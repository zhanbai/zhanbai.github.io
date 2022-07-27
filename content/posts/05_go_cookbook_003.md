+++ 
draft = false
date = 2022-07-27T21:27:51+08:00
title = "Go 实例：替换子串"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

希望用另外一个不同的字符串来替换一个子串。例如，打印一个信用卡号之前，想要对除了后四位以外的部分模糊处理。

## 解决方案

使用 `strings.Replace`，如下所示。

```go
// 将字符串 s 中的前 n 个 old 字符串替换为 new 字符串，如果 n 为 -1 则全部替换
newString := strings.Replace(s, old, new, n)
```

## 讨论

下面封装了一个 `replace` 方法，方便替换从 start 到指定 length 长度的字符，其中 start 和 length 不可为负数。

```go
package main

import (
	"fmt"
	"strings"
)

func replace(s string, new string, start int, length int) string {
	return strings.Replace(s, string([]rune(s)[start:start+length]), new, 1)
}

func main() {
	s := "My pet is a blue dog."
	fmt.Println(replace(s, "fish.", 12, len([]rune(s))-12))
	fmt.Println(replace(s, "green", 12, 4))
	creditCard := "4111 1111 1111 1111"
	fmt.Println(replace(creditCard, "xxxx ", 0, len([]rune(creditCard))-4))
	fmt.Println(replace(s, "Title: ", 0, 0))
}
```

输出如下：

```bash
My pet is a fish.
My pet is a green dog.
xxxx 1111
```

如果 start 和 length 为 0，新字串将插入到 s 的开始位置：

```go
fmt.Println(replace("My pet is a blue dog.", "Title: ", 0, 0))
```

输出如下：

```bash
Title: My pet is a blue dog.
```

如果文本过大，无法一次全部显示，你可能希望显示一部分文本，另外还有一个链接指向其余的文本。在这种情况下，函数 `replace` 就很有用。如下所示显示了一个消息的前 25 个字符，后面还有一个省略号作为链接，可以链接到一个显示更多文本的页面。

```go
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func replace(s string, new string, start int, length int) string {
	return strings.Replace(s, string([]rune(s)[start:start+length]), new, 1)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	id := 1
	message := "希望用另外一个不同的字符串来替换一个子串。例如，打印一个信用卡号之前，想要对除了后四位以外的部分模糊处理。"
	fmt.Fprintf(w, "<a href=\"more-text?id=%d\">%s</a>", id, replace(message, " ...", 25, len([]rune(message))-25))
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
```

示例中引用的 more-text 页面可以使用查询串中传入的消息 ID 来获取完整的消息，并显示这个消息。

## 参见

`strings.Replace` 的有关文档（https://golang.google.cn/pkg/strings/#Replace）。