+++ 
draft = false
date = 2022-08-02T08:55:51+08:00
title = "Go 实例：控制大小写"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

需要将字符串中的字母改为全大写或全小写，或者改变字母的大小写。例如，希望名字中首字母大写，而其余字母都为小写。

## 解决方案

使用自定义函数 `ucfirst` 或 `strings.Title` 将一个或多个单词中的首字母大写，如下所示：

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 将字符串的首字母转换为大写
func ucfirst(s string) string {
	for _, v := range s {
		u := string(unicode.ToUpper(v))
		return u + s[len(u):]
	}
	return ""
}

func main() {
	s := "how do you do today?"
	fmt.Println(ucfirst(s))
	s = "the prince of wales"
	fmt.Println(strings.Title(s))
}
```

输出如下：

```bash
How do you do today?
The Prince Of Wales
```

使用 `strings.ToUpper` 或 `strings.ToLower` 可以修改整个字符串的大小写，如下所示：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("i'm not yelling!"))
	fmt.Println(strings.ToLower("<H1>HELLO WORLD</H1>"))
}
```

输出如下：

```bash
I'M NOT YELLING!
<h1>hello world</h1>
```

## 讨论

使用 `ucfirst` 将字符串中的首字母改为大写：

```go
fmt.Println(ucfirst("monkey face"))
fmt.Println(ucfirst("1 monkey face"))
```

输出如下：

```bash
Monkey face
1 monkey face
```

注意，第二个短语并不是 “1 Monkey face”。

使用 `strings.Title` 将一个字符串中各个单词的首字母改为大写：

```go
fmt.Println(strings.Title("1 monkey face"))
fmt.Println(strings.Title("don't play zone defense against the philadelphia 76-ers"))
```

输出如下：

```bash
1 Monkey Face
Don'T Play Zone Defense Against The Philadelphia 76-Ers
```

`strings.Title` 将 “don't” 中的 “t” 改为大写。它也将 “76-ers” 中的 “e” 改为大写。它用于单词边界的规则不能正确处理 Unicode 标点符号，推荐使用 golang.org/x/text/cases 代替。

函数 `strings.ToUpper` 和 `strings.ToLower` 作用于整个字符串，而不只是单个字符。`strings.ToUpper` 将所有字母字符改为大写，`strings.ToLower` 将所有字母字符改为小写：

```go
fmt.Println(strings.ToUpper(`"since feeling is first" is a poem by e. e. cummings.`))
fmt.Println(strings.ToLower("I programmed the WOPR and the TRS-80."))
```

输入如下：

```bash
"SINCE FEELING IS FIRST" IS A POEM BY E. E. CUMMINGS.
i programmed the wopr and the trs-80.
```

## 参见

`strings.ToUpper` 的有关文档（https://pkg.go.dev/strings/#ToUpper）。

`strings.ToLower` 的有关文档（https://pkg.go.dev/strings/#ToLower）。

`strings.Title` 的有关文档（https://pkg.go.dev/strings/#Title）。