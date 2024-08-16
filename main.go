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
	root := Huffmanize(t)
	root.Encode("")

	fmt.Println(frequencyTable)
	FillFrequencyTable(root.GetLeaves(), frequencyTable)
	fmt.Println(frequencyTable)
	WriteOutputFile(frequencyTable)
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
