package main

import (
	"fmt"
	"os"
	"sort"
)

type Node struct {
	Char     byte
	Count    int
	Bits     int
	Encoding string
	Left     *Node
	Right    *Node
}

func (n Node) String() string {
	if n.Char == '\n' {
		return fmt.Sprintf("/N:(%d: '%v'[%d])", n.Count, n.Encoding, n.Bits)
	}
	return fmt.Sprintf("%c(%d: '%v'[%d])", n.Char, n.Count, n.Encoding, n.Bits)
}

func main() {
	fmt.Print("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var n = len(file)
	if false {
		fmt.Println(n)
	}
	// frequencyTable := freq(file[:n], n-1)
	frequencyTable := freq(file, n)
	// frequencyTable := map[byte]int{
	// 	'c': 32,
	// 	'd': 42,
	// 	'e': 120,
	// 	'k': 7,
	// 	'l': 42,
	// 	'm': 24,
	// 	'u': 37,
	// 	'z': 2,
	// }
	// printFreqTable(frequencyTable)
	// fmt.Println()

	t := initTree(frequencyTable)
	// fmt.Println("initTree -> ", len(t))
	// fmt.Println(frequencyTable)
	h := huffman(t)
	// fmt.Println(getLeaves(&h))
	encoding(&h, "")
	// fmt.Println()
	leaves := getLeaves(&h)
	fmt.Println(leaves)
	fmt.Println()
	header := createHeader(leaves)
	fmt.Println(header)
	fmt.Println()
	encs := getEncodings(leaves)
	fmt.Println(encs)
}

func createHeader(nodes []*Node) string {
	var header = ""
	var n = len(nodes)
	for i, node := range nodes {
		s := fmt.Sprintf("%v:%s", node.Char, node.Encoding)
		if i == n-1 {
			header += s
		} else {
			header += s
			header += ","
		}
	}
	return header
}

func convertToBitString(encs map[byte]string, encoding string) int {
	return 0
}

func encoding(tree *Node, prefix string) {
	if tree == nil {
		return
	}
	tree.Encoding = prefix
	tree.Bits = len(prefix)
	encoding(tree.Left, prefix+"0")
	encoding(tree.Right, prefix+"1")
	return
}

func printFreqTable(freq map[byte]int) {
	for k, v := range freq {
		if k == '\n' {
			fmt.Printf("\\N:%d, ", v)
		} else {
			fmt.Printf("%c:%d, ", k, v)
		}
	}
	fmt.Println()
}

func visualize(node *Node) {
	if node == nil {
		return
	}
	visualize(node.Left)
	fmt.Println(node.Count)
	visualize(node.Right)
	return
}

func getLeaves(tree *Node) []*Node {
	if tree == nil {
		return []*Node{}
	}
	if tree.Char != 0 {
		lt := append(getLeaves(tree.Left), tree)
		return append(lt, getLeaves(tree.Right)...)
	}
	return append(getLeaves(tree.Left), getLeaves(tree.Right)...)
}
func getEncodings(nodes []*Node) map[byte]string {
	encs := map[byte]string{}
	for _, node := range nodes {
		encs[node.Char] = node.Encoding
	}
	return encs
}

func huffman(nodes []Node) Node {
	// fmt.Println("huffman called: ", len(nodes))
	if len(nodes) == 1 {
		return nodes[0]
	}
	n1 := nodes[0]
	n2 := nodes[1]
	n := Node{Count: n1.Count + n2.Count, Left: &n1, Right: &n2, Char: 0}
	// fmt.Println(n1, n2, n)
	if len(nodes) == 2 {
		return n
	}
	h := huffman((insert(n, nodes[2:])))
	return h
}

func insert(node Node, nodes []Node) []Node {
	// fmt.Println(node, nodes)
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
	// fmt.Println(ans)
	return ans
}

func initTree(freq map[byte]int) []Node {
	var nodes []Node
	for k, v := range freq {
		nodes = append(nodes, Node{Char: k, Count: v, Left: nil, Right: nil})
	}
	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Count < nodes[j].Count
	})
	return nodes
}

func freq(file []byte, fileLength int) map[byte]int {
	dic := map[byte]int{}
	for i := 0; i < fileLength; i++ {
		dic[file[i]]++
	}
	return dic
}
