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
* Unpacking methods `Unpack()` and `UnpackExactly()` help extract and assign positional arguments
* Long-name flags are specified only with double-hyphen notation
* Short flags notation (`Parser.SetShortFlag("v", "verbose")`)
    * Short flags can be combined with single-hyphen notation (e.g. `-eux` for `-e -u -x`, or `-vv` for `-v -v` or `--verbose --verbose`)
* Help obtainable as string or printed; help arguments always listed in declaration order
* Help flags `--help` and `-h` automatically detected when using `ParseCliArgs()`, except when appearing after the first `--`


## Example

A basic example of usage. For further examples, see [unit tests](./unittests/)

```go
// Flags can appear before, after, or in between positionals

// Compare example commands:
//   go run tool.go send ./thing remote.lan
//   go run tool.go recv remote.lan --decrypt ./stuff -- nc 12.34.56.78 3000 "<" file.txt
package main
import (
    "fmt"
    "os"
    "github.com/taikedz/goargs/goargs"
)

func main() {
    var action string

    // Use `Unpack()` for processing leftmost positional arguments
    //   and retain the remainder in `moreargs`
    moreargs := goargs.Unpack(os.Args[1:], &action)
    
    if action == "send" {
        send_p := goargs.NewParser("send FILE SERVER")
        var file string
        var server string

        // Use the parser to detect any/unexpected flags
        // And automatically produce help text if "--help" is in the args
        if err := send_p.Parse(moreargs); err != nil {
            fmt.Printf("%v\n", err)
            os.Exit(1)
        }
        // Unpack the positionals now that eventual flags have been removed
        // Expect the exact number of tokens to number of variables
        if err := goargs.UnpackExactly(send_p.Args(), &file, &server); err != nil {
            fmt.Printf("%v\n", err)
            os.Exit(2)
        }

        DoSend(file, server, encrypt) // ...

    } else if action == "recv" {
        recv_p := goargs.NewParser("recv [--decrypt] SERVER FILE -- SERVER_COMMAND ...")
        var file string
        var server string

        // Declare an argument and variable to access result from
        decrypt := recv_p.Bool("decrypt", false, "Attempt decrypt on incoming data")
        // Also allow the flag to have a short notation
        recv_p.SetShortFlag('d', "decrypt")

        // Detect flags, isolate positionals and extras
        if err := recv_p.Parse(moreargs); err != nil {
            fmt.Printf("%v\n", err)
            os.Exit(1)
        }
        if count_err := goargs.UnpackExactly(recv_p.Args(), &server, &file); count_err != nil {
            fmt.Printf("%v\n", error)
            os.Exit(2)
        }

        // The extra args after "--" are passed along directly, raw
        ServerCommand(recv_p.PassdownArgs)
        DoSave(file, server, *decrypt) // ...
    } else {
        fmt.Printf("Unknown action: %s", action)
        os.Exit(10)
    }
}
```

## Alternatives

Why use this `taikedz/GoArgs` ? If your needs are minimal and/or you literally need to copy the files in, then maybe you have a case to use this lib. I have made a point of keeping the feature set relatively straighforward and flexible, and in keeping with the standard library's style. I have also tried to make a point to keep the code itself straightforwad so that you may audit it.

Elsewise, please treat it as a learning tool for its easy-to-read implementation. This package did begin as a learning project, started whilst on an airplane.

More-established packages exist that also offer partial drop-in capability, as well as support for short options, and `--` terminated arguments sequences:

* Two which I am aware of:
    * <https://pkg.go.dev/github.com/spf13/pflag>
    * <https://pkg.go.dev/github.com/jessevdk/go-flags>
* Search the go package listings in general: <https://pkg.go.dev/search?q=flag&m=>
