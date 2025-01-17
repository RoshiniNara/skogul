
.. image:: https://goreportcard.com/badge/github.com/telenornms/skogul
   :target: https://goreportcard.com/report/github.com/telenornms/skogul

.. image:: https://godoc.org/github.com/telenornms/skogul?status.svg
   :target: https://godoc.org/github.com/telenornms/skogul

.. image:: https://cloud.drone.io/api/badges/telenornms/skogul/status.svg
   :target: https://cloud.drone.io/telenornms/skogul

======================================
Skogul - generic metric/data collector
======================================

Skogul is a generic tool for moving metric data around. It can serve as a
collector of data, but is primarily designed to be a framework for building
bridges between data collectors and storage engines.

This repository contains the Skogul library/package, and ``cmd/skogul``,
which parses a JSON-config to set up Skogul.

The release also includes extensive documentation autogenerated, provided
both as a manual page (``man skogul``), and RST in ``docs/skogul.rst`` once
built.

See also the examples in `docs/examples/basics`_ 

.. contents:: Table of contents
   :depth: 2
   :local:

Quickstart - RPM
----------------

If you're on CentOS/RHEL 7 or newer, you should use our naively built RPM,
available at https://github.com/telenornms/skogul/releases/latest.

It will install skogul, set up a systemd service and install a simple
configuration in ``/etc/skogul/conf.d/default.json``.

There's also a 64-bit Linux build there, which should work for most
non-RPM-based installations. But we make no attempt to really maintain it.

Quickstart - Docker
-------------------

A docker image is published to ghcr.io/telenornms/skogul as of version
v0.17.0. As of version v0.18.0, it is also published as "latest".

Quickstart - source
-------------------

Building from source is not difficult. First you need Golang. Get it at 
https://golang.org/dl/ (I think you want go 1.19 or newer).

Building ``skogul``, including cloning::

   $ git clone https://github.com/telenornms/skogul
   (...)
   $ make
   $ make install

You don't have to use make - there's very little magic involved for regular
building, but it will ensure ``-version`` works, along with building the
auto-generated documentation.

Running ``make install`` installs the binary and default configuration, but
does NOT install any systemd unit or similar.

Also see ``make help`` for other make targets.

About
-----

Skogul collects, mutates and transmits metric data. It was written to
support an extensive eco system of data collectors and storage engines in
constant motion. With Skogul, the goal is to disconnect how data is
collected from how it is stored: If you decide to change storage engine,
you should not have to even touch your collector. Or vice versa.

It can be used for very simple setups, and expanded to large,
multi-datacenter infrastructures with a mixture of new and old systems
attached to it.

To accomplish this, you set up chains that define how data is received, how
it is treated, where it goes and what happens if something goes wrong.

A Skogul chain is built from one or more independent receivers which
receive data and pass it on to a sender. A sender can either transmit data
to an external system (including an other Skogul instance), or add some
internal routing logic before passing it on to one or more other senders.

.. image:: docs/imgs/basic.png

Unlike most APIs or collectors of metrics, Skogul does NOT have a
preference when it comes to storage engine. It is explicitly designed to
disconnect the task of how data is collected from how it is stored.

The rationale is that the problem of writing an efficient SNMP collector
should not be tightly coupled to where you store the data. And where you
store the data should not be tightly coupled with how you receive it, or
what you do with it.

This enables an organization to gradually shift from older to newer stacks,
as Skogul can both receive data on old and new transport mechanisms,
and store it both in new and old systems. That way, older collectors can
continue working how they always how worked, but send data to Skogul.
During testing/maturing, Skogul can store the data in both legacy system
and replacement system. When the legacy system is removed, no change is
needed on the side of the collector.

Extra care has been put into making it trivial to write senders and
receivers.

See the package documentation over at godoc for development-related
documentation: 
https://godoc.org/github.com/telenornms/skogul

More discussion on architecture can be found in `docs/`.

Modules
-------

Skogul is all based around modules. We started out with three modules, but
today we have a total of five different module types. To see every module,
use ``skogul -help``, read the manual, or simply look in the receiver,
sender, parser, encoder and transformer directory and read the ``auto.go``
file there which lists all modules.

Each family of module has its own section in the JSON configuration file.

While all modules are fundamentally equal, receivers and senders are so
important that they are a little bit *more* equal, so we call them core
modules.

.. image:: docs/imgs/modules-drawio.png

............
Core modules
............

Senders and receivers
.....................

There are two essential module types, the receiver which defines how data
is acquired, and senders that determine what to do with the data. A handler
is just a collection of parser, optional transformer and reference to the
first sender.

Commonly used receivers are the HTTP receiver, UDP, kafka receiver
(consumer), various file receivers, SQL receiver

Senders all have the same general API, but come in two distinct types

External senders
................

External senders transmit data out of Skogul and are the classic and
easy-to-understand senders. Examples are InfluxDB sender to store data in
InfluxDB, UDP sender, Kafka sender (producer), SQL sender, MQTT, and more.

The debug or "print" sender is a little special: It just prints data to
stdout and is HIGHLY useful for testing.

Internal/Logic senders
......................

