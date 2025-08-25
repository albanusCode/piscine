package main

func SaveAndMiss(arg string, num int) string {
    if num <= 0 {
        return arg
    }

    result := ""

    // We will use an index to move through the string
    for i := 0; i < len(arg); i += num * 2 {
        // Save chunk from i to i+num (or end of string)
        end := i + num
        if end > len(arg) {
            end = len(arg)
        }
        result += arg[i:end]
        // Then skip next chunk (i + num to i + 2*num) automatically by increment
    }

    return result
}
