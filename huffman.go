package main

import (
	"fmt"
	"sort"
)

// Represents a Node in an Huffman Tree
type Node struct {
	Char     byte
	Count    int
	Bits     int
	isLeaf   bool
	Encoding string

	Left  *Node
	Right *Node
}

// String representation of huffman node
func (n Node) String() string {
	if n.Char == '\n' {
		return fmt.Sprintf("/N:(%d: '%v'[%d])", n.Count, n.Encoding, n.Bits)
	}
	return fmt.Sprintf("%c(%d: '%v'[%d])", n.Char, n.Count, n.Encoding, n.Bits)
}

// Creates Initial Tree([]Node) from a Frequency Table
func NewHuffman(freqTable FrequencyTable) []Node {
	var nodes []Node
	for k, v := range freqTable {
		n := Node{
			Char:   k,
			Count:  v.Count,
			Left:   nil,
			Right:  nil,
			isLeaf: true,
		}
		nodes = append(nodes, n)
	}
	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Count < nodes[j].Count
	})
	return nodes
}

// Collapses nodes([]Node) into a single node
// construct the Huffman tree
// the final return is the root of the tree
func Huffmanize(nodes []Node) Node {
	if len(nodes) == 1 {
		return nodes[0]
	}
	n1 := nodes[0]
	n2 := nodes[1]
	n := Node{Count: n1.Count + n2.Count, Left: &n1, Right: &n2, isLeaf: false}
	if len(nodes) == 2 {
		return n
	}
	newNodes := insertNode(n, nodes[2:])
	return Huffmanize(newNodes)
}

// A helper function to Insert in a sorted list
func insertNode(node Node, nodes []Node) []Node {
	ans := []Node{}
	i := 0
	for ; i < len(nodes); i++ {
		if nodes[i].Count >= node.Count {
			break
		}
	}
	ans = append(ans, nodes[:i]...)
	ans = append(ans, node)
	ans = append(ans, nodes[i:]...)
	return ans
}

// Now we have a huffman tree ready, we will encode it
func (root *Node) Encode(prefix string) {
	if root == nil {
		return
	}
	root.Encoding = prefix
	root.Bits = len(prefix)
	root.Left.Encode(prefix + "0")
	root.Right.Encode(prefix + "1")
	return
}

func (root *Node) GetLeaves() []*Node {
	if root == nil {
		return []*Node{}
	}
	leaves := []*Node{}
	left := root.Left.GetLeaves()
	right := root.Right.GetLeaves()

	leaves = append(leaves, left...)
	if root.isLeaf {
		leaves = append(leaves, root)
	}
	leaves = append(leaves, right...)

	return leaves
}
