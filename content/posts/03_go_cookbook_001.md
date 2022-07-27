+++ 
draft = false
date = 2022-07-27T07:05:51+08:00
title = "Go 实例：访问子串"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

想知道一个字符串是否包含一个特定的子串。例如，想查看一个 email 地址是否包含一个 @。

## 解决方案

使用 `strings.Index` 如下所示：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	email := "bai.zhan@qq.com"
	if strings.Index(email, "@") != -1 {
		fmt.Println("There was @ in the e-mail address!")
	}
}
```

## 讨论

`strings.Index` 的返回值是子串（就像“要捞的针”）在字符串（就像“大海”）中出现的第一个位置。如果大海（字符串）中根本没有找到针（子串），`strings.Index` 将返回 -1 。如果子串位于这个字符串的起始位置，`strings.Index` 将返回 0，因为位置 0 表示字符串的开始位置。

## 参见

`strings.Index` 的有关文档，https://golang.google.cn/pkg/strings/#Index 。