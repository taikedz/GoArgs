# GoArgs - a better Go Arguments Parser

Go's default `flag` library is rudimentary; GoArgs aims to provide a simple yet more featureful module for parsing arguments.

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

This basic feature set allows flexible argument parsing in a simple package. The codebase remains small, ensuring it remains easy to audit.

```go
// Flags can appear before, after, or in between positionals

// go run tool.go send ./thing --encrypt remote.lan
// go run tool.go recv remote.lan ./stuff -- nc 12.34.56.78 3000 "<" file.txt

func main() {
    var command goargs.Parser
    var action string

    // Use `Unpack()` for processing positional arguments
    moreargs := goargs.Unpack(os.Args[1:], &action)
    
    if action == "send" {
        var send_p goargs.Parser
        var file string
        var server string
        var encrypt bool

        send_p.BoolArg(&encrypt, "encrypt", false)
        if err := send_p.Parse(moreargs, false); err != nil {
            fmt.Printf("%v\n", err)
            os.Exit(1)
        }
        goargs.Unpack(send_p.Positionals, &file, &server)

        DoSend(file, server, encrypt) // ...

    } else if action == "recv" {
        var recv_p goargs.Parser
        var file string
        var server string
        var encrypt bool

        if err := recv_p.Parse(moreargs, false); err != nil {
            fmt.Printf("%v\n", err)
            os.Exit(1)
        }
        goargs.Unpack(recv_p.Positionals, &server, &file)

        ServerCommand(recv_p.PassdownArgs) // all tokens after the first "--"
        DoSave(file, server) // ...
    }
}
```

