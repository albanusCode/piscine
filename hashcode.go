package main

func HashCode(dec string) string {
	size := len(dec)
	result := ""

	for i := 0; i < size; i++ {
		val := (int(dec[i]) + size) % 127
		if val < 32 { // make sure it's printable
			val += 33
		}
		if val > 127 {
			temp := val - 127
			if temp < 32 {
				temp += 33
				result += string(rune(temp))
			}
		}
		result += string(rune(val))
	}

	return result
}


// func HashCode(s string) int {
// 	hash := 0
// 	p := 31
// 	mod := 1_000_000_007 // big prime
// 	for i := 0; i < len(s); i++ {
// 		hash = (hash*p + int(s[i])) % mod
// 	}
// 	return hash
// }


// package piscine

// func HashCode(s string) int {
// 	hash := 0
// 	p := 31 // base prime
// 	for i := 0; i < len(s); i++ {
// 		hash = hash*p + int(s[i])
// 	}
// 	return hash
// }
