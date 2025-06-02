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

