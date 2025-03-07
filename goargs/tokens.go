package goargs

func SplitTokensBefore(delimiter string, args []string) ([]string, []string) {
    for i,arg := range(args) {
        if arg == delimiter {
            return args[:i], args[i+1:]
        }
    }

    return args[0:], []string{}
}
