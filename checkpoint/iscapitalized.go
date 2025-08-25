package main

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	words := []rune{}
	wordStart := true

	// We'll manually split the string by spaces (to avoid built-in strings.Fields)
	wordsList := [][]rune{}
	currentWord := []rune{}

	for _, ch := range s {
		if ch == ' ' {
			if len(currentWord) > 0 {
				wordsList = append(wordsList, currentWord)
				currentWord = []rune{}
			}
		} else {
			currentWord = append(currentWord, ch)
		}
	}
	if len(currentWord) > 0 {
		wordsList = append(wordsList, currentWord)
	}

	if len(wordsList) == 0 {
		return false
	}

	for _, w := range wordsList {
		firstChar := w[0]
		if firstChar >= 'a' && firstChar <= 'z' {
			return false
		}
	}

	return true
}
