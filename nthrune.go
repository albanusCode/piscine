package piscine

func NRune(s string, n int) rune {
	// Convert string to slice of runes (handles UTF-8 characters correctly)
	runes := []rune(s)

	// Check if n is valid (1-based index)
	if n <= 0 || n > len(runes) {
		return 0
	}

	// Return the nth rune (index n-1 because slices are 0-based)
	return runes[n-1]
}