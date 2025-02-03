package goargs

import (
    "fmt"
)




type Parser struct {
    // We just store the string. Conversion happens during retrieval
    opt_items map[string]string
    pos_items []string

    parse_done bool // FIXME can we default this to false ?

    flag_defs []ArgDefType
}


func (self *Parser) Parse(tokens []string, ignore_unknown bool) error {
    if self.parse_done {
        return fmt.Errorf("Dev error: Attempted to re-parse after parse already performed. Use `reset()` first?\n")
    }

    self.subParse(tokens, ignore_unknown)

    self.parse_done = true
}

func (self *Parser) subParse(tokens []string, ignore_unknown bool) error {

    for i := 0; i<len(tokens); i++ {
        tk := tokens[i]
        
        if ! strings.StartsWith("-", tk) {
            append(pos_tems, tk)
        } else if strings.StartsWith("--") {
            name, valuecount, settrue, err := self.getFlagRequirement()
            if err != nil {
                if ignore_unknown continue
                else return err
            }
            if valuecount > 0 {
            } else {
                self.opt_items = ! settrue
            }
        }

        // ----
        if seq := extractShortFlags(tk); seq != nil {
            for _, argdef := seq {
                if argdef.is_boolean {
                    self.opt_items[argdef.name] = true
                } else {
                    i++
                    self.opt_items[name] = tokens[i]
                }
            }
        } else if ! self.defines(tk) {
            if ignore_unknown continue
            else return fmt.Errorf("Unkown parameter '%s'")
        }
    }

    return nil
}

funct (self *Parser) defines(name string) {
}

/** Clear all parsed arguments data from parser.
 *
 * Allow re-use of parser by cleansing it first
 */
func (self *Parser) Reset() {
    clear(self.items)
    self.parse_done = false
}

func (self *Parser) AddArgument(definition ArgDefType) {
    append(self.arg_defs, definition)
}


func extractShortFlags(combined string) []ParamCount {
    // parse combination of short flags into their names, and number of paramters each consumes
    // and return as a sequence
}


// Argument fetching

func (self *Parser) AsString(name string) string {
}

func (self *Parser) AsInt32(name string) int32 {
}

func (self *Parser) AsUint32(name string) uint32 {
}

func (self *Parser) AsFloat32(name string) float32 {
}

func (self *Parser) AsBool(name string) bool {
}

