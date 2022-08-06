+++ 
draft = false
date = 2022-08-06T08:55:51+08:00
title = "Go 实例：解析逗号分隔数据"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

已经有逗号分隔值格式的数据，例如，从 Excel 或数据库导出的一个文件，希望将记录和字段抽取为可以在 Go 中处理的一种格式。

## 解决方案

如果 CSV 数据在一个文件中（或可以通过一个 URL 得到），用 `os.Open` 打开文件，并使用 `csv.NewReader` 读入数据。如下示例将 CSV 数据输出到一个 HTML 表格中。

```go
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
```

## 讨论

`Reader.Read` 会读入一整行数据。不要试图绕过 `csv.NewReader`，而只想读入一行再使用 `strings.Split` 方法按逗号进行解析。CSV 比这要更复杂，它可以处理包含特殊符号的字段值，如上述示例中的 "All Regions"，正确解析结果为 All Regions， 使用 `csv.NewReader` 可以避免这样一些微妙的错误。

## 参见

`os.Open` 的有关文档（https://pkg.go.dev/os/#Open）。

`csv.NewReader` 的有关文档（https://pkg.go.dev/encoding/csv/#NewReader）。

`Reader.Read` 的有关文档（https://pkg.go.dev/encoding/csv#Reader.Read）。