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
    value *string
}
func (self StringDef) getName() string { return self.name }
func (self StringDef) assign(value string) error { *self.value = value; return nil }

func (p *Parser) StringVar(value *string, name string, defval string) {
    vdef := StringDef{name, value}
    *vdef.value = defval
    p.definitions[name] = vdef
}
func (p *Parser) String(name string, defval string) *string {
    var s string
    p.StringVar(&s, name, defval)
    return &s
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

func (p *Parser) IntVar(value *int, name string, defval int) {
    vdef := IntDef{name, value}
    *vdef.value = defval
    p.definitions[name] = vdef
}
func (p *Parser) Int(name string, defval int) *int {
    var s int
    p.IntVar(&s, name, defval)
    return &s
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

func (p *Parser) FloatVar(value *float32, name string, defval float32) {
    vdef := FloatDef{name, value}
    *vdef.value = defval
    p.definitions[name] = vdef
}
func (p *Parser) Float(name string, defval float32) *float32 {
    var s float32
    p.FloatVar(&s, name, defval)
    return &s
}


// ======

type BoolDef struct {
    name string
    value *bool
    defval bool
}
func (self BoolDef) getName() string { return self.name }
func (self BoolDef) assign(value string) error { return fmt.Errorf("Invalid operation") }
func (self BoolDef) activate() { *self.value = !self.defval }

func (p *Parser) BoolVar(value *bool, name string, defval bool) {
    vdef := BoolDef{name, value, defval}
    *vdef.value = defval
    p.definitions[name] = vdef
}
func (p *Parser) Bool(name string, defval bool) *bool {
    var s bool
    p.BoolVar(&s, name, defval)
    return &s
}

// ======

type DurationDef struct {
    name string
    value *time.Duration
}

func (self DurationDef) getName() string { return self.name }
func (self DurationDef) assign(value string) error {
    if duration, err := time.ParseDuration(value); err != nil {
        return err
    } else {
        *self.value = duration
        return nil
    }
}

func (p *Parser) DurationVar(value *time.Duration, name string, defval time.Duration) {
    vdef := DurationDef{name, value}
    *vdef.value = defval
    p.definitions[name] = vdef
}
func (p *Parser) Duration(name string, defval time.Duration) *time.Duration {
    var s time.Duration
    p.DurationVar(&s, name, defval)
    return &s
}

