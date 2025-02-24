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
func (self CountDef) assign(value string) error { return fmt.Errorf("Internal error: Invalid aciton. Use increment()") }
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