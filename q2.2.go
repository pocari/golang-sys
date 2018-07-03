package main

import (
	"encoding/csv"
	"os"
)

func Q2_2() {
	w := csv.NewWriter(os.Stdout)
	w.Write([]string{"1-1", "1-2", "1-3"})
	w.Flush()
}
