+++ 
draft = false
date = 2022-07-29T06:05:51+08:00
title = "Go 实例：生成随机字符串"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

希望生成一个随机字符串。

## 解决方案

自定义 `strRand` 方法，如下所示：

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机字符串
func strRand(l int, chars string) string {
	if chars == "" {
		chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	charsLen := len(chars) - 1
	str := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := l; i > 0; i-- {
		str += string(chars[r.Intn(charsLen)])
	}

	return str
}

func main() {
	fmt.Println(strRand(32, ""))
}
```

## 讨论

Go 提供了生成随机数的内置方法，不过没有生成随机字符串的方法。`strRand` 方法返回一个由字母和数字构成的字符串。

传入一个整数可以改变返回字符串的长度。如果要使用另一个字符集合，可以相应地传入一个字符串作为第二个参数。例如，可以如下得到一个 16 位摩斯电码：

```go
fmt.Println(strRand(16, ".-"))
```

输出如下：

```bash
.--.-.-----....-
```

## 参见

`rand` 的有关文档（https://golang.google.cn/pkg/math/rand/）。