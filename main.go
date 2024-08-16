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

	// fmt.Println(frequencyTable)
	FillFrequencyTable(root.GetLeaves(), frequencyTable)
	fmt.Println(frequencyTable)
	bitString := CreateBitString(string(file), frequencyTable)
	// fmt.Println("inputString: ", string(file))
	// fmt.Println("bitString: ", bitString)
	// decodedString := root.Decode(&root, bitString, "")
	// fmt.Println("decodedString: ", decodedString)
	WriteOutputFile(string(file), bitString)
}
