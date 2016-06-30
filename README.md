expenv - Expand environment variables [![Build Status](https://travis-ci.org/blang/expenv.svg?branch=master)](https://travis-ci.org/blang/expenv)
======
Expenv replaces ${var} or $var in the input to the values of the current environment variables. Supports stdin/stdout and files.

There is also a newer version available, [expenv.sh](https://github.com/blang/expenv.sh), written as shell script.

Usage
-----
```bash
$ go get github.com/blang/expenv
$ $GOPATH/bin/expenv -f inputfile
```

If you don't want to build it, you might want to download the [ELF x86-64 binary](https://github.com/blang/expenv/releases/latest) of the latest release (build by travis-ci).

Examples
-----

```bash
$ expenv -f inputfile > outputfile
$ expenv < inputfile > outputfile
$ expenv -i -f inputfile // Replace inplace
```

Example input:
```bash
My PWD is $PWD
Whoami: ${USER}
I'm using $TERM
Expand $empty but don't expand $$empty # => Expand  but don't expand $empty
```

Motivation
-----

I need to make config files more dynamic using environment variables. In a docker environment this is a big issue for me.

Contribution
-----

Feel free to make a pull request. For bigger changes create a issue first to discuss about it.


License
-----

See [LICENSE](LICENSE) file.
