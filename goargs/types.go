package goargs

import (
	"fmt"
	"strconv"
	"time"
)

// ======

type def_String struct {
	name    string
	defval  string
	value   *string
	helpstr string
}

func (self def_String) getHelpString() string     { return self.helpstr }
func (self def_String) getName() string           { return self.name }
func (self def_String) assign(value string) error { *self.value = value; return nil }

// Register a string argument, storing to the supplied `value *string` pointer
func (p *Parser) StringVar(value *string, name string, defval string, helpstr string) {
	vdef := def_String{name, defval, value, helpstr}
	*vdef.value = defval
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register a string argument, storing to the returned `*string` pointer
func (p *Parser) String(name string, defval string, helpstr string) *string {
	var val string
	p.StringVar(&val, name, defval, helpstr)
	return &val
}

// ======

type def_Int struct {
	name    string
	defval  int
	value   *int
	helpstr string
}

func (self def_Int) getHelpString() string { return self.helpstr }
func (self def_Int) getName() string       { return self.name }
func (self def_Int) assign(value string) error {
	if val, err := strconv.Atoi(value); err != nil {
		return fmt.Errorf("Could not parse %s\n", value)
	} else {
		*self.value = val
	}
	return nil
}

// Register an int argument, storing to the supplied `value *int` pointer
func (p *Parser) IntVar(value *int, name string, defval int, helpstr string) {
	vdef := def_Int{name, defval, value, helpstr}
	*vdef.value = defval
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register an int argument, storing to the returned `*int` pointer
func (p *Parser) Int(name string, defval int, helpstr string) *int {
	var val int
	p.IntVar(&val, name, defval, helpstr)
	return &val
}

// ======

type def_Int64 struct {
	name    string
	defval  int64
	value   *int64
	helpstr string
}

func (self def_Int64) getHelpString() string { return self.helpstr }
func (self def_Int64) getName() string       { return self.name }
func (self def_Int64) assign(value string) error {
	if val, err := strconv.Atoi(value); err != nil {
		return fmt.Errorf("Could not parse %s\n", value)
	} else {
		*self.value = int64(val)
	}
	return nil
}

// Register an int64 argument, storing to the supplied `value *int64` pointer
func (p *Parser) Int64Var(value *int64, name string, defval int64, helpstr string) {
	vdef := def_Int64{name, defval, value, helpstr}
	*vdef.value = defval
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register an int64 argument, storing to the returned `*int64` pointer
func (p *Parser) Int64(name string, defval int64, helpstr string) *int64 {
	var val int64
	p.Int64Var(&val, name, defval, helpstr)
	return &val
}

// ======

type def_Uint struct {
	name    string
	defval  uint
	value   *uint
	helpstr string
}

func (self def_Uint) getHelpString() string { return self.helpstr }
func (self def_Uint) getName() string       { return self.name }
func (self def_Uint) assign(value string) error {
	if val, err := strconv.Atoi(value); err != nil {
		return fmt.Errorf("Could not parse %s\n", value)
	} else {
		*self.value = uint(val)
	}
	return nil
}

// Register an uint argument, storing to the supplied `value *uint` pointer
func (p *Parser) UintVar(value *uint, name string, defval uint, helpstr string) {
	vdef := def_Uint{name, defval, value, helpstr}
	*vdef.value = defval
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register an uint argument, storing to the returned `*uint` pointer
func (p *Parser) Uint(name string, defval uint, helpstr string) *uint {
	var val uint
	p.UintVar(&val, name, defval, helpstr)
	return &val
}

// ======

type def_Float struct {
	name    string
	defval  float32
	value   *float32
	helpstr string
}

func (self def_Float) getHelpString() string { return self.helpstr }
func (self def_Float) getName() string       { return self.name }
func (self def_Float) assign(value string) error {
	if val, err := strconv.ParseFloat(value, 32); err != nil {
		return fmt.Errorf("Could not parse %s\n", value)
	} else {
		*self.value = float32(val)
	}
	return nil
}

// Register a float argument, storing to the supplied `value *float` pointer
func (p *Parser) FloatVar(value *float32, name string, defval float32, helpstr string) {
	vdef := def_Float{name, defval, value, helpstr}
	*vdef.value = defval
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register a float argument, storing to the returned `*float` pointer
func (p *Parser) Float(name string, defval float32, helpstr string) *float32 {
	var val float32
	p.FloatVar(&val, name, defval, helpstr)
	return &val
}

// ======

type def_Float64 struct {
	name    string
	defval  float64
	value   *float64
	helpstr string
}

func (self def_Float64) getHelpString() string { return self.helpstr }
func (self def_Float64) getName() string       { return self.name }
func (self def_Float64) assign(value string) error {
	if val, err := strconv.ParseFloat(value, 64); err != nil {
		return fmt.Errorf("Could not parse %s\n", value)
	} else {
		*self.value = float64(val)
	}
	return nil
}

// Register a float64 argument, storing to the supplied `value *float64` pointer
func (p *Parser) Float64Var(value *float64, name string, defval float64, helpstr string) {
	vdef := def_Float64{name, defval, value, helpstr}
	*vdef.value = defval
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register a float64 argument, storing to the returned `*float64` pointer
func (p *Parser) Float64(name string, defval float64, helpstr string) *float64 {
	var val float64
	p.Float64Var(&val, name, defval, helpstr)
	return &val
}

// ======

type def_Bool struct {
	name    string
	defval  bool
	value   *bool
	helpstr string
}

func (self def_Bool) getHelpString() string { return self.helpstr }
func (self def_Bool) getName() string       { return self.name }
func (self def_Bool) assign(value string) error {
	panic("(goargs) Invalid call to assign() on BoolDef")
}
func (self def_Bool) activate() { *self.value = !self.defval }

// Register a bool argument, storing to the supplied `value *bool` pointer
func (p *Parser) BoolVar(value *bool, name string, defval bool, helpstr string) {
	vdef := def_Bool{name, defval, value, helpstr}
	*vdef.value = defval
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register a bool argument, storing to the returned `*bool` pointer
func (p *Parser) Bool(name string, defval bool, helpstr string) *bool {
	var val bool
	p.BoolVar(&val, name, defval, helpstr)
	return &val
}

// ======

type def_Duration struct {
	name    string
	defval  time.Duration
	value   *time.Duration
	helpstr string
}

func (self def_Duration) getHelpString() string { return self.helpstr }
func (self def_Duration) getName() string       { return self.name }
func (self def_Duration) assign(value string) error {
	if duration, err := time.ParseDuration(value); err != nil {
		return err
	} else {
		*self.value = duration
		return nil
	}
}

// Register a time.Duration argument, storing to the supplied `value *time.Duration` pointer
func (p *Parser) DurationVar(value *time.Duration, name string, defval time.Duration, helpstr string) {
	vdef := def_Duration{name, defval, value, helpstr}
	*vdef.value = defval
	p.definitions[name] = vdef
	p.enqueueName(name)
}

// Register a time.Duration argument, storing to the returned `*time.Duration` pointer
func (p *Parser) Duration(name string, defval time.Duration, helpstr string) *time.Duration {
	var val time.Duration
	p.DurationVar(&val, name, defval, helpstr)
	return &val
}
