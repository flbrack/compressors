package main

import (
    "testing"
    "reflect"
    "fmt"
)

func TestCreateNodeList(t *testing.T) {
    inputs := []string{"AAABBC", "AABBBCB", "BB", "CBCBAAA"};
    expected := []NodeList {
        {{"A", 3, nil, nil}, {"B", 2, nil, nil}, {"C", 1, nil, nil}},
        {{"B", 4, nil, nil}, {"A", 2, nil, nil}, {"C", 1, nil, nil}},
        {{"B", 2, nil, nil}},
        {{"A", 3, nil, nil}, {"C", 2, nil, nil}, {"B", 2, nil, nil}},
    };

    var outputs [4]NodeList;
    for i, val := range inputs {
        outputs[i] = CreateNodeList(val);
    }

    for i, val := range outputs {
        if !reflect.DeepEqual(val, expected[i]) {
            t.Fatalf("Output: %v\nExpected: %v", val, expected[i]);
        }
    }
    fmt.Println("Counter tests passed");
}

func TestIsLeafNode(t *testing.T) {
    inputs := NodeList{
        {"a", 4, nil, nil},
        {"b", 10, &Node{"c",4,nil,nil}, &Node{"t",2,nil,nil}},
        {"2", 10, nil, &Node{" ",2,nil,nil}},
    }
    expected := []bool{true, false, false}
    var outputs [3]bool;
    for i, val := range inputs {
        outputs[i] = val.isLeaf();
    }
    for i, val := range outputs {
        if val != expected[i] {
            t.Fatalf("Output: %v\nExpected: %v", val, expected[i])
        }
    }
    fmt.Println("isLeaf tests passed")
}

func TestHuffmanEncoding(t *testing.T) {
    t.Skip("skipping")
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
