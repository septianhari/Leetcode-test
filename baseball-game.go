package main

import (
	"strconv"
)

func calPoints(operations []string) int {
	var record []int

	for _, op := range operations {
		switch op {
		case "C":
			// Batalkan skor terakhir
			if len(record) > 0 {
				record = record[:len(record)-1]
			}
		case "D":
			// Gandakan skor terakhir
			if len(record) > 0 {
				record = append(record, 2*record[len(record)-1])
			}
		case "+":
			// Tambahkan dua skor terakhir
			if len(record) >= 2 {
				record = append(record, record[len(record)-1]+record[len(record)-2])
			}
		default:
			// Tambahkan skor baru
			score, _ := strconv.Atoi(op)
			record = append(record, score)
		}
	}

	// Hitung total skor
	total := 0
	for _, score := range record {
		total += score
	}
	return total
}

func main() {
	operations := []string{"5", "2", "C", "D", "+"}
	println(calPoints(operations))
}
