package main

func IsPrintable(s string) bool {
	for _, r := range s {
		// Printable ASCII range is from 32 (' ') to 126 ('~')
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}
