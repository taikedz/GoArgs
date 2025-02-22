# GoArgs - a simple and versatile Go arguments parser

Go's default `flag` library feels rudimentary; GoArgs aims to provide a simple yet more featureful module for parsing arguments.
The usage style matches the standard `flag` library for limited compatibility.

The codebase aims to remain small, ensuring it is easy to audit as an external dependency. It is not as fully featured as other implementations out there. See [alternatives](#alternatives) for more options.

Note that GoArgs does not intend to be a drop-in replacement for `flags` given its intent to improve on some behaviours. Simple usage of `flags` may be able to drop-in `goargs`, but
this is not a design goal.

Compatibility wtih `flag`:

* Create pointer from argument declaration (`flag.<Type>()` equivalents)
* Pass pointer into argument delcaration (`flag.<Type>Var()` equivalents)

Types:

* Basic: String, Int, Int64, Uint, Float, Float64, Bool, time.Duration (for general usage)
* Counter: incerments a counter every time the flag is seen (such as `-vvv` for incresed levels of verbosity)

Improved features:

* Flags can appear intermixed with positional arguments
* Parser operates on any developer-specified token list (not just `os.Args`)
* Parser recognises `--` as end of direct arguments, and stores subsequent raw tokens
* Parser can opt to ignore unknown flags, or return error on unknown arguments, as-needed.
* Unpacking methods `Unpack()` and `UnpackExactly()` help extract and assign positional arguments
* Long-name flags are specified only with double-hyphen notation (to support short flags combination notations)
* Help obtainable as string or printed; help arguments printed in declaration order
* Help flags `--help` and `-h` automatically detected when using `ParseCliArgs()`, except when appearing after first `--`

Yet to implement:

* Compatibility
    * Flag event function (`flag.Func` equivalent)
* StringChoices argument type
* Short flags
    * Specify short notation for boolean flags (`Parser.ShortFlag("v", "verbose")`, compatible with `flag.BoolFunc`)
    * Short flags can be combined with single-hyphen notation (e.g. `-eux` for `-e -u -x`, or `-vv` for `-v -v` or `--verbose --verbose`)

## Example

A basic example of usage. For further examples, see [goargs_test.go](./goargs_test.go) unit tests file

```go
// Flags can appear before, after, or in between positionals

// Compare example commands:
//   go run tool.go send ./thing remote.lan
//   go run tool.go recv remote.lan --decrypt ./stuff -- nc 12.34.56.78 3000 "<" file.txt
package main
import (
    "fmt"
    "os"
    "net.taikedz.goargs/goargs"
)

func main() {
    var command goargs.Parser
    var action string

    // Use `Unpack()` for processing leftmost positional arguments
    //   and retain the remainder in `moreargs`
    moreargs := goargs.Unpack(os.Args[1:], &action)
    
    if action == "send" {
        var send_p goargs.Parser
        var file string
        var server string

        // Use the parser to detect any/unexpected flags
        if err := send_p.Parse(moreargs, false); err != nil {
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
        var recv_p goargs.Parser
        var file string
        var server string

        // Declare an argument and variable to access result from
        decrypt := recv_p.Bool("decrypt", false)

        // Detect flags, isolate positionals and extras
        if err := recv_p.Parse(moreargs, false); err != nil {
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

Why use `taikedz/GoArgs` ? If your needs are minimal and you literally need to copy the files in, or to fully understand it (e.g. for auditability), then maybe you have a case to use this lib. Elsewise, please treat it as a learning tool for its easy-to-read implementation. This package did begin as a learning project, started whilst on an airplane.

More-established packages exist that also offer partial drop-in capability, as well as support for short options, and `--` terminated arguments sequences:

* <https://pkg.go.dev/github.com/spf13/pflag>
* <https://pkg.go.dev/github.com/jessevdk/go-flags>
