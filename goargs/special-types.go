package goargs

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type def_Count struct {
	name    string
	value   *int
	helpstr string
}

func (self def_Count) getHelpString() string { return self.helpstr }
func (self def_Count) getName() string       { return self.name }
func (self def_Count) assign(value string) error {
	panic("(goargs) Invalid call to assign() on CountDef")
}
func (self def_Count) increment() { *self.value++ }

// Register a Count argument, storing to the supplied `value *int` pointer
// A Count argument increments by 1 every time the flag is seen.
func (p *Parser) CountVar(value *int, name string, helpstr string) {
	vdef := def_Count{name, value, helpstr}
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register a Count argument, storing to the returned `*int` pointer
// A Count argument increments by 1 every time the flag is seen.
func (p *Parser) Count(name string, helpstr string) *int {
	var val int = 0
	p.CountVar(&val, name, helpstr)
	return &val
}

// =======

type def_Choices struct {
	name    string
	value   *string
	helpstr string
	choices []string
}

func (self def_Choices) getHelpString() string { return self.helpstr }
func (self def_Choices) getName() string       { return self.name }
func (self def_Choices) assign(value string) error {
	if !slices.Contains(self.choices, value) {
		return fmt.Errorf("Invalid choice '%s'. Valid choices: %v", value, self.choices)
	}
	*self.value = value
	return nil
}

// Register a Choices argument, storing to the supplied `value *string` pointer
// A Choices argument will only accept one of the specified `choices []string` elements
func (p *Parser) ChoicesVar(value *string, name string, choices []string, helpstr string) {
	vdef := def_Choices{name, value, helpstr, choices}
	*vdef.value = choices[0]
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register a Choices argument, storing to the returned `*string` pointer
// A Choices argument will only accept one of the specified `choices []string` elements
func (p *Parser) Choices(name string, choices []string, helpstr string) *string {
	var val string
	p.ChoicesVar(&val, name, choices, helpstr)
	return &val
}

// =======

type def_Appender struct {
	name    string
	value   *[]string
	helpstr string
}

func (self def_Appender) getHelpString() string { return self.helpstr }
func (self def_Appender) getName() string       { return self.name }
func (self def_Appender) assign(value string) error {
	*self.value = append(*self.value, value)
	return nil
}

// Register an Appender argument, storing to the supplied `value *[]string` pointer
// An Appender argument will append the associated value into the specified slice
func (p *Parser) AppenderVar(value *[]string, name string, helpstr string) {
	vdef := def_Appender{name, value, helpstr}
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register an Appender argument, storing to the returned `*[]string` pointer
// An Appender argument will append the associated value into the returned slice
func (p *Parser) Appender(name string, helpstr string) *[]string {
	var val []string
	p.AppenderVar(&val, name, helpstr)
	return &val
}

// =======

type def_Func struct {
	name      string
	helpstr   string
	innerfunc func(string) error
}

func (self def_Func) getHelpString() string     { return self.helpstr }
func (self def_Func) getName() string           { return self.name }
func (self def_Func) assign(value string) error { return self.innerfunc(value) }

// Register a Function argument
// The function defined at `funcdef` will be called each time the flag is seen, and be called
// with the associated value. If no value is needed, consider using Count instead
func (p *Parser) Func(name string, funcdef func(string) error, helpstr string) {
	vdef := def_Func{name, helpstr, funcdef}
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// =======

type def_Mode struct {
	name    string
	defval  string
	value   *string
	helpstr string
	modes   map[rune]string
}

func (self def_Mode) getHelpString() string { return self.helpstr }
func (self def_Mode) getName() string       { return self.name }
func (self def_Mode) assign(value string) error {
	// go through the modes map, and check that the mode value is found there
	var values []string
	for _, okval := range self.modes {
		if value == okval {
			*self.value = value
			return nil
		}
		values = append(values, okval)
	}
	return fmt.Errorf("Invalid mode '%s' - choose from: %s", value, strings.Join(values, ", "))
}
func (self def_Mode) setShortMode(short rune) {
	*self.value = self.modes[short]
}

/*
Register a Mode argument, storing the value in the specified `value *string` pointer
The stored value of a Mode argument corresponds to the last mode flag parsed.
Mode flags systematically use a map of runes to mode flag names

	// Example `modes` argument
	{
	    'b': "bright",
	    'd': "dark",
	    'm': "dim",
	}
	// allows for the short flags `-b`, `-d` and `-m` as well as `--bright`, `--dark`, and `--dim`
*/
func (p *Parser) ModeVar(value *string, name string, defval string, modes map[rune]string, helpstr string) {
	mode_helpstr := []string{}
	for k, v := range modes {
		mode_helpstr = append(mode_helpstr, fmt.Sprintf("      -%c : %s", k, v))
	}
	sort.Strings(mode_helpstr)
	mode_helpstr = slices.Insert(mode_helpstr, 0, helpstr)

	vdef := def_Mode{name, defval, value, strings.Join(mode_helpstr, "\n"), modes}
	*vdef.value = defval
	p.definitions[name] = vdef
	p.enqueueName(name)
	for r, _ := range modes {
		p.SetShortFlag(r, name)
	}
}

// Register a Mode argument, storing the value in the returned `*string` pointer
// See ModeVar for details
func (p *Parser) Mode(name string, defval string, modes map[rune]string, helpstr string) *string {
	var val string
	p.ModeVar(&val, name, defval, modes, helpstr)
	return &val
}
