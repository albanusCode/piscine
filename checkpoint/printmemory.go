package piscine

import "fmt"

func PrintMemory(arr [10]byte) {
    // Print bytes in hex, 4 bytes per line
    for i := 0; i < len(arr); i += 4 {
        end := i + 4
        if end > len(arr) {
            end = len(arr)
        }
        for j := i; j < end; j++ {
            fmt.Printf("%02x ", arr[j])
        }
        fmt.Println()
    }

    // Print ASCII representation
    for i := 0; i < len(arr); i++ {
        b := arr[i]
        if b >= 32 && b <= 126 {
            fmt.Printf("%c", b)
        } else {
            fmt.Printf(".")
        }
    }
    fmt.Println()
}
