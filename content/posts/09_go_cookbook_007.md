+++ 
draft = false
date = 2022-08-01T08:05:50+08:00
title = "Go 实例：扩展和压缩制表符"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

希望将字符串中的空格改为制表符（或者将制表符改为空格），但仍保证文本按制表位对齐。例如，可能想以一种标准化的方式为用户显示格式化文本。

## 解决方案

使用 `strings.Replace` 将空格替换为制表符，或者将制表符替换为空格，如下所示：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "白日依山尽 黄河入海流	欲穷千里目 更上一层楼"
	tabbed := strings.Replace(msg, " ", "\t", -1)
	spaced := strings.Replace(msg, "\t", " ", -1)
	fmt.Println("With Tabs: " + tabbed)
	fmt.Println("With Spaces: " + spaced)
}
```

输出如下：

```
With Tabs: 白日依山尽   黄河入海流      欲穷千里目      更上一层楼
With Spaces: 白日依山尽 黄河入海流 欲穷千里目 更上一层楼
```

不过，如果使用 `strings.Replace` 来完成替换，并不能保证按制表位对齐。如果希望每 8 个字符一个制表位，假设一行文本开头有一个 5 字母的单词和一个制表符，这个制表符就应当替换为 3 个空格，而不是一个空格。可以使用如下示例的 `tabExpand` 函数将制表符替换为空格，同时保证按制表位对齐。

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func tabExpand(t string) string {
	re := regexp.MustCompile(`^([^\t\n]*)(\t+)`)
	for strings.Index(t, "\t") != -1 {
		t = re.ReplaceAllStringFunc(t, tabExpandHelper)
	}
	return t
}

func tabExpandHelper(s string) string {
	ts := 8
	re := regexp.MustCompile(`^([^\t\n]*)`)
	s1 := re.FindString(s)
	re = regexp.MustCompile(`(\t+)`)
	s2 := re.FindString(s)
	return s1 + strings.Repeat(" ", len(s2)*ts-(len(s1)%ts))
}

func main() {
	msg := "白日依山尽	   黄河入海流	欲穷千里目	更上一层楼"
	fmt.Println(tabExpand(msg))
}
```

可以使用如下示例的 `tabUnexpand` 函数将空格转换回制表符。

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func tabExpand(t string) string {
	re := regexp.MustCompile(`^([^\t\n]*)(\t+)`)
	for strings.Index(t, "\t") != -1 {
		t = re.ReplaceAllStringFunc(t, tabExpandHelper)
	}
	return t
}

func tabExpandHelper(s string) string {
	ts := 8
	re := regexp.MustCompile(`^([^\t\n]*)`)
	s1 := re.FindString(s)
	re = regexp.MustCompile(`(\t+)`)
	s2 := re.FindString(s)
	return s1 + strings.Repeat(" ", len(s2)*ts-(len(s1)%ts))
}

func tabUnexpand(t string) string {
	ts := 8
	lines := strings.Split(t, "\n")
	for k, v := range lines {
		v = tabExpand(v)
		chunkCount := len([]rune(v)) / ts
		line := ""
		for i := 0; i < chunkCount; i++ {
			re := regexp.MustCompile(` {2,}$`)
			line += re.ReplaceAllString(string([]rune(v)[i*8:i*8+8]), "\t")
		}
		if string([]rune(v)[(chunkCount)*8:chunkCount*8+8]) == strings.Repeat(" ", ts) {
			line += "\t"
		} else {
			line += string([]rune(v)[(chunkCount)*8 : chunkCount*8+8])
		}
		lines[k] = line
	}
	return strings.Join(lines, "\n")
}

func main() {
	msg := "白日依山尽  黄河入海流	欲穷千里目	更上一层楼"
	fmt.Println(tabUnexpand(msg))
}
```

这两个函数都取一个字符串作为参数，并返回适当修改后的字符串。

## 讨论

这两个函数都假设每 8 个空格一个制表位，不过可以通过改变 `ts` 变量的设置来重新设定。

`tabExpand` 中的正则表达式可以匹配一组制表符，也可以匹配文本行中这组制表符之前的所有文本。之所以需要匹配制表符前的文本，是因为这个文本的长度会影响制表符要替换为多少个空格，从而保证后续的文本能够与下一个制表位对齐。这个函数并不只是把每个制表符替换为 8 个空格，它会调整制表符后的文本，与制表位对齐。

类似地，`tabUnexpand` 不只是查看 8 个连续的空格并将它们替换为一个制表符。它会把各个文本行分解为 8 字符的字符块，然后将这些字符块中的末尾空白符（至少两个空格）替换为制表符。这样不仅可以保证文本与制表位对齐，还可以节省字符串空间。

## 参见

`strings.Replace` 的有关文档（https://golang.google.cn/pkg/strings/#Replace）。

`Regexp.ReplaceAllStringFunc` 的有关文档（https://golang.google.cn/pkg/regexp/#Regexp.ReplaceAllStringFunc）。

`Regexp.ReplaceAllString` 的有关文档（https://golang.google.cn/pkg/regexp/#Regexp.ReplaceAllString）。