package main

import (
    "container/heap"
    "fmt"
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
    parent *Node
    index int
}

func (n Node) string() string {
    return fmt.Sprintf("Node: value:%v freq:%v", n.value, n.freq)
}

func (n *Node) isLeaf() bool {
    return n.right == nil && n.left == nil
}

type NodeHeap []*Node

func (nh NodeHeap) Less(i, j int) bool {
    if (nh[j].freq == nh[i].freq) {
        return nh[j].value < nh[i].value
    } else {
        return nh[j].freq < nh[i].freq
    }
}

func (nh NodeHeap) Len() int {
    return len(nh)
}

func (nh NodeHeap) Swap(i, j int) {
    nh[i], nh[j] = nh[j], nh[i]
    nh[i].index = i
    nh[j].index = j
}

func (nh *NodeHeap) Push(x any) {
    n := len(*nh)
    node := x.(*Node)
    node.index = n
    *nh = append(*nh, node)
}

func (nh *NodeHeap) Pop() any {
    old := *nh
    n := len(old)
    node := old[n-1]
    old[n-1] = nil
    node.index = -1
    *nh = old[:n-1]
    return node
}

func createNodeHeap(s string) NodeHeap {
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
    nodeHeap := make(NodeHeap, len(counter));
    i := 0;
    for k, v := range counter {
        nodeHeap[i] = &Node{value: k, freq: v, index: i};
        i++;
    }
    heap.Init(&nodeHeap)
    return nodeHeap
}


type Tree struct {
    head *Node
}

// func createHuffmanTree(nodeHeap []Node) Tree {
//     var least Node 
//     var secondLeast Node
//     for range nodeHeap {
//         least, secondLeast, nodeHeap = nodeHeap[len(nodeHeap)-1], nodeHeap[len(nodeHeap)-2], nodeHeap[:len(nodeHeap)-1]
//
//     }
// }

