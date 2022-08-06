package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./sales.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	r := csv.NewReader(f)

	fmt.Println("<table>")

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("<tr>")

		for _, v := range record {
			fmt.Print("<td>" + v + "</td>")
		}

		fmt.Println("</tr>")
	}

	fmt.Println("</table>")
}
