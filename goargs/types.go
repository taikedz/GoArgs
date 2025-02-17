package goargs

import (
    "fmt"
    "strconv"
)

// Unpack arguments into a series of variables, and return any unassigned values.
// Typically for use in conjunction with Parser.Positionals
func Unpack(tokens []string, vars ...interface{}) ([]string, error) {
    max := len(tokens)
    if len(vars) < max {
        max = len(vars)
    }

    for i := 0; i<max; i++ {
        tok := tokens[i]
        label := vars[i]

        switch t := label.(type) {
            default:
                return nil, fmt.Errorf("Unsupported type: %t", t)
            case *string:
                // `label` is an interface, and cannot be directly dereferenced and assigned
                //   so we performe some indirection via a new explicitly typed variable
                var lab *string = label.(*string)
                *lab = tok
            case *int:
                val, err := strconv.Atoi(tok)
                if err != nil {
                    return nil, fmt.Errorf("Could not parse int %s : %v", tok, err)
                }
                var lab *int = label.(*int)
                *lab = val
            case *float32:
                float, err := strconv.ParseFloat(tok, 32)
                if err != nil {
                    return nil, fmt.Errorf("Could not parse int %s : %v", tok, err)
                }
                var lab *float32 = label.(*float32)
                *lab = float32(float)
        }
    }

    return tokens[max:], nil
}


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

