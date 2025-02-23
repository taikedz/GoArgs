package goargs


import (
	"fmt"
	"strings"
)

func (p *Parser) SPrintHelp() string {
	// return a string of formatted help information
	helplines := []string{p.helptext, ""}
	for _,name := range p.longnames {
		def := p.definitions[name]

		// Flag format
		switch def.(type) {
		case BoolDef, CountDef:
			helplines = append(helplines, fmt.Sprintf("  --%s", name)) // TODO - list short flag
		default:
			helplines = append(helplines, fmt.Sprintf("  --%s VALUE", name))
		}

		// Flag default value
		switch def.(type) {
		case StringDef:
			helplines = append(helplines, fmt.Sprintf("    default: %s", def.(StringDef).defval))
		case IntDef: // Eventually "case IntDef, UintDef:"
			helplines = append(helplines, fmt.Sprintf("    default: %d", def.(IntDef).defval))
		case FloatDef: // FloatDef, Float64Def
			helplines = append(helplines, fmt.Sprintf("    default: %f", def.(FloatDef).defval))
		case BoolDef:
			helplines = append(helplines, fmt.Sprintf("    default: %t", def.(BoolDef).defval))
		case DurationDef: // DurationDef, UnmarshalerDef
			helplines = append(helplines, fmt.Sprintf("    default: %v", def.(DurationDef).defval))
		case CountDef:
			helplines = append(helplines, fmt.Sprintf("(each appearance of '--%s' is counted)", name))
		}

		// Flag help string
		// TODO - wrap on terminal width splitting at spaces
		helplines = append(helplines, fmt.Sprintf("    %s", def.getHelpString()))
	}

	return strings.Join(helplines, "\n")
}

func (p *Parser) PrintHelp() {
	print(p.SPrintHelp())
}

func FindHelpFlag(tokens []string) int {
	for i,token := range tokens {
		if token == "--help" || token == "-h" {
			return i
		} else if token == "--" {
			break
		}
	}
	return -1
}