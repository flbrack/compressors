package main

import (
    "container/heap"
    "fmt"
)

func Encode(s string) string {
    nh := createNodeHeap(s)
    tree := createHuffmanTree(&nh)
    m := make(map[string]string)
    createDict(tree.head, &m)
    output := ""
    for _, l := range s {
        output += m[string(l)]
    }
    return output
    // return "1000011101001000110010011101100111001001000111110010011111011111100010001111110100111001001011111011101000111111001"
}

func Decode(s string, t *HuffmanTree) string {
    output := ""
    current := t.head
    for _, val := range s {
        if (string(val) == "0") {
            current = current.left
        } else if (string(val) == "1") {
            current = current.right
        }
        if (current.isLeaf()) {
            output += current.value
            current = t.head
        }
    }
    return output
}

type Node struct {
    value string
    freq int
    right *Node
    left *Node
    parent *Node
    index int
}

func (n *Node) String() string {
    return fmt.Sprintf("Node - value: %v freq: %v", n.value, n.freq)
}

func (n *Node) isLeaf() bool {
    return n.right == nil && n.left == nil
}

type NodeHeap []*Node

func (nh NodeHeap) Less(i, j int) bool {
    if (nh[j].freq == nh[i].freq) {
        return nh[j].value < nh[i].value
    } else {
        return nh[i].freq < nh[j].freq
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
    counter := make(map[string]int)
    for _, val := range s {
        strVal := string(val)
        _, ok := counter[strVal]
        if !ok {
            counter[strVal] = 1
        } else {
            counter[strVal] += 1
        }
    }
    nodeHeap := make(NodeHeap, len(counter))
    i := 0
    for k, v := range counter {
        nodeHeap[i] = &Node{value: k, freq: v, index: i}
        i++
    }
    heap.Init(&nodeHeap)
    return nodeHeap
}


type HuffmanTree struct {
    head *Node
}

func createHuffmanTree(nodeHeap *NodeHeap) HuffmanTree {
    for nodeHeap.Len() > 1 {
        first := heap.Pop(nodeHeap).(*Node)
        second := heap.Pop(nodeHeap).(*Node)
        var parent = Node{
            value: first.value + second.value,
            freq: first.freq + second.freq,
            left: first,
            right: second}
        first.parent = &parent
        second.parent = &parent
        heap.Push(nodeHeap, &parent)
    }
    finalNode := heap.Pop(nodeHeap).(*Node)
    return HuffmanTree{finalNode}
}

func createDict(head *Node, m *map[string]string) {
    current := head
    if (current.isLeaf()) {
        _, ok := (*m)[current.value]
        if !ok {
            (*m)[current.value] = bubbleUp(current)
        }
        return
    }
    createDict(current.left, m)
    createDict(current.right, m)
}

func bubbleUp(node *Node) string {
    current := node
    s := ""
    for current.parent != nil {
        if (current.parent.left == current) {
            s = "0" + s
        } else {
            s = "1" + s
        }
        current = current.parent
    }
    return s
}

