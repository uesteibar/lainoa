# Lainoa

![ci](https://github.com/uesteibar/lainoa/workflows/ci/badge.svg?branch=master)

Programming language built following the interpreterbook.com

## Running locally

You'll need [golang](https://golang.org/) installed.

Install it:

```
make install
```

If you don't want to install it, you can also build the binary to `bin/lainoa`:

```
make build
```

Run the REPL:

```
lainoa repl
Hello uesteibar! This is the Lainoa programming language

Go ahead and type some stuff!
⛅️ >>
```

## Running with docker

```
> docker run -it uesteibar/lainoa repl
Hello root! This is the Lainoa programming language

Go ahead and type some stuff!
⛅️ >>
```
