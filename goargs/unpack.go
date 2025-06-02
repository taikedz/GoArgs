package goargs

import (
	"fmt"
	"strconv"
)

// Unpack arguments into a series of variables, and return any unassigned values.
// Typically for use in conjunction with Parser.Args()
//
// `vars` are pointers to supported types.
//
// Supported types: *string, *int, *float32
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
                return nil, fmt.Errorf("Unsupported type (did you use a pointer?): %t", t)
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
            case *bool:
                var lab *bool = label.(*bool)
                switch tok {
                case "false":
                    *lab = false
                case "0":
                    *lab = false
                case "true":
                    *lab = true
                case "1":
                    *lab = true
                default:
                    return nil, fmt.Errorf("Invalid string value for boolean: %v . Try 'true', 'false, '1', or '0'.", tok)
                }
        }
    }

    return tokens[max:], nil
}

// Unpack tokens, expecting the number of variables and number of tokens to match.
// Returns
//    any error
func UnpackExactly(tokens []string, vars ...interface{}) error {
    if len(tokens) != len(vars) {
        // FIXME - how does dev user differentiate a parse error from a mismatch error?
        //         dev will surely want to handle the two cases differently...
        return fmt.Errorf("Mismatch number of tokens (%d) to number of variables to populate (%d)", len(tokens), len(vars))
	}

	if _, err := Unpack(tokens, vars...); err != nil {
		return err
	}

	return nil
}
