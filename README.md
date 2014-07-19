expenv - Expand environment variables
======
Expenv replaces ${var} or $var in the input to the values of the current environment variables. Supports stdin/stdout and files.

Usage
-----
```bash
$ go get github.com/blang/expenv
$ $GOPATH/bin/expenv -f inputfile
```


Examples
-----

```bash
$ expenv -f inputfile > outputfile
$ expenv < inputfile > outputfile
$ expenv -i -f inputfile // Replace inplace
```

Example input:
```
My PWD is $PWD
Whoami: ${USER}
I'm using $TERM
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