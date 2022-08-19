+++ 
draft = false
date = 2022-08-18T08:55:51+08:00
title = "Go 实例：可下载的 CSV 文件"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

需要完成数据格式化，并将 CSV 文件发送到浏览器，并自动保存到一个电子表格中。

## 解决方案

结合使用 http 库和 csv 库的相关方法，可以完成数据格式化，并将 CSV 文件发送到浏览器，并自动保存到一个电子表格中。如下所示：

```go
package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sales := [][]string{
			{"Northeast", "2022-01-01", "2022-02-01", "12.54"},
			{"Northwest", "2022-01-01", "2022-02-01", "564.33"},
			{"Southeast", "2022-01-01", "2022-02-01", "93.26"},
			{"Southwest", "2022-01-01", "2022-02-01", "945.21"},
		}

		c, err := os.Create("./sales.csv")
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		cw := csv.NewWriter(c)
		cw.Write([]string{"Region", "Start Date", "End Date", "Amount"})
		total := 0.0
		// 保存各数据行，将 total 递增
		for _, row := range sales {
			_ = cw.Write(row)
			fs, _ := strconv.ParseFloat(row[3], 64)
			total += fs
		}
		cw.Write([]string{"All Regions", "--", "--", strconv.FormatFloat(total, 'E', -1, 32)})
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
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

其中，设置了两个头部请求参数确保浏览器正确处理 CSV 输出。第一个参数 Content-Type 告诉浏览器这个输出不是 HTML，而是 CSV。第二个参数 Content-Disposition 告诉浏览器不要显示这个输出，而需尝试加载一个外部程序来处理，其中 filename 属性提供了一个默认文件名，浏览器会对下载文件使用这个默认文件名。

## 讨论

如果希望提供同一数据的不同视图，可以通过传参来确定对数据采用哪一种格式化形式。比如：通过 format 查询字段，确定结果为 HTML 表格还是 CSV 文件，如下所示：

```go
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
```

如果查询中指定 format=csv，程序会返回 CSV 格式文件。否则返回 HTML 输出。这里只是将 format 设置为 CSV 或 HTML，按照这个逻辑可以很容易扩展到其它输出格式，如 JSON。如果很多地方都要用到这个功能，希望用多种格式输出相同的数据，可以把上述示例中的代码打包到一个函数中，这个函数接收一个要输出的数据和一个格式参数，然后显示正确的结果。

## 参见

`strconv.ParseFloat` 的有关文档（https://pkg.go.dev/strconv#ParseFloat）。

`http.ListenAndServe` 的有关文档（https://pkg.go.dev/net/http#ListenAndServe）。