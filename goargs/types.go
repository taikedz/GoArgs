package goargs

import (
    "fmt"
    "time"
    "strconv"
)

// To add: float64, int64, uint, encoding.TextUnmarshaler

// ======

type StringDef struct {
    name string
    defval string
    value *string
    helpstr string
}
func (self StringDef) getHelpString() string { return self.helpstr }
func (self StringDef) getName() string { return self.name }
func (self StringDef) assign(value string) error { *self.value = value; return nil }

func (p *Parser) StringVar(value *string, name string, defval string, helpstr string) {
    vdef := StringDef{name, defval, value, helpstr}
    *vdef.value = defval
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) String(name string, defval string, helpstr string) *string {
    var s string
    p.StringVar(&s, name, defval, helpstr)
    return &s
}


// ======

type IntDef struct {
    name string
    defval int
    value *int
    helpstr string
}
func (self IntDef) getHelpString() string { return self.helpstr }
func (self IntDef) getName() string { return self.name }
func (self IntDef) assign(value string) error {
    if val, err := strconv.Atoi(value); err != nil {
        return fmt.Errorf("Could not parse %s\n", value)
    } else {
        *self.value = val
    }
    return nil
}

func (p *Parser) IntVar(value *int, name string, defval int, helpstr string) {
    vdef := IntDef{name, defval, value, helpstr}
    *vdef.value = defval
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Int(name string, defval int, helpstr string) *int {
    var s int
    p.IntVar(&s, name, defval, helpstr)
    return &s
}


// ======

type FloatDef struct {
    name string
    defval float32
    value *float32
    helpstr string
}
func (self FloatDef) getHelpString() string { return self.helpstr }
func (self FloatDef) getName() string { return self.name }
func (self FloatDef) assign(value string) error {
    if val, err := strconv.ParseFloat(value, 32); err != nil {
        return fmt.Errorf("Could not parse %s\n", value)
    } else {
        *self.value = float32(val)
    }
    return nil
}

func (p *Parser) FloatVar(value *float32, name string, defval float32, helpstr string) {
    vdef := FloatDef{name, defval, value, helpstr}
    *vdef.value = defval
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Float(name string, defval float32, helpstr string) *float32 {
    var s float32
    p.FloatVar(&s, name, defval, helpstr)
    return &s
}


// ======

type BoolDef struct {
    name string
    defval bool
    value *bool
    helpstr string
}
func (self BoolDef) getHelpString() string { return self.helpstr }
func (self BoolDef) getName() string { return self.name }
func (self BoolDef) assign(value string) error { return fmt.Errorf("Invalid operation") }
func (self BoolDef) activate() { *self.value = !self.defval }

func (p *Parser) BoolVar(value *bool, name string, defval bool, helpstr string) {
    vdef := BoolDef{name, defval, value, helpstr}
    *vdef.value = defval
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Bool(name string, defval bool, helpstr string) *bool {
    var s bool
    p.BoolVar(&s, name, defval, helpstr)
    return &s
}

// ======

type DurationDef struct {
    name string
    defval time.Duration
    value *time.Duration
    helpstr string
}
func (self DurationDef) getHelpString() string { return self.helpstr }
func (self DurationDef) getName() string { return self.name }
func (self DurationDef) assign(value string) error {
    if duration, err := time.ParseDuration(value); err != nil {
        return err
    } else {
        *self.value = duration
        return nil
    }
}

func (p *Parser) DurationVar(value *time.Duration, name string, defval time.Duration, helpstr string) {
    vdef := DurationDef{name, defval, value, helpstr}
    *vdef.value = defval
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Duration(name string, defval time.Duration, helpstr string) *time.Duration {
    var s time.Duration
    p.DurationVar(&s, name, defval, helpstr)
    return &s
}

