package main

import (
	"fmt"
	"os"
	"strconv"
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

func CreateBitString(input string, freqTable FrequencyTable) string {
	fmt.Println("creating bit string...")
	outputFileBitString := ""
	for i := 0; i < len(input); i++ {
		outputFileBitString += freqTable[input[i]].StringEncoding
		outputFileBitString += ""
	}
	fmt.Println("created bit string!!!")
	return outputFileBitString
}

func WriteOutputFile(input string, bitString string) {
	output, err := os.Create("compressed.txt")
	if err != nil {
		fmt.Println("err WriteOutputFile: ", err)
		return
	}

	compressedBites := []byte{}
	newByte := ""
	byteLen := 0
	for i := 0; i < len(bitString); i++ {
		newByte += string(bitString[i])
		byteLen++

		if byteLen == 8 {
			b, err := strconv.ParseUint(newByte, 2, 8)
			if err != nil {
				fmt.Println("err ParseInt: ", err)
				return
			}
			compressedBites = append(compressedBites, byte(b))

			newByte = ""
			byteLen = 0
		}
	}

	b, err := strconv.ParseUint(newByte, 2, len(newByte))
	if err != nil {
		fmt.Println("err ParseInt: ", err)
		return
	}
	compressedBites = append(compressedBites, byte(b))

	_, err = output.Write(compressedBites)
	if err != nil {
		fmt.Println("err in Write: ", err)
	}
	return
}
