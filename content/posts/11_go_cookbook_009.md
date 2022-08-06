+++ 
draft = false
date = 2022-08-03T08:55:51+08:00
title = "Go 实例：字符串中的内插函数和表达式"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

希望在一个字符串中包含执行某个函数或表达式的结果。

## 解决方案

如果想要包含的值无法直接放在字符串中，就可以使用字符串连接操作符（+），如下所示：

```go
package main

import "fmt"

func getDiameter() string {
	return "20"
}

func main() {
	bn := 2
	gn := 2
	fmt.Println("You have " + fmt.Sprint(bn+gn) + " children.")
	w := "hello"
	fmt.Println("The word " + w + " is " + fmt.Sprint(len([]rune(w))) + " characters long.")
	p := "20.5"
	fmt.Println("You owe " + p + " immediately.")
	fmt.Println("My circle's diameter is " + getDiameter() + " inches.")
}
```

输出如下：

```bash
You have 4 children.
The word hello is 5 characters long.
You owe 20.5 immediately.
My circle's diameter is 20 inches.
```

## 讨论

无

## 参见

`fmt.Sprint` 的有关文档（https://pkg.go.dev/fmt/#Sprint）。