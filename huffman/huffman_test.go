package main

import (
    "testing"
    "reflect"
    "fmt"
)

func TestCounter(t *testing.T) {
    inputs := []string{"AAABBC", "AABBBCB", "BB", "CBCBAAA"};
    expected := [4]map[string]int {
        {"A":3, "B":2, "C":1},
        {"A":2, "B":4, "C":1},
        {"B":2},
        {"A":3, "B":2, "C":2},
    };

    var outputs [4]map[string]int;
    for i, val := range inputs {
        outputs[i] = Counter(val);
    }

    for i, val := range outputs {
        if !reflect.DeepEqual(val, expected[i]) {
            t.Fatalf("Output: %v\nExpected: %v", val, expected[i]);
        }
    }

    fmt.Println("Counter tests passed");
}

func TestSortCounter(t *testing.T) {
    input := map[string]int {"A":3, "B":2, "C":1};
    expected := PairList{
        {"A", 3},
        {"B", 2},
        {"C", 1},
    };
    output := SortCounter(input)

    if !reflect.DeepEqual(output, expected) {
        t.Fatalf("Output: %v\nExpected: %v", output, expected);
    }

    fmt.Println("SortCounter tests passed");
}

func TestHuffmanEncoding(t *testing.T) {
    input := "A_DEAD_DAD_CEDED_A_BAD_BABE_A_BEADED_ABACA_BED";
    expected := "1000011101001000110010011101100111001001000111110010011111011111100010001111110100111001001011111011101000111111001";
    output := Encode(input);

    if output != expected {
        t.Fatalf("Output: %s \nExpected: %s",
                output, expected);
    }

    fmt.Println("Encode tests passed");
    fmt.Printf("# Bits in Input: %d \t# Bits in Output: %d\n\n",
                len(input)*8, len(output));
}