Logical senders are used internally to route or do something related with
data. The by far most important internal sender is the batch sender, which
accepts data, batches it into user-defined sizes, then passes them on to an
other sender. There are a large amount of small but important logical
senders that can be combined to form powerful chains.

Other dev favorites are: The count sender for getting statistics about how
much data passes through Skogul, the switch sender for sending data to
different other senders based on metadata, dupe sender for sending the same
data to multiple other senders, the null sender for simply discarding data,

...............
Support modules
...............

Additionally, three support-type modules exists:

Parsers
.......

The parser takes a set of bytes received and decodes it into a Skogul
internal container. E.g.: JSON decoding, protocol buffers for Juniper-data,
Influx Line protocol data, etc and is used by receivers through handlers.

Encoders
........

The encoder does the opposite: It takes an internal Skogul container and
encodes it as a byte stream for external tranmission. Today, only a small
amount of senders use encoders, as they are quite new, but they will be
used more extensively in the future. Currently, only JSON is supported.
More to come.

Transformers
............

Transformers are used to, you guessed it, transform or mutate parsed
containers. Typically used to re-arrange source data to better match target
data, to add metadata, or to sanitize data.

................
Implicit modules
................

All modules can be defined in configuration, but several modules have zero
configuration options, or very common options. E.g.: The `skogul` parser
doesn't require any configuration to work, the `debug` sender works fine
without any settings, the `now` transformer doesn't need any configuration
to add current time to a metric. To save you from having to define a whole
lot of empty modules, these type of modules can be referenced by their
implementation name (class, if you like) and an instance will be created
behind the scenes. These are listed as "auto modules" in the manual page.

E.g., without this feature::

        {
                "receivers": {
                        "foo": {
                                "type": "test",
                                "handler": "myhandler"
                        }
                },
		"handlers": {
			"myhandler": {
				"parser": "skogul",
				"sender": "debug"
			}
		},
                "parsers": {
                        "skogul": {
                                "type": "skogul"
                        }
                },
                "senders": {
                        "debug": {
                                "type": "debug"
                        }
                }
        }

But since the Skogul parser and the debug sender has no configuration, you
can just omit their definition and Skogul will implicitly create them for
you::

        {
                "receivers": {
                        "foo": {
                                "type": "test",
                                "handler": "myhandler"
                        }
                },
		"handlers": {
			"myhandler": {
				"parser": "json",
				"sender": "debug"
			}
		}
        }


Performance
-----------

Skogul is meant to scale well. Early tests on a laptop proved to work very
well:

.. image:: docs/imgs/skogul-rates.png

The above graph is from a very simple test on a laptop (with a quad core
i7), using the provided tester to write data to influxdb. It demonstrates
that despite well-known weaknesses at the time (specially in the
influx-writer), we're able to push roughly 600-800k values/s through
Skogul. This has since been exceeded.

This was an early test, and since then Skogul has been run in production on
large scale systems, and generally out-performs anything it communicates
with.

Name
----

Skogul is a Valkyrie. After extensive research (5 minutes on Wikipedia with
a cross-check on duckduckgo), this name was selected because it is
reasonably unique and is also a Valkyrie, like Gondul, a sister-project.

Hacking
-------

There is little "exotic" about Skogul hacking, so the following sections
are aimed mostly at people who are unfamiliar with Go.

A few sources for more documentation:

- docs/CODE_OF_CONDUCT.md
- docs/CONTRIBUTING
- docs/CODING
- doc.go

.......
Testing
.......

In short: Use ``make check``. It will run ``go test -short ./...`` and
various other checks. There's also ``make covergui`` to do coverage
analysis and open it in a browser.

``make check`` is run on every commit.

Use ``make fmtfix`` to fix formatting issues, which also makes sure to not
mess with the bundled/generated go files.

.............
Documentation
.............

Documentation comes in two forms. One is aimed at end-users. This is
provided mainly by adding proper labels to your data structures (see any
sender or receiver implementation), and through hard-coded text found in
``cmd/skogul/main.go``. In addition to this, stand-alone examples of setups
are provided in the ``examples/`` directory.

For development, documentation is written and maintained using code
comments and runnable examples, following the ``godoc`` approach. Some
architecture comments are kept in ``docs/``, but by and large,
documentation should be consumed from godoc.

See https://godoc.org/github.com/telenornms/skogul for the online
version, or use ``go doc github.com/telenornms/skogul`` or similar,
as you would any other go package.

Examples are part of the test suite and thus extracted from ``*_test.go``
where applicable. But aren't really used much.

Roadmap
-------

We are doing frequent releases on github. I honestly don't know why no 1.0
release has been made. Mainly lazyness, we've been "almost there" for 2+
years.

Overall, the core modules and the scaffolding is getting pretty good.

The bigger things moving right now except new modules is logging, which has
never been quite right, and dealing with some legacy/deprecation.

Similarly, test cases need to be refreshed. Tests are written very
isolated, and a good bit of spaghetti-logic has arisen. We have decent
coverage, but it's getting trickier to scale test case writing.

We also need better integration tests now that Skogul integrates with a
wide variety of services.

Other than that, there are modules to be written and features to be added
which are mostly a matter of what needs arise.
