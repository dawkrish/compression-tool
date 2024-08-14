package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	frequencyTable := NewFrequencyTable(string(file))

	t := NewHuffman(frequencyTable)
	// fmt.Println("initTree -> ", len(t))
	// fmt.Println(frequencyTable)
	h := huffman(t)
	h.Encode("")
	leaves := h.GetLeaves()

	fmt.Println(leaves)
	fmt.Println()
	header := createHeader(leaves)
	fmt.Println(header)
	fmt.Println()
	encs := getEncodings(leaves)
	fmt.Println(encs)
}

func getEncodings(nodes []*Node) map[byte]string {
	encs := map[byte]string{}
	for _, node := range nodes {
		encs[node.Char] = node.Encoding
	}
	return encs
}

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
