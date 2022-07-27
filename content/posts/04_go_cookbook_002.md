+++ 
draft = false
date = 2022-07-27T09:22:13+08:00
title = "Go 实例：抽取子串"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

希望从字符串中的某个特定位置开始抽取这个字符串的一部分。例如，对于输入到一个表单的用户名，想要得到这个用户名的前 8 个字符。

## 解决方案

使用 `string` 和 `rune` 选择子串，如下所示：

```go
package main

import "fmt"

func main() {
	// 用法：string([]rune(str)[start:start+length])
	username := "zhanbai展白"
	username = string([]rune(username)[:8])
	fmt.Println(username)
}
```

## 讨论

`string` 和 `rune` 结合可以完全不用考虑 unicode 字节问题，一个中文就只占一个数组下标。start 和 length 必须是正整数并且不能超过字符串长度，start 可以忽略，即从 0 开始，以上会返回从 start 开始的 length 个字符。字符中的第一个位置为 0。如果忽略 start+length，会返回从 start 到原字符串末尾的子串，如下所示：

```go
package main

import "fmt"

func main() {
	str := "watch out for that tree"
	str = string([]rune(str)[17:])
	fmt.Println(str)
}
```

输出如下：

```bash
t tree
```

## 参见

`string` 的有关文档（https://golang.google.cn/ref/spec#String_literals）。

`rune` 的有关文档（https://golang.google.cn/ref/spec#Rune_literals）。