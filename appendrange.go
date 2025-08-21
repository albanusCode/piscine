package main

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	var result []int // nil slice to start with

	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}

// takes an int min and an int max as parameters. The function should return a slice of ints with all the values between the min and max.

// Min is included, and max is excluded.

// If min is greater than or equal to max, a nil slice is returned.

// make is not allowed.
