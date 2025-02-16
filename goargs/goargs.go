package goargs

import "fmt"

type Parser struct {
    definitions []ArgumentGeneric
}


func (self Parser) AddArgument(arg ArgumentGeneric) {
    self.definitions = append(self.definitions, arg)
}


type ArgumentGeneric interface {
    GetName() string
    GetValueString() string
    GetFlags() (string, rune)
}

type StrDef struct {
    value string
    name string
    longflag string
    shortflag rune
}
func (self StrDef) GetName() string { return self.name }
func (self StrDef) GetValueString() string { return self.value }
func (self StrDef) GetFlags() (string, rune) { return self.longflag, self.shortflag }

// ---------

func SpotCheck() {
    var p = Parser{[]ArgumentGeneric{}}

    myarg := StrDef{"alice", "name", "name", 'N'}

    p.AddArgument( myarg )


    fmt.Printf("%s -> %s\n", myarg.GetName(), myarg.GetValueString())

    fmt.Println(p)

    var holder []ArgumentGeneric
    holder = append(holder, myarg)
    fmt.Printf("%d/%d -> ", len(holder), cap(holder))
    fmt.Println(holder)
}
