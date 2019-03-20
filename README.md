Fsync bench
===========

Benchmark fsync

Download a [binary from github](https://github.com/ncw/fsyncbench/releases/latest)
or build from source (see later).

Usage
=====

    fsyncbench [flags]

Flags

    -n=100: Iterations to test - default 100
    -dir=/tmp: Target directory - default os.TmpDir()

It will produce output like this

    $ fsyncbench 
    2013/07/02 16:03:30 That took 988.255164ms for 100 fsyncs
    2013/07/02 16:03:30 That took 9.882551ms per fsync

Build
=====

You'll need go installed, then 

    go get github.com/ncw/fsyncbench

and this will build the binary in `$GOPATH/bin`.  You can then modify
the source and submit patches.

License
=======

This is free software under the terms of the MIT license (check the
COPYING file included in this package).

Contact and support
===================

The project website is at:

- https://github.com/ncw/fsyncbench

There you can file bug reports, ask for help or contribute patches.

Authors
=======

- Nick Craig-Wood <nick@craig-wood.com>

Contributors
------------

- Your name goes here!
- Matteo Olivi <matteoolivi7@gmail.com>
