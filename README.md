# GoArgs - a better Go Arguments Parser

Go's default `flag` library is rudimentary; GoArgs aims to provide a simple yet more featureful module for parsing arguments.
The usage style matches the standard `flag` library for compatibility, though it is not aiming to be a drop-in replacement.
The codebase remains small, ensuring it is easy to audit.

Compatibility:

* Pass pointer into argument delcaration (`flag.<Type>()` equivalents)
* Create pointer from argument declaration (`flag.<Type>Var()` equivalents)

Improved features:

* Flags can appear intermixed with positional arguments
* Parser operates on any developer-specified token list (not just `os.Args`)
* Parser recognises `--` as end of direct arguments, and stores subsequent raw tokens
* Parser can opt to ignore unknown flags, or return error on unknown arguments, as-needed.
* Unpacking methods `Unpack()` and `UnpackExactly()` help extract and assign positional arguments
* Long-name flags are specified only with double-hyphen notation (to support short flags combination notations)

Yet to implement:

* Flag event function (`flag.Func` equivalent)
* Bool flag counter (`flag.BoolFunc` equivalent)
* Additional types as found in `flag` standard lib
* StringChoices argument type
* Usage strings
* Help display
* Optional short flags (rune `-` to mean "no short flag")
* Short flags can be combined with single-hyphen notation (e.g. `-eux` for `-e -u -x`)

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
        if _, _, err := goargs.UnpackExactly(send_p.Args(), &file, &server); err != nil {
            fmt.Printf("%v\n", err)
            os.Exit(2)
        }

        DoSend(file, server, encrypt) // ...

    } else if action == "recv" {
        var recv_p goargs.Parser
        var file string
        var server string
        var decrypt bool

        recv_p.Bool(&decrypt, "decrypt", false)

        // Detect flags, isolate positionals and extras
        if err := recv_p.Parse(moreargs, false); err != nil {
            fmt.Printf("%v\n", err)
            os.Exit(1)
        }
        if _,_, count_err := goargs.UnpackExactly(recv_p.Args(), &server, &file); count_err != nil {
            fmt.Printf("%v\n", error)
            os.Exit(2)
        }

        // The extra args after "--" are passed along directly, raw
        ServerCommand(recv_p.PassdownArgs)
        DoSave(file, server) // ...
    } else {
        fmt.Printf("Unknown action: %s", action)
        os.Exit(10)
    }
}
```

