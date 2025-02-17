package main

import (
    "fmt"
    "net.taikedz.goargs/goargs"
)

func main() {
    // Try with:
    // go run GoArgs.go Bob --name=Jay --age 12 --height=1.7 --admit 15 and the rest --debug -- There is more --oops
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
        //return
    }
    fmt.Printf("%s (%.2f m) -> %d yo (%t)\n", person, height, age, admit)

    var person2 string
    var age2 int
    remains, err := goargs.Unpack(par.positionals, &person2, &age2)
    if err != nil {
        fmt.Printf("!! %v\n", err)
        //return
    }
    fmt.Printf("Unpacked: %s -> %d\nRemains: %v\n", person2, age2, remains)
}

