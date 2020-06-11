# Echelon shim

------

The shim server is a reasonable implementation for easily talking to the
echelon-http implementation using normal inbound json content-type.

------

## Usage

The server intended to be run as a shim between nginx and the standalone
implementation.

### Environmental variables

The echelon application has a set of environment variables to help tweak the
application for different setups (testing vs production). The application can
use the various strategies to then turn on an off various parts of the
application.

### Running

Running the echelon is relatively easy and can even be run side by side the
insert server by passing a different port to run on. If you just want to test
out the echelon application just run the following:

```bash
go run ./echelon-shim/main.go
```

Alternatively running the application with a different port, then just overwrite
the environmental variable.

```bash
HTTP_ADDRESS=":9002" go run echelon-shim/main.go
```
