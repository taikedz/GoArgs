package goargs

import (
    "fmt"
    "os"
    "strings"
)

type VarDef interface {
    getName() string
    assign(string) error
}


// A discrete Parser to hold a number of argument definitions.
// Each Parser can receive a distinct set of value type pointers
//  and be made to parse different sequences of argument tokens.
type Parser struct {
    definitions []VarDef
    // Non-flag tokens in the arguments
    positionals []string
    // All tokens found after the first instance of `--`
    PassdownArgs []string
}

// Args returns the positional tokens from the parsed arguments
// Args does not return pass-down arguments
func (p *Parser) Args() []string {
    return p.positionals[:]
}

// Arg returns the i'th positional argument, after flags have been processed.
// Arg does not process the pass-down arguments.
// Arg returns an error if the index is out of bounds
func (p *Parser) Arg(i int) (string, error) {
    if i < len(p.positionals) {
        return p.positionals[i], nil
    } else {
        return "", fmt.Errorf("Could not get item %d from a %d-length list", i, len(p.positionals))
    }
}

func (p *Parser) clearParsedData() {
    p.positionals = []string{}
    p.PassdownArgs = []string{}
}

// Look for a VarDef carying the given longname as its name, returning a pointer to that VarDef
// If none is found returns nil
func (p *Parser) fromName(longname string) *VarDef {
    for _, vdef := range p.definitions {
        if vdef.getName() == longname { return &vdef }
    }
    return nil
}

// Parse the program's CLI arguments. Must be called before accessing flags' variables.
// If ignore_unknown is true, returns an error for unrecognised flags
// If ignore_unknown is false, retains unrecognised flags in the positional arguments set
func (p *Parser) ParseCliArgs(ignore_unknown bool) error {
    return p.Parse(os.Args[1:], ignore_unknown)
}

// Parse custom token sequence. See ParseCliArgs.
func (p *Parser) Parse(args []string, ignore_unknown bool) error {
    for i := 0; i<len(args); i++ {
        token := args[i]
        var def_p *VarDef = nil
        var nextVal *string = nil

        if token[:2] == "--" {
            longname := token[2:]
            if longname == "" {
                // We found the first delimiter for passdown arguments
                // Retain them as such, and stop parsing
                p.PassdownArgs = args[i+1:]
                return nil
            }
            if strings.Contains(longname, "=") {
                seq := strings.SplitN(longname, "=", 2)
                longname = seq[0]
                nextVal = &seq[1]
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
                    if nextVal == nil {
                        i++
                        if i >= len(args) { return fmt.Errorf("Expected value after %s", token) }
                        nextVal = &args[i]
                    }
                    if err := def.assign(*nextVal); err != nil {
                        return err
                    }
            }

        } else {
            p.positionals = append(p.positionals, token)
        }
    }

    return nil
}

