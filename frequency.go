package main

type FrequencyTable = map[byte]*Property

type Property struct {
	Count    int
	Encoding string
}

func NewFrequencyTable(file string) FrequencyTable {
	table := FrequencyTable{}
	for i := 0; i < len(file); i++ {
		table[file[i]].Count++
	}
	return table
}
