package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true // 2 and 3 are prime
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	// Check divisors from 5 to sqrt(nb), skipping even numbers
	for i := 5; i*i <= nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}
