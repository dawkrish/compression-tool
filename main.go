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

			fi, err := os.Stat(fileName)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%v: %d bytes\n", fileName, fi.Size())

			fi, err = os.Stat(newFileName)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%v: %d bytes\n", newFileName, fi.Size())
			fmt.Println()
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
}
