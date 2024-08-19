package main

type FrequencyTable = map[byte]Property

type Property struct {
	Count          int
	StringEncoding string
	Bits           int
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

func FillFrequencyTable(leaves []*Node, freqTable FrequencyTable) {
	for _, leave := range leaves {
		entry, _ := freqTable[leave.Char]
		entry.StringEncoding = leave.Encoding
		entry.Bits = len(leave.Encoding)
		freqTable[leave.Char] = entry
	}
}
