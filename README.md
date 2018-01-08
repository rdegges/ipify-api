# ipify-api

*A Simple Public IP Address API*


This repository contains the source code for [ipify](https://www.ipify.org), one
of the largest and most popular IP address API services on the internet. ipify
serves over 30 billion requests per month!


## What does ipify do?

Have you ever needed to pragmatically get your public IP address? This is quite
common for developers provisioning cloud servers, for instance, where you might
be creating servers and running bootstrapping software on them without access to
server metadata.

Being able to quickly and reliably get access to your public IP address is
essential for configuring DNS, managing external services, and a number of other
operationally related tasks.

In general, there are a number of uses for public IP address information.


## What is ipify?

ipify is a free API service anyone can use to get their public IP address. It is
highly reliable (built on top of [Heroku](https://www.heroku.com/)) and fast.
Typical response times (server side) are between 1ms and 10ms.

ipify is also fully funded -- it's been running for years and isn't going
anywhere. The people behind ipify cover all expenses and maintenance, so you
can feel safe integrating with it knowing it won't be disappearing.

If you'd like to use ipify in your application, no permission is needed. You can
immediately start using the service without any restrictions. Simply visit our
[public website](https://www.ipify.org) for more information.


## What is this project?

This project is the source code that powers the ipify service. ipify is written
in the Go programming language for speed and efficiency purposes. You can read
an [article](https://www.rdegges.com/2018/to-30-billion-and-beyond/) written by
ipify's creator, [Randall Degges](https://twitter.com/rdegges), if you'd like
more information.

If you'd like to contribute to ipify's development, you can do so here. Pull
requests are encouraged.

Finally, if you'd like to deploy your own instance of ipify, you can easily do
so. Compiling this project will produce a single statically linked binary that
is designed to be run on Heroku. With minor modification, ipify can be ran on
any web hosting platform.

Please contact [Randall](mailto:r@rdegges.com) if you need assistance deploying
your own copy of ipify onto a non-Heroku host.


## Building ipify

To develop and build ipify, you'll need to have the Go programming language
setup on your computer. If you don't, you can read more about it here:
https://golang.org/

Once you have Go installed, you'll need to clone this project into your
computer's GOPATH. For me, this means I'll typically do something like:

```bash
$ git clone https://github.com/rdegges/ipify-api.git ~/go/src/github.com/rdegges/ipify-api
```

To build the project, change to the project directory and run:

```bash
$ go build
```

This will create the `ipify-api` binary in the current directory that you can
use for testing.


## Deploying ipify

If you'd like to deploy your own version of ipify to Heroku, you can do so
easily by clicking the button below. This will take you to Heroku and let you
instantly provision your own copy of the ipify service.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)


## Questions?

Got a question? Please create a Github issue!
