package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1) Get out the header
2) Use the decode function
*/

func Decompress(input string, newFileName string) {
	lines := strings.Split(input, "\n")

	header := lines[0]
	keyvals := strings.Split(header, ",")
	lookup := map[string]byte{}

	for _, kv := range keyvals {
		arr := strings.Split(kv, ":")
		if len(arr) == 1 {
			continue
		}
		b, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println("err strconv.Atoi: ", err)
		}
		lookup[arr[0]] = byte(b)
	}

	content := strings.Join(lines[1:], "\n")
	contentBitString := ""
	decompressedBytes := []byte{}
	last_i := len(content) - 1

	for i := 0; i < len(content); i++ {
		if i == last_i {
			contentBitString += fmt.Sprintf("%b", content[i])
		} else {
			contentBitString += fmt.Sprintf("%08b", content[i])
		}
		var possibleKey = ""
		for j := 0; j < len(contentBitString); j++ {
			possibleKey += string(contentBitString[j])
			entry, ok := lookup[possibleKey]
			if ok {
				decompressedBytes = append(decompressedBytes, entry)
				possibleKey = ""
			}
		}
		contentBitString = possibleKey
	}

	outputFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("err os.Create: ", err)
	}
	_, err = outputFile.Write(decompressedBytes)
	if err != nil {
		fmt.Println("err in Write: ", err)
	}
}

// if len(contentBitString) != 0 {
// 	fmt.Println("leftbitstring:", contentBitString)
// 	fmt.Println(lookup)
// 	entry, ok := lookup[contentBitString]
// 	fmt.Println(entry, ok)
// 	if ok {
// 		decompressedBytes = append(decompressedBytes, entry)
// 		contentBitString = ""
// 	}
// }
