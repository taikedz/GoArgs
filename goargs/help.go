package goargs


import (
	"fmt"
	"strings"
	"reflect"
)

func (p *Parser) runeFromLong(name string) (rune, error) {
	for char, def := range p.shortnames {
		if name == def.getName() {
			return char, nil
		}
	}
	return '-', fmt.Errorf("No short flag found for '%s'", name)
}

func typeName(def VarDef) string {
	tokens := strings.Split(fmt.Sprintf("%s", reflect.TypeOf(def)), ".")
	typename := tokens[len(tokens)-1]
	return strings.ToUpper(typename[:len(typename)-3])
}

func (p *Parser) SPrintHelp() string {
	// return a string of formatted help information
	helplines := []string{p.helptext, ""}
	for _,name := range p.longnames {
		def := p.definitions[name]

		// Flag format
		switch def.(type) {
		case BoolDef, CountDef:
			helplines = append(helplines, fmt.Sprintf("  --%s", name))
			if sflag, err := p.runeFromLong(name); err == nil {
				helplines = append(helplines, fmt.Sprintf("  -%c", sflag))
			}
		default:
            var tname string
            switch def.(type) {
            case ChoicesDef, AppenderDef, FuncDef, ModeDef:
                tname = "STRING"
            default:
                tname = typeName(def)
            }
			helplines = append(helplines, fmt.Sprintf("  --%s %s", name, tname))
			switch def.(type) {
			case ModeDef:
				// do nothing - this type registers multiple short flags to the same name.
			default:
				if sflag, err := p.runeFromLong(name); err == nil {
					helplines = append(helplines, fmt.Sprintf("  -%c %s", sflag, tname))
				}
			}
		}

		// Flag default value
		switch def.(type) {
		case StringDef:
			helplines = append(helplines, fmt.Sprintf("    default: %s", def.(StringDef).defval))
		case ChoicesDef:
			helplines = append(helplines, fmt.Sprintf("    default: %s", def.(ChoicesDef).choices[0]))
			helplines = append(helplines, fmt.Sprintf("    choices: %s", strings.Join(def.(ChoicesDef).choices, ", ") ))
		case IntDef, Int64Def, UintDef:
			helplines = append(helplines, fmt.Sprintf("    default: %d", def.(IntDef).defval))
		case FloatDef, Float64Def:
			helplines = append(helplines, fmt.Sprintf("    default: %f", def.(FloatDef).defval))
		case BoolDef:
			helplines = append(helplines, fmt.Sprintf("    default: %t", def.(BoolDef).defval))
		case DurationDef: // DurationDef, UnmarshalerDef
			helplines = append(helplines, fmt.Sprintf("    default: %v", def.(DurationDef).defval))
		case CountDef:
			helplines = append(helplines, fmt.Sprintf("    (each appearance is counted)"))
		case AppenderDef:
			helplines = append(helplines, fmt.Sprintf("    (can be specified multiple times)"))
		case ModeDef:
			helplines = append(helplines, fmt.Sprintf("    default: %s", def.(ModeDef).defval))
        case FuncDef:
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

// Print the help message to stderr
func (p *Parser) PrintHelp() {
	print(p.SPrintHelp())
    println("")
}

// Identify the index of a token matching "-h" or "--help"
func findHelpFlag(tokens []string) int {
	for i,token := range tokens {
		if token == "--help" || token == "-h" {
			return i
		} else if token == "--" {
			break
		}
	}
	return -1
}
