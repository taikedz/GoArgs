package goargs

import (
	"fmt"
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