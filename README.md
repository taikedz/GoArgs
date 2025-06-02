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
* Unpacking methods `Unpack()` and `UnpackExactly()` help extract and parse positional arguments (supported var types: `*string`, `*int`, `*float`)
* Long-name flags are specified only with double-hyphen notation
* Short flags notation (`Parser.SetShortFlag("v", "verbose")`)
    * Short flags can be combined with single-hyphen notation (e.g. `-eux` for `-e -u -x`, or `-vv` for `-v -v` or `--verbose --verbose`)
* Help obtainable as string or printed; help arguments always listed in declaration order
* Help flags `--help` and `-h` automatically detected when using `ParseCliArgs()`, except when appearing after the first `--`


## Examples

A basic example of usage. For further examples, see [unit tests](./unittests/)

```go
package main

import (
    "os"
    "fmt"
    "strings"

    "github.com/taikedz/goargs/goargs"
)

const WAVE_MOJI string = "👋"

func main() {
    // Declare a new parser, and its basic help string
    parser := goargs.NewParser("salute NAME [--wave N] [--grouped] [--with SALUTATION]")

    // Declare a string flag "--with", its default value "Hello", and its help string
    // The returned value is a pointer to a memory location of type string
    salutation := parser.String("with", "Hello", "Salutation to use")
    parser.SetShortFlag('w', "with")

    // Again with an int flag
    wave_count := parser.Int("wave", 0, fmt.Sprintf("Add N hand wave emojis (%s)", WAVE_MOJI) )
    parser.SetShortFlag('W', "wave")

    // Again with a bool flag
    grouped := parser.Bool("grouped", false, "Group names to a single salutation")
    parser.SetShortFlag('g', "grouped")

    // Perform the parse. If `--help` is found amongst the flags, prints the help and exits
    if err := parser.ParseCliArgs(); err != nil {
        fmt.Printf("! -> %v\n", err)
        os.Exit(1)
    }

    // The variable is a pointed, remember to dereference!
    mojis := wave(*wave_count)

    if *grouped {
        names := parser.Args()
        lead_names := strings.Join(names[:len(names)-1], ", ")
        last_name := names[len(names)-1]

        fmt.Printf("%s %s and %s %s\n", *salutation, lead_names, last_name, mojis)
    } else {
        for _, name := range(parser.Args()) {
            fmt.Printf("%s, %s %s\n", *salutation, name, mojis)
        }
    }
}

func wave(times int) string {
    var hands []string

    for i:=0; i<times; i++ {
        hands = append(hands, WAVE_MOJI)
    }

    return strings.Join(hands, "")
}
```

Usage:

```sh
go run salute.go Alex Sam -W 2 --grouped Jay

# Prints:
# Hello Alex, Sam and Jay 👋👋
```

## Alternatives

Why use this `taikedz/GoArgs` ? If your needs are minimal and/or you literally need to copy the files in, then maybe you have a case to use this lib. I have made a point of keeping the feature set relatively straighforward and flexible, and in keeping with the standard library's style. I have also tried to make a point to keep the code itself straightforwad so that you may audit it.

Elsewise, please treat it as a learning tool for its easy-to-read implementation. This package did begin as a learning project, started whilst on an airplane.

More-established packages exist that also offer partial drop-in capability, as well as support for combined short options, and the `--` arguments sequence separator:

* Two which I am aware of:
    * <https://pkg.go.dev/github.com/spf13/pflag>
    * <https://pkg.go.dev/github.com/jessevdk/go-flags>
* Search the go package listings in general: <https://pkg.go.dev/search?q=flag&m=>
