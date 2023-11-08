# Temporal library for Slack activities


## Development

### Requirements

- Go 1.19+
- GNU Make
- git

### Building

This package does not generate any direct usable assets (it's a library). You can still run the build scripts to validate your code, and generate coverage information.

```bash
# Default target is 'build'
$ make

# Explicitly run build
$ make build

```

### Testing

You can run the unit tests with the supplied `All Tests` run config if you're using a Jetbrains IDE, or you can run them via `make` with 
the following command
```bash
$ make test
```
