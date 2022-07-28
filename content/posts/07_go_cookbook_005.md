+++ 
draft = false
date = 2022-07-28T09:27:51+08:00
title = "Go 实例：按单词或字节反转字符串"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

希望反转一个字符串中的单词或字节。

## 解决方案

自定义 `strrev` 方法按字节反转，如下所示：

```go
package main

import (
	"fmt"
)

// 按字节反转字符串
func strrev(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func main() {
	fmt.Println(strrev("This is not a palindrome."))
}
```

输出如下：

```bash
.emordnilap a ton si sihT
```

要按单词反转，需要根据单词边界分解字符串，反转这些单词，然后再重新组合，如下所示：

```go
package main

import (
	"fmt"
	"strings"
)

// 反转单词数组
func arrayReverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	s := "Once upon a time there was a turtle."
	// 将字符串分解为单词
	words := strings.Split(s, " ")
	// 反转单词数组
	words = arrayReverse(words)
	// 重新构建字符串
	s = strings.Join(words, " ")
	fmt.Println(s)
}
```

输出如下：

```go
turtle. a was there time a upon Once
```

## 讨论

也可以如下所示，按单词反转字符串，代码更加简洁。

```go
s := "Once upon a time there was a turtle."
fmt.Println(strings.Join(arrayReverse(strings.Split(s, " ")), " "))
```

## 参见

`strings.Split` 的有关文档（https://golang.google.cn/pkg/strings/#Split）。

`strings.Join` 的有关文档（https://golang.google.cn/pkg/strings/#Join）。