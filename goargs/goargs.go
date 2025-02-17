package goargs

import (
    "fmt"
    "os"
)


// FIXME - VarDef is not seen outside the package, make methods private
type VarDef interface {
    getName() string
    assign(string) error
}


type Parser struct {
    definitions []VarDef
    positionals []string
    passdown_args []string
}

func (p *Parser) fromName(longname string) *VarDef {
    for _, vdef := range p.definitions {
        if vdef.getName() == longname { return &vdef }
    }
    return nil
}

func (p *Parser) ParseCliArgs(ignore_unknown bool) error {
    return p.Parse(os.Args[1:], ignore_unknown)
}

func (p *Parser) Parse(args []string, ignore_unknown bool) error {
    for i := 0; i<len(args); i++ {
        token := args[i]
        var def_p *VarDef = nil

        if token[:2] == "--" {
            longname := token[2:]
            if longname == "" {
                // We found the first delimiter for passdown arguments
                // Retain them as such, and stop parsing
                p.passdown_args = args[i+1:]
                return nil
            }
            def_p = p.fromName(longname)

            if def_p == nil && !ignore_unknown {
                return fmt.Errorf("Unknown flag %s", token)
            }

        }

        // TODO - support short names?

        if def_p != nil {
            def := *def_p
            switch def.(type) {
                case BoolDef:
                    def.(BoolDef).activate() // switches a boolean to opposite of its default value
                default:
                    i++
                    if i >= len(args) { return fmt.Errorf("Expected value after %s", token) }
                    def.assign(args[i])
            }

        } else {
            p.positionals = append(p.positionals, token)
        }
    }

    return nil
}

