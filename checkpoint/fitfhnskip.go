package main

func FifthAndSkip(str string) string {
    // Remove spaces first
    cleaned := []rune{}
    for _, ch := range str {
        if ch != ' ' {
            cleaned = append(cleaned, ch)
        }
    }

    if len(cleaned) == 0 {
        return "\n"
    }

    if len(cleaned) < 5 {
        return "Invalid Input\n"
    }

    result := ""
    i := 0
    n := len(cleaned)

    for i < n {
        // Take 5 chars or less if at the end
        end := i + 5
        if end > n {
            end = n
        }
        // Add group of chars
        result += string(cleaned[i:end])
        i = end

        // Skip 1 char if possible
        if i < n {
            i++
        }

        // Add space if more chars left
        if i < n {
            result += " "
        }
    }

    return result
}
