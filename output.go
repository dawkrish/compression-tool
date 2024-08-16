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

func WriteOutputFile(freqTable FrequencyTable) {
	file, err := os.Create("compressed.txt")
	if err != nil {
		fmt.Println("err WriteOutputFile: ", err)
		return
	}
	compressedBites := []byte{}
	for _, v := range freqTable {
		encs := []byte{}
		ns := v.Encoding / 255
		mod := v.Encoding % 255
		for range ns {
			encs = append(encs, 255)
		}
		encs = append(encs, byte(mod))
		// fmt.Println(v.Encoding, encs)
		compressedBites = append(compressedBites, encs...)
	}
	fmt.Println(compressedBites)
	fmt.Println(len(compressedBites))
	_, err = file.Write(compressedBites)
	if err != nil {
		fmt.Println("err in Write: ", err)
	}
	return
}
