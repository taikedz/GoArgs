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
}

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

func (p *Parser) Parse(args []string) {
    for i, str := range args {
        if i >= len(p.definitions) { return }

        p.definitions[i].Assign(str)
    }
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

    fmt.Printf("Definitions: %v\n", par.definitions)
    fmt.Printf("%s -> %d\n", person, age)
    // */
}
