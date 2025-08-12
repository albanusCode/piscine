package piscine

func GCD(a, b int) int {
	// Make sure both numbers are positive
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// package piscine

// func GCD(a, b int) int {
// 	if a < 0 {
// 		a = -a
// 	}
// 	if b < 0 {
// 		b = -b
// 	}
// 	for b != 0 {
// 		a, b = b, a%b
// 	}
// 	return a
// }

// func GCD3(a, b, c int) int {
// 	return GCD(GCD(a, b), c)
// }
