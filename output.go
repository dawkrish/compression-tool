package main

import (
	"fmt"
	"os"
)

func CreateHeader(leaves []*Node) string {
	var header = ""
	var n = len(leaves)
	for i, leave := range leaves {
		s := fmt.Sprintf("%v:%s", leave.Char, leave.Encoding)
		if i == n-1 {
			header += s
		} else {
			header += s
			header += ","
		}
	}
	return header
}

func WriteOutputFile(input string, freqTable FrequencyTable) {
	output, err := os.Create("compressed.txt")
	if err != nil {
		fmt.Println("err WriteOutputFile: ", err)
		return
	}
	compressedBites := []byte{}

	for i := 0; i < len(input); i++ {
		compressedBites = append(compressedBites, freqTable[input[i]].ArrayEncoding...)
	}

	_, err = output.Write(compressedBites)
	if err != nil {
		fmt.Println("err in Write: ", err)
	}
	return
}
