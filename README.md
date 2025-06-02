# GoArgs - a simple and flexible Go arguments parser

Go's default `flag` library feels limited as a basic library:

* flags must come before positional arguments
* no support for `--` "raw" passdown tokens after main CLI arguments
* no support for short flags and aggregated short flags (`-x -y -z` as `-xyz`)

GoArgs aims to provide a simple yet more flexible module for parsing arguments.

The codebase aims to remain small, ensuring it is easy to audit as an external dependency. It is not as fully featured as other implementations out there. See [alternatives](#alternatives) for more options.

## To Do

* Documentation
* Make relevant items private
* Re-evaluate multi-error-type returns/error classfication

## Features

Compatibility wtih `flag`:

* Create pointer from argument declaration (`flag.<Type>()` equivalents)
* Pass pointer into argument delcaration (`flag.<Type>Var()` equivalents)
* Flag event function (`flag.Func` equivalent)

Types:

* Basic: String, Int, Int64, Uint, Float, Float64, Bool, time.Duration
* Additional flag types:
    * Counter: incerments a counter every time the flag is seen (such as `-vvv` for incresed levels of verbosity)
    * Choices: predefine a number of possible values for a given flag
    * Appender: allow using the same flag multiple times (`--mount /this:/right/here --mount /that:/over/there` for two mounts)

Improved features:

* Flags can appear intermixed with positional arguments
* Parser operates on any developer-specified `[]string` of tokens (not just `os.Args`)
* Parser recognises `--` as end of direct arguments, and stores subsequent "raw" passdown tokens
* Parser can opt to ignore unknown flags, or return error on unknown arguments, as-needed.
* Unpacking methods `Unpack()` and `UnpackExactly()` help extract and parse positional arguments (supported var types: `*string`, `*int`, `*float`, `*bool`)
* Long-name flags are specified only with double-hyphen notation
* Short flags notation (`Parser.SetShortFlag("v", "verbose")`)
    * Short flags can be combined with single-hyphen notation (e.g. `-eux` for `-e -u -x`, or `-vv` for `-v -v` or `--verbose --verbose`)
* Help obtainable as string or printed; help arguments always listed in declaration order
* Help flags `--help` and `-h` automatically detected when using `ParseCliArgs()`, except when appearing after the first `--`


## Examples

A simple example of some of the basic notations.

```go
// Declare a parser with help
parser := goargs.NewParser("cmd ARG1 ARG2 --opt OPTVAL --num N ...")

// Use a reference variable from the parser
opt := parser.String("opt", "hello", "A simple word")
// Set a short flag for the option by name
parser.SetShortFlag('o', "opt")

// Create your own variable and pass it in
var count int
parser.Intvar(&count, "num", 0, "A number")
parser.SetShortFlag('n', "num")

parser.ParseCliArgs() // parse the os.Args[...] values

// ... or parse values from an array/slice:
/*
tokens := []string{"one", "two", "-n", "5"}
parser.Parse(tokens)
*/

fmt.Printf("%v, %v -> %v", count, *opt, parser.Args())
```

Some runnable example files are linked below. For further examples, see [unit tests](./unittests/)

[salute.go](examples/salute.go)

Usage:

```sh
go run examples/salute.go Alex Sam -W 2 --grouped Jay

# Prints:
# Hello Alex, Sam and Jay 👋👋
```

[splat.go](examples/splat.go)

Try running the following:

```sh
go run examples/splat.go w one.txt two.txt -- "hi and bye" "this and that"
go run examples/splat.go w one.txt --help two.txt -- "hi and bye" "this and that"
go run examples/splat.go w one.txt two.txt -- "hi and bye" --help "this and that"
```

## Alternatives

Why use this `taikedz/GoArgs` ? If your needs are minimal and/or you literally need to copy the files in to your project directly, then maybe you have a case to use this lib. I have made a point of keeping the feature set relatively straighforward and flexible, and in keeping with the standard library's style. I have also tried to make a point to keep the code itself straightforwad so that you may audit it.

Elsewise, please treat it as a learning tool for its easy-to-read implementation. This package did begin as a learning project, started whilst on an airplane.

More-established packages exist that also offer partial drop-in capability, as well as support for combined short options, and the `--` arguments sequence separator:

* Two which I am aware of:
    * <https://pkg.go.dev/github.com/spf13/pflag>
    * <https://pkg.go.dev/github.com/jessevdk/go-flags>
* Search the go package listings in general: <https://pkg.go.dev/search?q=flag&m=>
