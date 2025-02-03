package goargs

import "fmt"

type Parser struct {
    definitions []*ArgumentGeneric
}


func (self *Parser) AddArgument(arg *ArgumentGeneric) {
    append(self.definitions, arg)
}


type ArgumentGeneric interface {
    GetName() string
    GetFlags() (string, rune)
}

type StrDef struct {
    value *string
    name string
    longflag string
    shortflag rune
}
func (self *StrDef) GetName() string { return self.name }
func (self *StrDef) GetFlags() (string, rune) { return self.longflag, self.shortflag }

func SpotCheck() {
    alice := "alice"

    var p Parser

    p.AddArgument( &StrDef{&alice, "name", "name", 'N'} )

    fmt.Println(p)
}
