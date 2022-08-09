+++ 
draft = false
date = 2022-08-09T08:55:51+08:00
title = "Go 实例：解析固定宽度字段数据记录"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

需要将固定宽度记录分解为字符串。

## 解决方案

使用自定义函数 `substr`，如下所示：

```go
package main

import (
	"fmt"
	"strings"
)

// 返回字符串的子串
func substr(s string, start int, length int) string {
	if length > len([]rune(s)) {
		return s
	}
	return s[start : start+length]
}

func main() {
	booklist := `Elmer Gantry             Sinclair Lewis 1927
The Scarlatti InheritanceRobert Ludlum  1971
The Parsifal Mosaic      William Styron 1979`
	bookArr := make(map[int]map[string]string)
	book := make(map[string]string)
	books := strings.Split(booklist, "\n")
	for i := 0; i < len(books); i++ {
		book["title"] = strings.Trim(substr(books[i], 0, 25), " ")
		book["author"] = strings.Trim(substr(books[i], 25, 15), " ")
		book["publication_year"] = strings.Trim(substr(books[i], 40, 4), " ")
		bookArr[i] = book
	}
	fmt.Println(bookArr)
}
```

## 讨论

上述图书列表数据中，每一行会为每个字段分配固定数目的字符，包括书名、作者和出版日期。每一行上，书名占前 25 个字符，作者姓名占接下来 15 个字符，出版年代则占后面的 4 个字符。知道了这些字段宽度，就能很容使用 `substr` 将这些字段解析为一个二维 map。通过将 booklist 分解为一个文本行切片，这样不论是处理一个字符串还是处理从文件读入的一系列文本行，循环代码都是一样的。

## 参见

`strings.Split` 的有关文档（https://pkg.go.dev/strings/#Split）。

`strings.Trim` 的有关文档（https://pkg.go.dev/strings/#Trim）。