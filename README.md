logrus formatters
=================

[![Build Status](https://travis-ci.org/meatballhat/logrus-formatters.svg?branch=master)](https://travis-ci.org/meatballhat/logrus-formatters)

## l2met formatter

An [l2met](http://r.32k.io/l2met-introduction) formatter for
[logrus](https://github.com/Sirupsen/logrus) adapted from work by
[@henrikhodne](https://github.com/henrikhodne).  Take a look at the
[example web app](./examples/l2met-web.go), or peek at this:

``` golang
// in file "web.go"

log := logrus.New()
log.Formatter = l2met.NewFormatter("myapp:")

// ...

log.WithFields(logrus.Fields{"mood": "stellar", "shrugs": nshrugs}).Info("wat")
```

Such an example might produce a log record like this:

```
myapp: mood=stellar shrugs=99 msg=wat filename=/path/to/web.go
```
