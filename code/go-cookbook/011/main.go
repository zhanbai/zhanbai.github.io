package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

func main() {
	sales := [][]string{
		{"Northeast", "2022-01-01", "2022-02-01", "12.54"},
		{"Northwest", "2022-01-01", "2022-02-01", "564.33"},
		{"Southeast", "2022-01-01", "2022-02-01", "93.26"},
		{"Southwest", "2022-01-01", "2022-02-01", "945.21"},
		{"All Regions", "--", "--", "1597.34"},
	}

	// c, err := os.Create("./sales.csv")
	// if err != nil {
	// 	log.Fatalf("failed creating file: %s", err)
	// }

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	for _, row := range sales {
		_ = w.Write(row)
	}
	w.Flush()
	fmt.Println(buf.String())
	// c.Close()
}
