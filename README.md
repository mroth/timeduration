timeduration
============

timeduration is pretty much the simplest possible wrapper around the
[`time.ParseDuration()`][1] function in Go, designed for a binary to let me
specify durations in more convenient formats for CLI utilities that expect to
receive an argument in number of seconds.

This isn't even really a program, it's the sort of thing that normally would
just go into a ZSH function in my dotfiles, but this seemed like a faster way to
accomplish it.  Really this is only publicly on GitHub to make it easier for me
to drop it on new workstations. :information_desk_person:

Example
-------
Base case:

    $ timeduration 3h15m
    11700

But really this is mostly for using in as an argument to another command line
utility, for example I use the utility Darwin utility `caffeinate` to prevent my
iMac from going to sleep when streaming a movie via DLNA to my Playstation.

    $ caffeinate -s -t `timeduration 2h10m`

Pretty simple but faster than doing mental math.  Typically in real world usage
I alias this to something short like `td` so it can easily just become part of
quick shell usage.

Installation
------------
Just do `go get github.com/mroth/timeduration`.  Or just glance at the source
code and then write your own, really.

License
-------
C'mon, do you really need one? This thing is so ludicrously simple, it's just a
quick script for my own use.

...Fine, available under the [WTFPL][2].

[1]: http://golang.org/pkg/time/#ParseDuration
[2]: http://www.wtfpl.net
