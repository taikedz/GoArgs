package goargs

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

const _VALID_SFLAGS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

type t_VarDef interface {
	getName() string
	assign(string) error
	getHelpString() string
}

// A discrete Parser to hold a number of argument definitions.
// Each Parser can receive a distinct set of value type pointers
// and be made to parse different sequences of argument tokens.
type Parser struct {
	definitions      map[string]t_VarDef
	shortnames       map[rune]t_VarDef
	longnames        []string
	helptext         string
	post_helptext    string
	require_flagdefs bool
	// Non-flag tokens in the arguments
	positionals []string
	// All tokens found after the first instance of `--`
	passdown_args []string
}

/*
Create a new parser instance, with initial help text.
Help text is printed before the flags' individual help strings are printed.
*/
func NewParser(helptext string) Parser {
	var p Parser
	p.definitions = make(map[string]t_VarDef)
	p.shortnames = make(map[rune]t_VarDef)
	p.helptext = helptext
	p.require_flagdefs = true
	return p
}

// Determine whether flags need to be defined. If false, treat unrecognised flags as
// basic string arguments.
func (p *Parser) RequireFlagDefs(require bool) {
	p.require_flagdefs = require
}

// register a flag in the parser
func (p *Parser) enqueueName(name string) {
	if slices.Contains(p.longnames, name) {
		panic(fmt.Sprintf("Flag '--%s' already defined.", name))
	}
	if matched, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_-]+$", name); !matched {
		panic(fmt.Sprintf("Invalid flag name '%s'. Must be minimum two characters long and start with letter", name))
	}
	p.longnames = append(p.longnames, name)
}

/*
Set a single-character notation for an existing long flag.
Panics if the code attempts to set a short flag rune that already exists,
or if the long flag is not yet registered,
or if the rune value is not alpha-numeric.
*/
func (p *Parser) SetShortFlag(short rune, longname string) {
	if !strings.ContainsRune(_VALID_SFLAGS, short) {
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

// Args returns the positional tokens from the parsed arguments.
// Args does not return pass-down arguments (after the first "--" token).
// See ExtraArgs() for passdown-arguments
func (p *Parser) Args() []string {
	return p.positionals[:]
}

// Arg(i) returns the i'th positional argument, after flags have been processed.
// Arg(i) cannot be used to access pass-down arguments.
// Arg(i) returns an error if the index is out of bounds
func (p *Parser) Arg(i int) (string, error) {
	if i < len(p.positionals) {
		return p.positionals[i], nil
	} else {
		return "", fmt.Errorf("Could not get item %d from a %d-length list", i, len(p.positionals))
	}
}

// ExtraArgs returns the list of tokens found after the first "--", if any.
// If none are found, the returned slice is empty.
func (p *Parser) ExtraArgs() []string {
	return p.passdown_args[:]
}

// Unpack positional arguments starting at index position, into specified pointer locations
// and return all remaining tokens after consumed tokens.
//
// e.g. `parser.UnpackArgs(0, &name1, &name2)` is like `parser.UnpackArgs(0, &name1); parser.UnpackArgs(1, &name2)`
func (p *Parser) UnpackArgs(idx int, ref ...interface{}) ([]string, error) {
	remains, err := Unpack(p.positionals[idx:], ref...)
	if err != nil {
		return nil, err
	}
	return remains, nil
}

func (p *Parser) clearParsedData() {
	p.positionals = []string{}
	p.passdown_args = []string{}
}

/*
Parse custom token sequence.

* If flag definitions are required (default), returns an error for unrecognised flags
* Else, unrecognised flags are stored unparsed in the positional arguments
* See `RequireFlagDefs(bool)`
*/
func (p *Parser) Parse(args []string) error {
	args, passdowns := splitTokensBefore("--", args)
	p.passdown_args = passdowns

	// CONFESSION : I don't like that this function is so convoluted.

	for i := 0; i < len(args); i++ {
		token := args[i]
		var def_ifc t_VarDef = nil // Interface types are a bit pointery (can be nil), but cannot ever be indirected with `*`
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
				return fmt.Errorf("unknown flag %s", token)
			}

		} else if len(token) > 1 && token[:1] == "-" {
			// Typically do not retain short flag aggregates
			// However if short flag is not found, retain the lot
			retain_token = false
			for _, sflag := range token[1:] {
				def, found_sflag := p.shortnames[sflag]
				if !found_sflag && p.require_flagdefs {
					return fmt.Errorf("unknown short flag '%c'", sflag)
				} else if !found_sflag {
					retain_token = true
					break
				}
				switch def.(type) {
				case def_Bool:
					def.(def_Bool).activate()
					continue
				case def_Count:
					def.(def_Count).increment()
					continue
				case def_Mode:
					def.(def_Mode).setShortMode(sflag)
					continue
				default:
					if len(token) == 2 {
						def_ifc = def
					} else {
						return fmt.Errorf("please specify '%c' on its own as '-%c VALUE'", sflag, sflag)
					}
				}
			}
		}

		if def_ifc != nil {
			switch def_ifc.(type) {
			case def_Bool:
				def_ifc.(def_Bool).activate()
			case def_Count:
				def_ifc.(def_Count).increment()
			default:
				if nextVal == nil {
					i++
					if i >= len(args) {
						return fmt.Errorf("expected value after %s", token)
					}
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

/*
Parse the program's CLI arguments. Must be called before accessing flags' variables.
See Parse() for further behaviours.
*/
func (p *Parser) ParseCliArgs() error {
	return p.Parse(os.Args[1:])
}
