package main

import (
	"fmt"
	"strconv"
)

type FrequencyTable = map[byte]Property

type Property struct {
	Count    int
	Encoding int64
}

func NewFrequencyTable(file string) FrequencyTable {
	table := FrequencyTable{}
	for i := 0; i < len(file); i++ {
		entry, ok := table[file[i]]
		if !ok {
			table[file[i]] = Property{Count: 1, Encoding: 0}
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
		n, err := strconv.ParseInt(leave.Encoding, 2, maxBits)
		if err != nil {
			fmt.Println("error in ParseInt: ", err)
		} else {
			fmt.Printf("Converted %v(%c) to: %d\n", leave.Encoding, leave.Char, n)
		}
		entry.Encoding = n
		freqTable[leave.Char] = entry
	}
}
