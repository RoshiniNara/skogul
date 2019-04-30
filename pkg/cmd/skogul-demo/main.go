/*
 * skogul, main method/init
 *
 * Copyright (c) 2019 Telenor Norge AS
 * Author(s):
 *  - Kristian Lyngstøl <kly@kly.no>
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
 * 02110-1301  USA
 */

/*
Skogul is primarily a framework, where you use it to build your OWN
binaries. This package is provided to show-case a complete chain of
Skogul.

While it is usually possible to write it "the right way up" - start with
where we receive data and add senders - it's easier to write it "up-side-down"
and start with the final resting place for data. So you might consider reading
this bottom-up if it looks weird.

We are going to set up a chain of senders, starting with HTTP receiver which has
three distinct paths and handlers/chains: / sends to the primary handler,
/debug sends to a debug handler that just echos the parsed JSON to log.Print, and
/debug/notemplate does the same, but does not expand any provided template.

The primary chain looks sort of like this:

	http -> fallback -> delay -> counter ->  dupe2 -> postgres
	            \                    \             `-> influx
	             \			 `------------/
		      \
		       `- dupe --> Log(print "the following failed")
		               `-> debug(print json to stdout)

*/
package main

import (
	"github.com/KristianLyng/skogul/pkg"
	"github.com/KristianLyng/skogul/pkg/parsers"
	"github.com/KristianLyng/skogul/pkg/receivers"
	"github.com/KristianLyng/skogul/pkg/senders"
	"github.com/KristianLyng/skogul/pkg/transformers"
	"time"
)

func main() {
	// Let's start by setting up two "final" storage senders
	influx := &senders.InfluxDB{URL: "http://127.0.0.1:8086/write?db=test", Measurement: "test"}
	//postgres := &senders.Postgres{ConnStr: "user=postgres dbname=test host=localhost port=5432 sslmode=disable"}
	// Init is optional, see the skogul.senders.Postgres documentation
	//postgres.Init()

	// Set up a duplicator and hook influx and postgres up to it -
	// Everything going to the duplicator will go to both influx and
	// postgres.
	dupe2 := senders.Dupe{Next: []skogul.Sender{influx}}

	// Set up a handler for where to send statistics. In this case, we
	// just send it to influx.
	countHandler := skogul.Handler{
		Sender:       influx,
		Transformers: []skogul.Transformer{}}

	// The counter generates statistics for us every Period time
	// (assuming data) and sends it to the Stats-handler (here:
	// influx). While it might seem strange to have a handler instead
	// of just a Sender at first, this allows us to provide arbitrary
	// transformers to the stats, e.g.: add metadata.
	counter := &senders.Counter{Next: dupe2, Stats: countHandler, Period: 1 * time.Second}

	// Let's also inject a random delay for testing!
	delay := senders.Sleeper{counter, 5000 * time.Millisecond, false}

	fanout := senders.Fanout{Next: &delay}

	// Let's detach
	detach := senders.Detacher{Next: &fanout}

	// An other duplicator. This one just prints "The following failed"
	// and then uses the Debug-sender to print the metrics.
	dupe := senders.Dupe{Next: []skogul.Sender{senders.Log{"The following failed"}, senders.Debug{}}}

	// the Fallback sender tries to write to the delay-sender
	// (delay->counter->dupe2->{postgres,influx}), but if this
	// fails, it will write to the dupe-sender (print "the following
	// failed" and the request).
	fb := senders.Fallback{}
	fb.Add(&detach)
	fb.Add(&dupe)

	// That takes care of the sender-chains. Let's set up three
	// receiver handlers.

	// This is the "normal" one - send to the fallback sender and
	// that's it. It also has a single transformer that - prior to
	// sending the data on - expands any template provided.
	h := skogul.Handler{
		Parser:       parsers.JSON{},
		Sender:       &fb,
		Transformers: []skogul.Transformer{transformers.Templater{}}}

	// This is the same - but just print the request.
	debugtemplate := skogul.Handler{
		Parser:       parsers.JSON{},
		Sender:       senders.Debug{},
		Transformers: []skogul.Transformer{transformers.Templater{}}}

	// Print the request, but do NOT expand the template. Demonstrates
	// what a template does and what the template transformer does.
	debugnotemplate := skogul.Handler{
		Parser:       parsers.JSON{},
		Sender:       senders.Debug{},
		Transformers: []skogul.Transformer{}}

	// Set up a HTTP receiver
	receiver := receivers.HTTP{Address: "[::1]:8080"}

	// Add the various handlers to relevant paths.
	receiver.Handle("/", &h)
	receiver.Handle("/debug", &debugtemplate)
	receiver.Handle("/debug/notemplate", &debugnotemplate)

	// Start it
	receiver.Start()
}