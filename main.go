package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	f := flag.Bool("d", false, "to decompress")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("provide a file name!")
		os.Exit(1)
	}
	if *f == false {
		for _, fileName := range args {
			file, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			freqTable := NewFrequencyTable(string(file))
			t := NewHuffman(freqTable)
			root := Huffmanize(t)
			root.Encode("")
			FillFrequencyTable(root.GetLeaves(), freqTable)
			newFileName := fmt.Sprintf("%v_compressed.txt", fileName)
			Compress(string(file), freqTable, newFileName)
		}
	} else {
		for _, fileName := range args {
			file, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			newFileName := fmt.Sprintf("%v_decompressed.txt", fileName)
			Decompress(string(file), newFileName)
		}
	}

	fmt.Println(args)
}

// frequencyTable := NewFrequencyTable(string(file))
// t := NewHuffman(frequencyTable)
// root := Huffmanize(t)
// root.Encode("")

// // fmt.Println(frequencyTable)
// FillFrequencyTable(root.GetLeaves(), frequencyTable)
// fmt.Println(frequencyTable)
// // fmt.Println("inputString: ", string(file))
// // fmt.Println("bitString: ", bitString)
// // decodedString := root.Decode(&root, bitString, "")
// // fmt.Println("decodedString: ", decodedString)
// Compress(string(file), frequencyTable)

// fmt.Println()
// file2, _ := os.ReadFile("result.txt")
// Decompress(string(file2))
