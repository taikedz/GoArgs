# GoArgs - a simple and flexible Go arguments parser

Go's default `flag` library feels limited as a basic library:

* flags must come before positional arguments
* no support for `--` "raw" passdown tokens after main CLI arguments
* no support for short flags and aggregated short flags (`-x -y -z` as `-xyz`)

GoArgs aims to provide a simple yet more flexible module for parsing arguments, with similar feel to the default `flag` package.

The codebase aims to remain small, ensuring it is easy to audit as an external dependency. It is not as fully featured as other implementations out there.

I hope it is useful to someone who just needs something straighforward to get things done. For a real life usage o fthis package, see [AlpackaNG](https://github.com/taikedz/alpacka-ng) by the same author.

## To Do

* Documentation overview
* Re-evaluate multi-error-type returns/error classfication

## Features

Similarity wtih `flag` typical usage:

* Create pointer from argument declaration (`flag.<Type>()` equivalents)
* Pass pointer into argument delcaration (`flag.<Type>Var()` equivalents)
* Flag event function (`flag.Func` equivalent)

Types:

* Basic: String, Int, Int64, Uint, Float, Float64, Bool, time.Duration
* Additional flag types:
    * Choices: predefine a number of possible values for a given flag
    * Counter: increments a counter every time the flag is seen (such as `-v -v -v` or `-vvv` for incresed levels of verbosity)
    * Appender: allow using the same flag multiple times (`--mount /this:/right/here --mount /that:/over/there` for two mounts)
    * Mode : define groups of mutually-exclusive flags, with the last-sepcified flag taking priority over the previous ones

Improved features:

* Flags can appear intermixed with positional arguments
* Long-name flags are specified only with double-hyphen notation
* Short flags notation (`Parser.SetShortFlag("v", "verbose")`)
    * Short flags can be combined with single-hyphen notation (e.g. `-eux` for `-e -u -x`, or `-vv` for `-v -v` or `--verbose --verbose`)
* Parser operates on any developer-specified `[]string` of tokens (not just `os.Args`)
* Parser recognises `--` as end of direct arguments, and stores subsequent "raw" passdown tokens
* Parser can opt to ignore unknown flags (storing them unparsed as positionals), or return error on unknown arguments, as-needed.
* Unpacking methods `Unpack()` and `UnpackExactly()` help extract and parse positional arguments (supported var types: `*string`, `*int`, `*float`, `*bool`)
* Help obtainable as string or printed; help arguments always listed in declaration order

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

Some runnable example files are linked below. For further examples, see [examples](./examples/):

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
