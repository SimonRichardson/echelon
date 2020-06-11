# Echelon walker

------

The walker server is a reasonable implementation of a daemon that is intended to
walk over the various keys within the storage to repair any damange caused by
partitions.

------

## Usage

The server intended to be run as a standalone implementation.

### Environmental variables

The echelon application has a set of environment variables to help tweak the
application for different setups (testing vs production). The application can
use the various strategies (see above) to then turn on an off various parts of
the application:

### Running

Running the echelon is relatively easy and can even be run side by side the
insert server by passing a different port to run on. If you just want to test
out the echelon application just run the following:

```bash
go run ./echelon-walker/main.go
```

Alternatively running the application with a different port, then just overwrite
the environmental variable.

```bash
HTTP_ADDRESS=":9002" go run echelon-walker/main.go
```

### Operations

Echelon expects to interact with a set of independent Redis instances, which
operators should deploy, monitor and manage. In general, Redis will use a lot of
RAM and comparatively little CPU and Echelon will use very little RAM and
comparatively large amount of CPU. It may make sense to co-locate a Echelon
instance with every Redis instance.
