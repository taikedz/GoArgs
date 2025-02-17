package main

import (
    "fmt"
    "net.taikedz.goargs/goargs"
)

func main() {
    var par goargs.Parser

    var person string
    var age int
    var height float32
    var admit bool
    var debug bool

    par.StringArg(&person, "name", "nobody")
    par.IntArg(&age, "age", -1)
    par.FloatArg(&height, "height", 0.0)
    par.BoolArg(&admit, "admit", false)
    par.BoolArg(&debug, "debug", false)

    if err := par.ParseCliArgs(false); err != nil {
        fmt.Printf("!! %v\n", err)
    }

    if(debug) {
        fmt.Printf("Parser: %v\n", par)
    }
    fmt.Printf("%s (%.2f m) -> %d yo (%t)\n", person, height, age, admit)
}

