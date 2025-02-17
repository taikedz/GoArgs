package goargs

import (
    "fmt"
    "strconv"
)


// ======

type StringDef struct {
    name string
    value *string
}
func (self StringDef) getName() string { return self.name }
func (self StringDef) assign(value string) error { *self.value = value; return nil }

func (p *Parser) StringArg(value *string, name string, defval string) {
    s := StringDef{name, value}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}


// ======

type IntDef struct {
    name string
    value *int
}
func (self IntDef) getName() string { return self.name }
func (self IntDef) assign(value string) error {
    if val, err := strconv.Atoi(value); err != nil {
        return fmt.Errorf("Could not parse %s\n", value)
    } else {
        *self.value = val
    }
    return nil
}

func (p *Parser) IntArg(value *int, name string, defval int) {
    s := IntDef{name, value}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}


// ======

type FloatDef struct {
    name string
    value *float32
}
func (self FloatDef) getName() string { return self.name }
func (self FloatDef) assign(value string) error {
    if val, err := strconv.ParseFloat(value, 32); err != nil {
        return fmt.Errorf("Could not parse %s\n", value)
    } else {
        *self.value = float32(val)
    }
    return nil
}

func (p *Parser) FloatArg(value *float32, name string, defval float32) {
    s := FloatDef{name, value}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}


// ======

type BoolDef struct {
    name string
    value *bool
    defval bool
}
func (self BoolDef) getName() string { return self.name }
func (self BoolDef) assign(value string) error { return fmt.Errorf("Invalid operation") }
func (self BoolDef) activate() {
    *self.value = !self.defval
}

func (p *Parser) BoolArg(value *bool, name string, defval bool) {
    s := BoolDef{name, value, defval}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}

