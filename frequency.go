package main

type FrequencyTable = map[byte]Property

type Property struct {
	Count    int
	Encoding string
}

func NewFrequencyTable(file string) FrequencyTable {
	table := FrequencyTable{}
	for i := 0; i < len(file); i++ {
		entry, ok := table[file[i]]
		if !ok {
			table[file[i]] = Property{Count: 1, Encoding: ""}
		} else {
			entry.Count += 1
			table[file[i]] = entry
		}
	}
	return table
}
