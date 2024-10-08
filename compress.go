package main

import (
	"fmt"
	"os"
	"strconv"
)

func CreateHeader(freqTable FrequencyTable) string {
	var header = ""
	for k, v := range freqTable {
		h := fmt.Sprintf("%v:%d,", v.StringEncoding, k)
		header += h
	}
	header += "\n"
	return header
}

func Compress(input string, freqTable FrequencyTable, newFileName string) {
	header := CreateHeader(freqTable)
	var compressedBytes = []byte{}
	compressedBytes = append(compressedBytes, []byte(header)...)

	bitString := ""
	bitStringLen := 0

	for i := 0; i < len(input); i++ {
		entry := freqTable[input[i]]
		bitString += entry.StringEncoding
		bitStringLen += entry.Bits
		for {
			if bitStringLen < 8 {
				break
			}
			b, err := strconv.ParseUint(bitString[:8], 2, 8)
			if err != nil {
				fmt.Println("err strconv.ParseUint: ", err)
			}
			compressedBytes = append(compressedBytes, byte(b))

			bitString = bitString[8:]
			bitStringLen -= 8
		}
	}
	if bitStringLen != 0 {
		b, err := strconv.ParseUint(bitString, 2, 8)
		if err != nil {
			fmt.Println("err strconv.ParseUint: ", err)
		}
		compressedBytes = append(compressedBytes, byte(b))

	}

	outputFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("err os.Create: ", err)
	}
	_, err = outputFile.Write(compressedBytes)
	if err != nil {
		fmt.Println("err in Write: ", err)
	}
	return
}
