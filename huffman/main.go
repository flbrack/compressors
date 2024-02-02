package main

import (
    "os"
    "fmt"
)

func main() {
    input := os.Args[1]
    output := Encode(input)

    fmt.Println("Original string: " + input)
    fmt.Println("Compressed string: " + output)
    fmt.Printf("# Bits in Input: %d \t# Bits in Output: %d\n",
                len(input)*8, len(output))
}
