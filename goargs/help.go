package goargs

/*
$ go run goargs/help.go --name Stan -h who am i
Usage of /tmp/go-build2333072002/b001/exe/help:
  -help string
    	Something (default "nothing")
  -name string
    	Someone (default "nobody")

*/


import (
	"fmt"
)

func (p *Parser) SPrintHelp() string {
	// return a string of formatted help information
	var helplines []string
	for name,def := range p.definitions {
		switch def.(type) { // Flag format
		case BoolDef:
			helplines = append(helplines, fmt.Sprintf("  --%s\n", name)) // TODO - list short flag
		case default:
			helplines = append(helplines, fmt.Sprintf("  --%s VALUE\n", name))
		}

		switch def.(type) { // Flag default value
		case StringDef:
			helplines = append(helplines, fmt.Sprintf("    default: %s\n", def.(StringDef).defval))
		case IntDef: // Eventually "case IntDef, UintDef:"
			helplines = append(helplines, fmt.Sprintf("    default: %d\n", def.(IntDef).defval))
		case FloatDef: // FloatDef, Float64Def
			helplines = append(helplines, fmt.Sprintf("    default: %f\n", def.(FloatDef).defval))
		case BoolDef:
			helplines = append(helplines, fmt.Sprintf("    default: %s\n", def.(BoolDef).defval))
		case DurationDef: // DurationDef, UnmarshalerDef
			helplines = append(helplines, fmt.Sprintf("    default: %v\n", def.(DurationDef).defval))
		}

		// Flag help string
		// TODO - wrap on terminal width splitting at spaces
		helplines = append(helplines, fmt.Sprintf("    %s\n", def.getHelpstring()))
	}
}

func PrintHelp() {
	print(SPrintHelp())
}