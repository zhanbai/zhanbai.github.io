+++ 
draft = false
date = 2022-08-12T08:55:51+08:00
title = "Go 实例：使文本在指定行长度自动换行"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

需要实现字符串自动换行。例如，希望使用 `<pre>` 和 `</pre>` 标记显示文本，但要求保持在一个常规大小的浏览器窗口以内。

## 解决方案

使用自定义函数 `wordwrap` 如下所示：

```go
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// 自动换行
func wordwrap(text string, lineWidth int) string {
	wrap := make([]byte, 0, len(text)+2*len(text)/lineWidth)
	eoLine := lineWidth
	inWord := false
	for i, j := 0, 0; ; {
		r, size := utf8.DecodeRuneInString(text[i:])
		if size == 0 && r == utf8.RuneError {
			r = ' '
		}
		if unicode.IsSpace(r) {
			if inWord {
				if i >= eoLine {
					wrap = append(wrap, '\n')
					eoLine = len(wrap) + lineWidth
				} else if len(wrap) > 0 {
					wrap = append(wrap, ' ')
				}
				wrap = append(wrap, text[j:i]...)
			}
			inWord = false
		} else if !inWord {
			inWord = true
			j = i
		}
		if size == 0 && r == ' ' {
			break
		}
		i += size
	}
	return string(wrap)
}

func main() {
	s := "Four score and seven years ago our fathers brought forth on this continent a new nation, conceived in liberty and dedicated to the proposition that all men are created equal."
	fmt.Println(wordwrap(s, 75))
}
```

输出如下：

```bash
Four score and seven years ago our fathers brought forth on this continent
a new nation, conceived in liberty and dedicated to the proposition that
all men are created equal.
```

## 讨论

自定义函数 `wordwrap` 第二个参数可以指定一个行长度。

## 参见

`utf8.DecodeRuneInString` 的有关文档（https://pkg.go.dev/unicode/utf8#DecodeRuneInString）。