package goargs

import (
    "fmt"
    "os"
    "strings"
    "slices"
    "regexp"
)

type VarDef interface {
    getName() string
    assign(string) error
    getHelpString() string
}


// A discrete Parser to hold a number of argument definitions.
// Each Parser can receive a distinct set of value type pointers
//  and be made to parse different sequences of argument tokens.
type Parser struct {
    definitions map[string]VarDef
    longnames []string
    helptext string
    // Non-flag tokens in the arguments
    positionals []string
    // All tokens found after the first instance of `--`
    PassdownArgs []string
}

func NewParser(helptext string) Parser {
    var p Parser
    p.definitions = make(map[string]VarDef)
    p.helptext = helptext
    return p
}

func (p *Parser) enqueueName(name string) {
    if slices.Contains(p.longnames, name) {
        panic(fmt.Sprintf("Flag '--%s' already defined.", name))
    }
    if matched, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_-]+$", name); !matched {
        panic(fmt.Sprintf("Invalid flag name '%s'. Must be minimum two characters long and start with letter"))
    }
    p.longnames = append(p.longnames, name)
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

func (p *Parser) ClearParsedData() {
    p.positionals = []string{}
    p.PassdownArgs = []string{}
}

/*
Parse custom token sequence.

* If ignore_unknown is false, returns an error for unrecognised flags
* If ignore_unknown is true, retains unrecognised flags in the positional arguments set

Panics if a flag is defined twice on the same name, or if the flag has a bad name.
Acceptable flag names must be at least two characters long, and start with an ASCII-7 alphabetical character.
*/
func (p *Parser) Parse(args []string, ignore_unknown bool) error {
    for i := 0; i<len(args); i++ {
        token := args[i]
        var def_ifc VarDef = nil // Interface types are a bit pointery (can be nil), but cannot ever be indirected with `*`
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
            def_ifc = p.definitions[longname]

            if def_ifc == nil && !ignore_unknown {
                return fmt.Errorf("Unknown flag %s", token)
            }

        }

        // TODO - support short names?

        if def_ifc != nil {
            switch def_ifc.(type) {
                case BoolDef:
                    // switches a boolean to opposite of its default value
                    def_ifc.(BoolDef).activate()
                case CountDef:
                    def_ifc.(CountDef).increment()
                default:
                    if nextVal == nil {
                        i++
                        if i >= len(args) { return fmt.Errorf("Expected value after %s", token) }
                        nextVal = &args[i]
                    }
                    if err := def_ifc.assign(*nextVal); err != nil {
                        return err
                    }
            }

        } else {
            p.positionals = append(p.positionals, token)
        }
    }

    return nil
}

/* Parse the program's CLI arguments. Must be called before accessing flags' variables.

If "-" or "--help" is found before the first "--", then help is printed.

See Parse() for further behaviours.
*/
func (p *Parser) ParseCliArgs(ignore_unknown bool) error {
    if i := FindHelpFlag(os.Args[1:]); i >= 0 {
        p.PrintHelp()
        os.Exit(0)
    }
    return p.Parse(os.Args[1:], ignore_unknown)
}

