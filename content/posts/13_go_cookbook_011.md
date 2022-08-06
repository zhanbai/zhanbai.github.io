+++ 
draft = false
date = 2022-08-05T08:55:51+08:00
title = "Go 实例：生成逗号分隔数据"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

希望将数据格式化为逗号分隔值（comma-separated values, CSV），从而可以由电子表格或数据库导入。

## 解决方案

可以使用 `encoding/csv` 标准库的 `Write` 方法根据一个数据数组生成一个 CSV 格式的文本行，如下所示把 sales 中的数据写入一个文件。

```go
package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	sales := [][]string{
		{"Northeast", "2022-01-01", "2022-02-01", "12.54"},
		{"Northwest", "2022-01-01", "2022-02-01", "564.33"},
		{"Southeast", "2022-01-01", "2022-02-01", "93.26"},
		{"Southwest", "2022-01-01", "2022-02-01", "945.21"},
		{"All Regions", "--", "--", "1597.34"},
	}

	c, err := os.Create("./sales.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w := csv.NewWriter(c)
	for _, row := range sales {
		_ = w.Write(row)
	}

	w.Flush()
	c.Close()
}
```

## 讨论

要输出 CSV 格式的数据而不是写入文件，可以使用标准输出 `os.Stdout`，如下所示：

```go
package main

import (
	"encoding/csv"
	"os"
)

func main() {
	sales := [][]string{
		{"Northeast", "2022-01-01", "2022-02-01", "12.54"},
		{"Northwest", "2022-01-01", "2022-02-01", "564.33"},
		{"Southeast", "2022-01-01", "2022-02-01", "93.26"},
		{"Southwest", "2022-01-01", "2022-02-01", "945.21"},
		{"All Regions", "--", "--", "1597.34"},
	}

	w := csv.NewWriter(os.Stdout)
	for _, row := range sales {
		_ = w.Write(row)
	}

	w.Flush()
}
```

要把 CSV 格式的数据放入一个字符串，而不是输出或写至一个文件，可以结合输出缓冲区来实现，如下所示：

```go
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

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	for _, row := range sales {
		_ = w.Write(row)
	}
	w.Flush()
	fmt.Println(buf.String())
}
```
## 参见

`os.Create` 的有关文档（https://pkg.go.dev/os/#Create）。

`csv.Write` 的有关文档（https://pkg.go.dev/encoding/csv/#Writer）。

`bytes.Buffer` 的有关文档（https://pkg.go.dev/bytes/#Buffer）。