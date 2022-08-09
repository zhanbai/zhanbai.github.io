+++ 
draft = false
date = 2022-08-04T08:55:51+08:00
title = "Go 实例：去除字符串首尾的空格"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

希望从字符串开头和末尾删除空白符。例如，在验证用户输入之前，可能希望先完成清理。

## 解决方案

可以使用 `strings.TrimLeft`、`strings.TrimRight` 或 `strings.Trim`。`strings.TrimLeft` 方法从字符串开头删除空白符，`strings.TrimRight` 从字符串末尾删除空白符，`strings.Trim` 则会删除字符串开头和末尾的空白符，如下所示：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim(" zipcode ", "  \t\n\r\000\x0B"))
	fmt.Println(strings.TrimRight("text ", "  \t\n\r\000\x0B"))
	fmt.Println(strings.TrimLeft(" name", "  \t\n\r\000\x0B"))
}
```

输出如下：

```bash
zipcode
text
name
```

## 讨论

对于这些使用方法，空白符定义为以下字符：普通空格符、制表符、换行符、回车符、空字节符以及垂直制表符。

从字符串去除首尾的空白符不仅可以节省存储空间，还有利于更准确的显示格式化数据或文本，如 `<pre>` 标记中的文本。如果要对用户输入完成比较，首先需要去除数据首尾的空白符。如果有人在 310000 后面错误的输入了一些空格，并把它作为邮政编码，通过去除字符串首尾的空白符，就不会要求修正这个本不算错误的 “错误”。完成本文比较之前先去除首尾空白符还可以确保完成正确的比较，例如 “salami\n” 等于 “salami”。将字符串数据存储到数据库之前，还可以通过去除首尾的空白符来规范化数据，这也是一个很好的想法。

`strings.Trim` 方法还可以从字符串删除用户指定的其它字符。可以将想要删除的字符作为第二个参数传入这个方法。

## 参见

`strings.TrimLeft` 的有关文档（https://pkg.go.dev/strings/#TrimLeft）。

`strings.TrimRight` 的有关文档（https://pkg.go.dev/strings/#TrimRight）。

`strings.Trim` 的有关文档（https://pkg.go.dev/strings/#Trim）。