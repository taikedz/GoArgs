package goargs

import (
    "fmt"
    "time"
    "strconv"
)

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
    var val string
    p.StringVar(&val, name, defval, helpstr)
    return &val
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
    var val int
    p.IntVar(&val, name, defval, helpstr)
    return &val
}


// ======

type Int64Def struct {
    name string
    defval int64
    value *int64
    helpstr string
}
func (self Int64Def) getHelpString() string { return self.helpstr }
func (self Int64Def) getName() string { return self.name }
func (self Int64Def) assign(value string) error {
    if val, err := strconv.Atoi(value); err != nil {
        return fmt.Errorf("Could not parse %s\n", value)
    } else {
        *self.value = int64(val)
    }
    return nil
}

func (p *Parser) Int64Var(value *int64, name string, defval int64, helpstr string) {
    vdef := Int64Def{name, defval, value, helpstr}
    *vdef.value = defval
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Int64(name string, defval int64, helpstr string) *int64 {
    var val int64
    p.Int64Var(&val, name, defval, helpstr)
    return &val
}


// ======

type UintDef struct {
    name string
    defval uint
    value *uint
    helpstr string
}
func (self UintDef) getHelpString() string { return self.helpstr }
func (self UintDef) getName() string { return self.name }
func (self UintDef) assign(value string) error {
    if val, err := strconv.Atoi(value); err != nil {
        return fmt.Errorf("Could not parse %s\n", value)
    } else {
        *self.value = uint(val)
    }
    return nil
}

func (p *Parser) UintVar(value *uint, name string, defval uint, helpstr string) {
    vdef := UintDef{name, defval, value, helpstr}
    *vdef.value = defval
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Uint(name string, defval uint, helpstr string) *uint {
    var val uint
    p.UintVar(&val, name, defval, helpstr)
    return &val
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
    var val float32
    p.FloatVar(&val, name, defval, helpstr)
    return &val
}


// ======

type Float64Def struct {
    name string
    defval float64
    value *float64
    helpstr string
}
func (self Float64Def) getHelpString() string { return self.helpstr }
func (self Float64Def) getName() string { return self.name }
func (self Float64Def) assign(value string) error {
    if val, err := strconv.ParseFloat(value, 64); err != nil {
        return fmt.Errorf("Could not parse %s\n", value)
    } else {
        *self.value = float64(val)
    }
    return nil
}

func (p *Parser) Float64Var(value *float64, name string, defval float64, helpstr string) {
    vdef := Float64Def{name, defval, value, helpstr}
    *vdef.value = defval
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Float64(name string, defval float64, helpstr string) *float64 {
    var val float64
    p.Float64Var(&val, name, defval, helpstr)
    return &val
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
func (self BoolDef) assign(value string) error { panic("Invalid call to assign() on BoolDef") }
func (self BoolDef) activate() { *self.value = !self.defval }

func (p *Parser) BoolVar(value *bool, name string, defval bool, helpstr string) {
    vdef := BoolDef{name, defval, value, helpstr}
    *vdef.value = defval
    p.definitions[name] = vdef
    p.enqueueName(name)
}
func (p *Parser) Bool(name string, defval bool, helpstr string) *bool {
    var val bool
    p.BoolVar(&val, name, defval, helpstr)
    return &val
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
    var val time.Duration
    p.DurationVar(&val, name, defval, helpstr)
    return &val
}

