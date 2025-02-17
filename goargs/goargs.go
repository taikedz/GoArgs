package goargs

import (
    "fmt"
    "os"
)


type VarDef interface {
    GetName() string
    Assign(string) error
}


type Parser struct {
    definitions []VarDef
    positionals []string
    passdown_args []string
}

func (p *Parser) fromName(longname string) *VarDef {
    for _, vdef := range p.definitions {
        if vdef.GetName() == longname { return &vdef }
    }
    return nil
}

func (p *Parser) Parse(args []string) error {
    for i := 0; i<len(args); i++ {
        token := args[i]
        var def VarDef = nil

        if token[:2] == "--" {
            longname := token[2:]
            if longname == "" {
                // We found the first delimiter for passdown arguments
                // Retain them as such, and stop parsing
                p.passdown_args = args[i+1:]
                return nil
            }
            def = *p.fromName(longname)

        }

        // TODO - support short names?

        if def != nil {
            i++
            // TODO - this won't apply to bool flag - check type
            // https://stackoverflow.com/a/7006853/2703818
            if i >= len(args) { return fmt.Errorf("Expected value after %s", token) }

            def.Assign(args[i])
        } else {
            p.positionals = append(p.positionals, token)
        }
    }

    return nil
}

// Supported types

func (p *Parser) StringArg(value *string, name string, defval string) {
    s := StringDef{name, value}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}

func (p *Parser) IntArg(value *int, name string, defval int) {
    s := IntDef{name, value}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}

// ---------

func SpotCheck() {
    //*
    var person string
    var age int
    var par Parser
    par.StringArg(&person, "name", "nobody")
    par.IntArg(&age, "age", -1)
    par.Parse(os.Args[1:])

    fmt.Printf("Parser: %v\n", par)
    fmt.Printf("%s -> %d\n", person, age)
    // */
}

