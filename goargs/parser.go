package goargs

import (
    "fmt"
    "os"
    "strings"
    "slices"
    "regexp"
)

const _VALID_SFLAGS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

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
    shortnames map[rune]VarDef
    longnames []string
    helptext string
    post_helptext string
    help_on_empty_args bool
    require_flagdefs bool
    // Non-flag tokens in the arguments
    positionals []string
    // All tokens found after the first instance of `--`
    passdown_args []string
}

func NewParser(helptext string) Parser {
    var p Parser
    p.definitions = make(map[string]VarDef)
    p.shortnames = make(map[rune]VarDef)
    p.helptext = helptext
    p.require_flagdefs = true
    return p
}

func (p *Parser) SetHelpOnEmptyArgs(onempty bool) {
    p.help_on_empty_args = onempty
}

func (p *Parser) RequireFlagDefs(require bool) {
    p.require_flagdefs = require
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

func (p *Parser) SetShortFlag(short rune, longname string) {
    if ! strings.ContainsRune(_VALID_SFLAGS, short) {
        panic(fmt.Sprintf("Internal error: cannot use rune %c", short))
    }
    if gotname, ok := p.shortnames[short]; ok {
        panic(fmt.Sprintf("'-%c' already defined against '%s'", short, gotname))
    }
    def, ok := p.definitions[longname]
    if !ok {
        panic(fmt.Sprintf("Flag '--%s' not yet defined", longname))
    }
    p.shortnames[short] = def
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

func (p *Parser) ExtraArgs() []string {
    return p.passdown_args[:]
}

/// Unpack positional arguments starting at index position, into specified pointer locations
//   amd return all remaining tokens after consumed tokens.
// e.g. `parser.UnpackArgs(0, &name1, &name2)` is like `parser.UnpackArgs(0, &name 1); parser.UnpackArgs(1, &name2)`
func (p *Parser) UnpackArgs(idx int, ref ... interface{}) ([]string, error) {
    remains, err := Unpack(p.positionals[idx:], ref...)
    if err != nil {
        return nil, err
    }
    return remains, nil
}

func (p *Parser) ClearParsedData() {
    p.positionals = []string{}
    p.passdown_args = []string{}
}

/*
Parse custom token sequence.

* If flag definitions are required (default), returns an error for unrecognised flags
* Else, retain_tokens unrecognised flags in the positional arguments set
* See `RequireFlagDefs(bool)`

If "-h" or "--help" is found before the first "--", then help is printed and process exits.
*/
func (p *Parser) Parse(args []string) error {
    p.autoHelp(args)
    args, passdowns := SplitTokensBefore("--", args)
    p.passdown_args = passdowns

    // CONFESSION : I don't like that this function is so convoluted.

    for i := 0; i<len(args); i++ {
        token := args[i]
        var def_ifc VarDef = nil // Interface types are a bit pointery (can be nil), but cannot ever be indirected with `*`
        var nextVal *string = nil
        var retain_token = true

        if len(token) >= 2 && token[:2] == "--" {
            longname := token[2:]
            if strings.Contains(longname, "=") {
                seq := strings.SplitN(longname, "=", 2)
                longname = seq[0]
                nextVal = &seq[1]
            }
            def_ifc = p.definitions[longname]

            if def_ifc == nil && p.require_flagdefs {
                return fmt.Errorf("Unknown flag %s", token)
            }

        } else if len(token) > 1 && token[:1] == "-" {
            // Typically do not retain short flag aggregates
            // However if short flag is not found, retain the lot
            retain_token = false
            for _, sflag := range token[1:] {
                def, found_sflag := p.shortnames[sflag]
                if !found_sflag && p.require_flagdefs {
                    return fmt.Errorf("Unknown short flag '%c'", sflag)
                } else if !found_sflag {
                    retain_token = true
                    break
                }
                switch def.(type) {
                case BoolDef:
                    def.(BoolDef).activate()
                    continue
                case CountDef:
                    def.(CountDef).increment()
                    continue
                case ModeDef:
                    def.(ModeDef).setShortMode(sflag)
                    continue
                default:
                    if len(token) == 2 {
                        def_ifc = def
                    } else {
                        return fmt.Errorf("Please specify '%c' on its own as '-%c VALUE'", sflag, sflag)
                    }
                }
            }
        }

        if def_ifc != nil {
            switch def_ifc.(type) {
                case BoolDef:
                    def_ifc.(BoolDef).activate()
                case CountDef:
                    def_ifc.(CountDef).increment()
                default:
                    if nextVal == nil {
                        i++
                        if i >= len(args) { return fmt.Errorf("Expected value after %s", token) }
                        nextVal = &args[i]
                    }
                    switch def_ifc.(type) {
                    // case FuncDef:
                    //     def_ifc.(FuncDef).call(*nextVal)
                    default:
                        if err := def_ifc.assign(*nextVal); err != nil {
                            return err
                        }
                    }
            }

        } else {
            if retain_token {
                p.positionals = append(p.positionals, token)
            }
        }
    }

    return nil
}

/* Parse the program's CLI arguments. Must be called before accessing flags' variables.

See Parse() for further behaviours.
*/
func (p *Parser) ParseCliArgs() error {
    return p.Parse(os.Args[1:])
}

func (p *Parser) autoHelp(args []string) {
    if i := FindHelpFlag(args); (p.help_on_empty_args && len(os.Args) <= 1) || i >= 0 {
        p.PrintHelp()
        os.Exit(0)
    }
}
