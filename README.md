# GoArgs - a better Go Arguments Parser

Go's default `flag` library is rudimentary; GoArgs aims to provide a simple yet more featureful module for parsing arguments.
The usage style matches the standard `flag` library for compatibility.
The codebase remains small, ensuring it remains easy to audit.

Notably:

* Flags can appear intermixed with positional arguments
* Long-name flags are specified with double-dash notation
* Parser operates on any developer-specified token list (not just `os.Args`)
* Parser can opt to ignore unknown flags, or return error on unknown arguments, as-needed.
* Parser recognises `--` as end of relevant arguments, and stores subsequent raw tokens

Yet to implement:

* Pointer-creation (`flag.<type>Var()` equivalents)
* Flag event function (`flag.Func` equivalent)
* Usage strings
* Help display
* Optional short flags (rune `-` to mean "no short flag")
* Short flags can be combined with single-dash notation

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
        if _, _, err := goargs.UnpackExactly(send_p.Positionals, &file, &server); err != nil {
            fmt.Printf("%v\n", err)
            os.Exit(2)
        }

        DoSend(file, server, encrypt) // ...

    } else if action == "recv" {
        var recv_p goargs.Parser
        var file string
        var server string
        var decrypt bool

        recv_p.BoolArg(&decrypt, "decrypt", false)

        // Detect flags, isolate positionals and extras
        if err := recv_p.Parse(moreargs, false); err != nil {
            fmt.Printf("%v\n", err)
            os.Exit(1)
        }
        if _,_, count_err := goargs.UnpackExactly(recv_p.Positionals, &server, &file); count_err != nil {
            fmt.Printf("%v\n", error)
            os.Exit(2)
        }

        // The extra args after "--" are passed along directly, raw
        ServerCommand(recv_p.PassdownArgs)
        DoSave(file, server) // ...
    }
}
```

