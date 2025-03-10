package goargs

import (
	"fmt"
    "slices"
)


type CountDef struct {
    name string
    value *int
    helpstr string
}
func (self CountDef) getHelpString() string { return self.helpstr }
func (self CountDef) getName() string { return self.name }
func (self CountDef) assign(value string) error { panic("Invalid call to assign() on CountDef") }
func (self CountDef) increment() { *self.value++ }

func (p *Parser) CountVar(value *int, name string, helpstr string) {
    vdef := CountDef{name, value, helpstr}
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Count(name string, helpstr string) *int {
    var val int = 0
    p.CountVar(&val, name, helpstr)
    return &val
}

// =======

type ChoicesDef struct {
    name string
    value *string
    helpstr string
    choices []string
}
func (self ChoicesDef) getHelpString() string { return self.helpstr }
func (self ChoicesDef) getName() string { return self.name }
func (self ChoicesDef) assign(value string) error {
    if !slices.Contains(self.choices, value) {
        return fmt.Errorf("Invalid choice '%s'. Valid choices: %v", value, self.choices)
    }
    *self.value = value
    return nil
}

func (p *Parser) ChoicesVar(value *string, name string, choices []string, helpstr string) {
    vdef := ChoicesDef{name, value, helpstr, choices}
    *vdef.value = choices[0]
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Choices(name string, choices []string, helpstr string) *string {
    var val string
    p.ChoicesVar(&val, name, choices, helpstr)
    return &val
}

// =======

type AppenderDef struct {
    name string
    value *[]string
    helpstr string
}
func (self AppenderDef) getHelpString() string { return self.helpstr }
func (self AppenderDef) getName() string { return self.name }
func (self AppenderDef) assign(value string) error {
    *self.value = append(*self.value, value)
    return nil
}

func (p *Parser) AppenderVar(value *[]string, name string, helpstr string) {
    vdef := AppenderDef{name, value, helpstr}
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Appender(name string, helpstr string) *[]string {
    var val []string
    p.AppenderVar(&val, name, helpstr)
    return &val
}

// =======

type FuncDef struct {
    name string
    helpstr string
    innerfunc func(string) error
}
func (self FuncDef) getHelpString() string { return self.helpstr }
func (self FuncDef) getName() string { return self.name }
func (self FuncDef) assign(value string) error { panic("Invalid call to assign() on FuncDef") }
func (self FuncDef) call(s string) error {return self.innerfunc(s) }

func (p *Parser) Func(name string, funcdef func(string) error, helpstr string) {
    vdef := FuncDef{name, helpstr, funcdef}
    p.definitions[name] = vdef
    p.enqueueName(name)
}
