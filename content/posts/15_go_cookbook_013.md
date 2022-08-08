+++ 
draft = false
date = 2022-08-08T08:55:51+08:00
title = "Go 实例：生成固定宽度字段数据记录"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

需要格式化数据记录，使得每个字段占据指定数目的字符。

## 解决方案

使用自定义函数 `substr` 和 `strPad` 将一个记录数组转换为固定宽度的记录，其中包含用 . 填充的字段。如下所示：

```go
package main

import "fmt"

// 返回字符串的子串
func substr(s string, start int, length int) string {
	if length > len([]rune(s)) {
		return s
	}
	return s[start : start+length]
}

// 使用另一个字符串填充字符串为指定长度
func strPad(input string, padLength int, padString string) string {
	output := ""
	inputLen := len(input)
	if inputLen >= padLength {
		return input
	}
	for i := 0; i < (padLength - inputLen); i = i + len(padString) {
		output += padString
	}
	return input + output
}

func main() {
	books := [][]string{
		{"Elmer Gantry", "Sinclair Lewis", "1927"},
		{"The Scarlatti Inheritance", "Robert Ludlum", "1971"},
		{"The Parsifal Mosaic", "William Styron", "1979"},
	}

	for _, book := range books {
		title := strPad(substr(book[0], 0, 25), 25, ".")
		author := strPad(substr(book[1], 0, 15), 15, ".")
		year := strPad(substr(book[2], 0, 4), 4, ".")
		fmt.Println(title + author + year)
	}
}
```

## 讨论

自定义函数 `substr` 来确保字段值不会过长，另外使用自定义函数 `strPad` 来确保字段值不会过短，其中使用 . 作为填充字段。

## 参见

无