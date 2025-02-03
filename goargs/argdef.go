

type Range struct {
    low uint
    high uint
}

func (self *Range) contains(number uint) bool {
    return self.low <= number && number <= self.high
}

type ARGDEF int


type ArgDef struct {
    name string
    long_flag string
    short_flag string
    value_count bool

    string_default string
    int_default int
    bool_default bool

    // restrict possible values
    options []string
    val_range Range
}

func StrFlag(name string, long_flag string, short_flag string, def string) {
    return ArgDef{name, long_flag, short_flag, 1, def, 0, false, nil, nil}
}

func StrFlagRestricted(name string, long_flag string, short_flag string, def string, restrict []string) {}

func IntFlag(name string, long_flag string, short_flag string, def int) {}

func IntFlagRestricted(name string, long_flag sting, short_flag string, def int, restrict Range) {}

func BoolFlag(name string, long_flag string, short_flag string, def bool) {}

