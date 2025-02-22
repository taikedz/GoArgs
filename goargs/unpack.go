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
        }
    }

    return tokens[max:], nil
}

// Unpack tokens, expecting the number of variables and number of tokens to match.
// Returns
//    remainder (if any, or nil)
//    parse error (if any, or nil)
//    count error (if any, or nil)
func UnpackExactly(tokens []string, vars ...interface{}) ([]string, error, error) {
	var err error
	remains, err := Unpack(tokens, vars...)

	if err != nil {
		return remains, err, nil
	}
	if len(tokens) > len(vars) {
		return remains, nil, fmt.Errorf("Excess tokens")
	}
	if len(tokens) < len(vars) {
		return nil, nil, fmt.Errorf("Last %d vars were uninitialized", len(vars)-len(tokens))
	}

	return nil, nil, nil
}