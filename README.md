# Lainoa

![ci](https://github.com/uesteibar/lainoa/workflows/ci/badge.svg?branch=master)

Built following the https://interpreterbook.com.

I never got far enough in college to the point where I could take the _compilers_ subject,
and that's been an itch I've always wanted to scratch. So this is me scratching that itch!

Lainoa is an [interpreted language](https://en.wikipedia.org/wiki/Interpreted_language) writen in
[go](https://golang.org/). It's pretty slow, doesn't have a lot of features, and you shouldn't
be using it for anything sane. It's here for educational purposes, and I had a blast 🎉 building it.

Ah, `lainoa` stands for `cloud` in [Euskera](https://en.wikipedia.org/wiki/Basque_language).

## Installing and running

### Running locally

You'll need [golang](https://golang.org/) installed.

Install it:

```
make install
```

If you don't want to install it, you can also build the binary to `bin/lainoa`:

```
make build
```

### Run a file

```
# examples/hello_world.ln

let name = "World"

puts("Hello " + name + "!")
```

```
> lainoa run examples/hello_world.ln
Hello World!
```

### Run the REPL:

```
lainoa repl
Hello uesteibar! This is the Lainoa programming language

Go ahead and type some stuff!
⛅️ >>
```

or with docker

```
> docker run -it uesteibar/lainoa repl
```

## Features

Lainoa is as simple as a programming language can get.


You have numbers and can do fun math with them:

```
let one = 1
let two = 2

let five = one + two * 2
```

Strings are there too:

```
let name = "Unai"
let last_name = "Esteibar"

let full_name = name + " " + last_name # => "Unai Esteibar"

puts(full_name) # prints Unai Esteibar
```

And of course booleans and boolean operations:

```
let lainoa_is_cool = true
let should_i_use_it = false

15 > 10
15 < 10
15 == 10
15 != 10
```

You can declare functions and pass them around:

```
let result = 0

let add = fun(a) {
  let number = a

  return fun(b) {
    return number + b
  }
}

let addFive = add(5)

result = addFive(10)
```

There's also conditionals of course, otherwise life would be pretty boring:

```
let status = fun(age) {
  if (age < 18) {
    "little-adult"
  } else {
    "adult"
  }
}

status(21) # => "adult"
status(17) # => "little-adult"
```

Arrays, because otherwise how would you build a ToDo app?

```
let shopping_list = [
  "milk",
  "cereals",
  "bread"
]

puts(shopping_list[0]) # => "milk"
puts(shopping_list[1]) # => "cereals"
puts(shopping_list[2]) # => "bread"
```

Oh, you can use `;` if you want to do things inline, but they're not mandatory otherwise:

```
let a = 1; a = a + 2;
```

For funky cases, there's `nil`:

```
let a = if (false) { "never" }

puts(a) # => nil
```

