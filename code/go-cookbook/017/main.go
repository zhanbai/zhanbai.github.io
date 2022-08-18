package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		format := values.Get("format")

		if format != "csv" {
			format = "html"
		}

		sales := [][]string{
			{"Northeast", "2022-01-01", "2022-02-01", "12.54"},
			{"Northwest", "2022-01-01", "2022-02-01", "564.33"},
			{"Southeast", "2022-01-01", "2022-02-01", "93.26"},
			{"Southwest", "2022-01-01", "2022-02-01", "945.21"},
		}

		var headers = []string{"Region", "Start Date", "End Date", "Amount"}
		var totalLine = []string{"All Regions", "--", "--", ""}

		if format == "csv" {
			c, err := os.Create("./sales.csv")
			if err != nil {
				log.Fatalf("failed creating file: %s", err)
			}

			cw := csv.NewWriter(c)
			cw.Write(headers)
			total := 0.0
			// 保存各数据行，将 total 递增
			for _, row := range sales {
				_ = cw.Write(row)
				fs, _ := strconv.ParseFloat(row[3], 64)
				total += fs
			}
			totalLine[3] = strconv.FormatFloat(total, 'E', -1, 32)
			cw.Write(totalLine)
			cw.Flush()
			c.Close()

			// 告诉浏览器这是一个 CSV 文件
			w.Header().Set("Content-Type", "application/csv")
			w.Header().Set("Content-Disposition", `attachment; filename="sales.csv"`)

			f, err := os.Open("./sales.csv")
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}

			io.Copy(w, f)
			f.Close()
		} else {
			s := "<table><tr><th>" + strings.Join(headers, "</th><th>") + "</th></tr>"
			total := 0.0
			for _, row := range sales {
				s += "<tr><td>" + strings.Join(row, "</td><td>") + "</td></tr>"
				fs, _ := strconv.ParseFloat(row[3], 64)
				total += fs
			}
			totalLine[3] = strconv.FormatFloat(total, 'E', -1, 32)
			s += "<tr><td>" + strings.Join(totalLine, "</td><td>") + "</td></tr></table>"
			fmt.Fprintln(w, s)
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
