package main

import "sort"

func Encode(s string) string {
    return ""
    // return "1000011101001000110010011101100111001001000111110010011111011111100010001111110100111001001011111011101000111111001";
}

func Counter(s string) map[string]int {
    counter := make(map[string]int);
    for _, val := range s {
        strVal := string(val);
        _, ok := counter[strVal];
        if !ok {
            counter[strVal] = 1;
        } else {
            counter[strVal] += 1;
        }
    }
    return counter;
}

type Node struct {
    right *Node
    left *Node
}

type LeafNode struct {
    Node
    value string
}

type Pair struct {
    Key string
    Value int
}

func SortCounter(c map[string]int) []Pair {
    p := make([]Pair, len(c));
    i := 0;
    for k, v := range c {
        p[i] = Pair{k, v};
        i++;
    }
    sort.Slice(p, func(i, j int) bool {
        return p[i].Value > p[j].Value
    })
    return p;
}

