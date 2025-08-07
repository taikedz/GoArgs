package goargs

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func (p *Parser) runeFromLong(name string) (rune, error) {
	for char, def := range p.shortnames {
		if name == def.getName() {
			return char, nil
		}
	}
	return '-', fmt.Errorf("no short flag found for '%s'", name)
}

func typeName(def t_VarDef) string {
	tokens := strings.Split(fmt.Sprintf("%s", reflect.TypeOf(def)), ".")
	typename := tokens[len(tokens)-1]
	return strings.ToUpper(typename[4:])
}

// Produce help text string and return it.
// Panics if an unknown type is unimplemented (goargs developer error. please report it!)
func (p *Parser) SPrintHelp() string {
	// return a string of formatted help information
	helplines := []string{p.helptext, ""}
	for _, name := range p.longnames {
		def := p.definitions[name]

		// Flag format
		switch def.(type) {
		case def_Bool, def_Count:
			helplines = append(helplines, fmt.Sprintf("  --%s", name))
			if sflag, err := p.runeFromLong(name); err == nil {
				helplines = append(helplines, fmt.Sprintf("  -%c", sflag))
			}
		default:
			var tname string
			switch def.(type) {
			case def_Choices, def_Appender, def_Func, def_Mode:
				tname = "STRING"
			default:
				tname = typeName(def)
			}
			helplines = append(helplines, fmt.Sprintf("  --%s %s", name, tname))
			switch def.(type) {
			case def_Mode:
				helplines = append(helplines, "    (use short flag or STRING value)")
			default:
				if sflag, err := p.runeFromLong(name); err == nil {
					helplines = append(helplines, fmt.Sprintf("  -%c %s", sflag, tname))
				}
			}
		}

		// Flag default value
		switch def.(type) {
		case def_String:
			helplines = append(helplines, fmt.Sprintf("    default: %s", def.(def_String).defval))
		case def_Choices:
			helplines = append(helplines, fmt.Sprintf("    default: %s", def.(def_Choices).choices[0]))
			helplines = append(helplines, fmt.Sprintf("    choices: %s", strings.Join(def.(def_Choices).choices, ", ")))
		case def_Int, def_Int64, def_Uint:
			helplines = append(helplines, fmt.Sprintf("    default: %d", def.(def_Int).defval))
		case def_Float, def_Float64:
			helplines = append(helplines, fmt.Sprintf("    default: %f", def.(def_Float).defval))
		case def_Bool:
			helplines = append(helplines, fmt.Sprintf("    default: %t", def.(def_Bool).defval))
		case def_Duration: // DurationDef, UnmarshalerDef
			helplines = append(helplines, fmt.Sprintf("    default: %v", def.(def_Duration).defval))
		case def_Count:
			helplines = append(helplines, fmt.Sprintf("    (each appearance is counted)"))
		case def_Appender:
			helplines = append(helplines, fmt.Sprintf("    (can be specified multiple times)"))
		case def_Mode:
			helplines = append(helplines, fmt.Sprintf("    default: %s", def.(def_Mode).defval))
		case def_Func:
			// do nothing. the user help will explain all.
		default:
			panic(fmt.Sprintf("Internal error (goargs): Uncatered type '%t'", def))
		}

		// Flag help string
		// TODO - wrap on terminal width splitting at spaces
		helplines = append(helplines, fmt.Sprintf("    %s", def.getHelpString()))
	}

	if len(p.post_helptext) > 0 {
		helplines = append(helplines, p.post_helptext)
	}

	return strings.Join(helplines, "\n")
}

/* Set the text to print at the end of the help message, after the parameters have been listed.
 */
func (p *Parser) SetPostHelptext(text string) {
	p.post_helptext = text
}

// Print the help message to stdout, uses SPrintHelp(), can panic
func (p *Parser) PrintHelp() {
	fmt.Println(p.SPrintHelp())
}

// Print the help message to stderr, uses SPrintHelp(), can panic
func (p *Parser) PrintHelpE() {
	print(p.SPrintHelp())
	println("")
}

// Identify the index of a token matching "-h" or "--help"
// e.g. `if FindHelpFlag(nil) >= 0 { ... printHelp() ... }`
func FindHelpFlag(tokens []string) int {
	if tokens == nil {
		tokens = os.Args[1:]
	}
	for i, token := range tokens {
		if token == "--help" || token == "-h" {
			return i
		} else if token == "--" {
			break
		}
	}
	return -1
}
