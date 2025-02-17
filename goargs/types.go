package goargs

import (
    "fmt"
    "time"
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
func (p *Parser) StringVar(name string, defval string) *string {
    var s string
    p.StringArg(&s, name, defval)
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

func (p *Parser) IntArg(value *int, name string, defval int) {
    s := IntDef{name, value}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}
func (p *Parser) IntVar(name string, defval int) *int {
    var s int
    p.IntArg(&s, name, defval)
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

func (p *Parser) FloatArg(value *float32, name string, defval float32) {
    s := FloatDef{name, value}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}
func (p *Parser) FloatVar(name string, defval float32) *float32 {
    var s float32
    p.FloatArg(&s, name, defval)
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

func (p *Parser) BoolArg(value *bool, name string, defval bool) {
    s := BoolDef{name, value, defval}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}
func (p *Parser) BoolVar(name string, defval bool) *bool {
    var s bool
    p.BoolArg(&s, name, defval)
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

func (p *Parser) DurationArg(value *time.Duration, name string, defval time.Duration) {
    s := DurationDef{name, value}
    *s.value = defval
    p.definitions = append(p.definitions, s)
}
func (p *Parser) DurationVar(name string, defval time.Duration) *time.Duration {
    var s time.Duration
    p.DurationArg(&s, name, defval)
    return &s
}

