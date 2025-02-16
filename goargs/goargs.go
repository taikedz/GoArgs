package goargs

import (
    "fmt"
    "os"
    "strconv"
)


type VarDef interface {
    GetName() string
    Assign(string)
}

type StringDef struct {
    name string
    value *string
}
func (self StringDef) GetName() string { return self.name }
func (self StringDef) Assign(value string) { *self.value = value }

type IntDef struct {
    name string
    value *int
}
func (self IntDef) GetName() string { return self.name }
func (self IntDef) Assign(value string) {
    if value, err := strconv.Atoi(value); err != nil {
        panic(fmt.Sprintf("Could not parse %s\n", value) )
    } else {
        *self.value = value
    }
}


var definitions = []VarDef{}


func StringArg(value *string, name string, defval string) {
    s := StringDef{name, value}
    *s.value = defval
    definitions = append(definitions, s)
}

func IntArg(value *int, name string, defval int) {
    s := IntDef{name, value}
    *s.value = defval
    definitions = append(definitions, s)
}

func Parse() {
    for i, str := range os.Args[1:] {
        definitions[i].Assign(str)
    }
}

// ---------

func SpotCheck() {
    var person string
    var age int
    StringArg(&person, "name", "nobody")
    IntArg(&age, "age", -1)
    Parse()
    fmt.Printf("%s -> %d\n", person, age)
    fmt.Printf("%v\n", definitions)
}
