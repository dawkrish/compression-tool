package main

import (
	"fmt"
)

type FrequencyTable = map[byte]Property

type Property struct {
	Count          int
	StringEncoding string
	// NumericEncoding int64
	// ArrayEncoding   []byte
}

func NewFrequencyTable(file string) FrequencyTable {
	table := FrequencyTable{}
	for i := 0; i < len(file); i++ {
		entry, ok := table[file[i]]
		if !ok {
			table[file[i]] = Property{Count: 1}
		} else {
			entry.Count += 1
			table[file[i]] = entry
		}
	}
	return table
}

func MaxBits(leaves []*Node) int {
	max := leaves[0].Bits
	for _, leave := range leaves {
		if leave.Bits > max {
			max = leave.Bits
		}
	}
	return max
}

func FillFrequencyTable(leaves []*Node, freqTable FrequencyTable) {
	maxBits := MaxBits(leaves) + 1
	fmt.Println("Max Bits: ", maxBits)
	for _, leave := range leaves {
		entry, _ := freqTable[leave.Char]
		// n, err := strconv.ParseInt(leave.Encoding, 2, maxBits)
		// if err != nil {
		// 	fmt.Println("error in ParseInt: ", err)
		// } else {
		// 	fmt.Printf("Converted %v(%c) to: %d\n", leave.Encoding, leave.Char, n)
		// }
		// ns := n / 255
		// mod := n % 255
		// encs := []byte{}
		// for range ns {
		// 	encs = append(encs, 255)
		// }
		// encs = append(encs, byte(mod))

		entry.StringEncoding = leave.Encoding
		// entry.NumericEncoding = n
		// entry.ArrayEncoding = encs

		freqTable[leave.Char] = entry
	}
}
