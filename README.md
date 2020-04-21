# c2c

c2c is a command line utility which converts the log files produced by the Caddy web server into the standard combined log format. Thus, making it easier to process with a log file analyser like GoAccess.

The binaries folder contains pre-compiled executables for 64 bit machines. There is a version for Windows and Linux.

# Example Usage

c2c has been designed to work from the command line so that the log file can be piped into the command. The converted output can then be piped to an output file like:

``` bash
./c2c < caddy.log > combined.log
```

This allows c2c to be used as part of a cron job which processes the log file(s) every evening.