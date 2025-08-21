package main

func SecondHalf(slice []int) []int {
    // find the starting index of the second half
    start := len(slice) / 2

    // if odd length, include the middle element in the second half
    if len(slice)%2 != 0 {
        start = len(slice) / 2
    }

    // return slice starting from 'start'
    return slice[start:]
}
