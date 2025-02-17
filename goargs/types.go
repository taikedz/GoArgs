package goargs

import (
    "fmt"
    "strconv"
)


type StringDef struct {
    name string
    value *string
}
func (self StringDef) GetName() string { return self.name }
func (self StringDef) Assign(value string) error { *self.value = value; return nil }


type IntDef struct {
    name string
    value *int
}
func (self IntDef) GetName() string { return self.name }
func (self IntDef) Assign(value string) error {
    if value, err := strconv.Atoi(value); err != nil {
        return fmt.Errorf("Could not parse %s\n", value)
    } else {
        *self.value = value
    }
    return nil
}


type BoolDef {
    name string
    value *bool
}
func (self BoolDef) GetName() string { return self.name }
func (self BoolDef) Assign(value string) error {
    _ = value
    *self.value = !self.defval
    return nil
}

