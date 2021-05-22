# Shorty
**Shorty** is an awesome service to shorten urls.

## Requirement

* [**Go (1.13.8 or later)**](https://golang.org/doc/install)

## Installation

Please install the Go at least 1.13.8 and make sure it works perfectly on your local machine.

After Go have been installed, please clone **Shorty** project into your local machine.

```
> git clone git@github.com:irfanalfarabbi/shorty.git
```

After **Shorty** already cloned, go to its directory then run it in your local machines.

```
> cd shorty
> make run
```

## Other Useful Commands

To run all the unit tests:

```
> make test
```

For the quick tests:

```
> make testlocal
```

To compile **Shorty** to a binary file:

```
> make compile
```

To build a docker image of **Shorty** locally:

```
> make build
```
