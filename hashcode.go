func HashCode(s string) int {
	hash := 0
	p := 31
	mod := 1_000_000_007 // big prime
	for i := 0; i < len(s); i++ {
		hash = (hash*p + int(s[i])) % mod
	}
	return hash
}


// package piscine

// func HashCode(s string) int {
// 	hash := 0
// 	p := 31 // base prime
// 	for i := 0; i < len(s); i++ {
// 		hash = hash*p + int(s[i])
// 	}
// 	return hash
// }
