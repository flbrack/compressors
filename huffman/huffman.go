package main

import (
    "sort"
    // "fmt"
)

func Encode(s string) string {
    return ""
    // return "1000011101001000110010011101100111001001000111110010011111011111100010001111110100111001001011111011101000111111001";
}

type Node struct {
    value string
    freq int
    right *Node
    left *Node
}

type NodeList []Node

func (n Node) isLeaf() bool {
    return n.right == nil && n.left == nil
}

func (n NodeList) Less(i, j int) bool {
    return n[j].freq < n[i].freq
}

func (n NodeList) Len() int {
    return len(n)
}

func (n NodeList) Swap(i, j int) {
    n[i], n[j] = n[j], n[i]
}

func CreateNodeList(s string) NodeList {
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
    nodeList := make(NodeList, len(counter));
    i := 0;
    for k, v := range counter {
        nodeList[i] = Node{k, v, nil, nil};
        i++;
    }
    sort.Sort(nodeList)
    return nodeList
}


type Tree struct {
    head *Node
}

// func CreateHuffmanTree(nodeList []Node) Tree {
//     return
// }

