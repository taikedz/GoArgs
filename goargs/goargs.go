package goargs

import (
    "fmt"
    "os"
)


type VarDef struct {
    name string
    value *string
    defval string
}


var definitions = []VarDef{}


func StringArg(value *string, name string, defval string) {
    definitions = append(definitions, VarDef{name, value, defval} )
}

func Parse() {
    if len(os.Args) == 1 {
        *definitions[0].value = definitions[0].defval
    } else {
        *definitions[0].value = os.Args[1]
    }
}

// ---------

func SpotCheck() {
    var person string
    StringArg(&person, "name", "nobody")
    Parse()
    fmt.Println(person)
    fmt.Printf("%v\n", definitions)
}
