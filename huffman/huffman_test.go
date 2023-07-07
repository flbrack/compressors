package main

import (
    "testing"
    "reflect"
    "fmt"
    "container/heap"
)

func TestCreateNodeHeap(t *testing.T) {
    inputs := []string{"AAABBC", "AABBBCB", "BB", "CBCBAAA"};
    expected := []NodeHeap {
        {
            {value: "A", freq: 3, index: -1},
            {value: "B", freq: 2, index: -1},
            {value: "C", freq: 1, index: -1},
        },
        {
            {value: "B", freq: 4, index: -1},
            {value: "A", freq: 2, index: -1},
            {value: "C", freq: 1, index: -1},
        },
        {
            {value: "B", freq: 2, index: -1},
        },
        {
            {value: "A", freq: 3, index: -1},
            {value: "C", freq: 2, index: -1},
            {value: "B", freq: 2, index: -1},
        },
    };

    var outputs [4]NodeHeap;
    for i, val := range inputs {
        outputs[i] = createNodeHeap(val);
    }

    for i, nodeHeap := range outputs {
        for j := range nodeHeap {
            node := heap.Pop(&nodeHeap).(*Node)
            if !reflect.DeepEqual(node, expected[i][j]) {
                t.Fatalf("Output: %v\nExpected: %v", node, expected[i][j]);
            }
        }
    }
    fmt.Println("Create heap tests passed");
}

func TestIsLeafNode(t *testing.T) {
    inputs := NodeHeap{
        {value: "a", freq: 4},
        {
            value: "b",
            freq: 10, 
            right: &Node{value: "c", freq: 4},
            left: &Node{value: "t", freq: 2},
        },
        {value: "2", freq: 10, right: &Node{value: " ", freq: 2}},
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
