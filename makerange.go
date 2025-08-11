package piscine

func makerange(min, max int) []int {
	if min >= max {
		return nil // return nil slice if invalid range
	}

	// Create a slice big enough to hold all numbers from min to max-1
	size := max - min
	result := make([]int, size)

	// Fill the slice
	for i := 0; i < size; i++ {
		result[i] = min + i
	}

	return result
}

// Check invalid range → if min >= max, return nil.

// Calculate size → the number of elements will be max - min.

// Create slice with make → make([]int, size) creates a slice full of zeroes with the exact size.

// Fill manually → result[i] = min + i puts the correct number at each position.

// Return → slice is complete, no append needed.