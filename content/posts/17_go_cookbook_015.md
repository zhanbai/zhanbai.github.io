+++ 
draft = false
date = 2022-08-10T08:55:51+08:00
title = "Go 实例：分解字符串"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

需要将一个字符串分解为片段。例如，希望访问用户在一个 `<textarea>` 表单域中输入的各行文本。

## 解决方案

如果各个片段由一个固定的字符串分隔，则使用 `strings.Split` 如下所示：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Split("My sentence is not very complicated", " "))
}
```

输出如下：

```bash
[My sentence is not very complicated]
```

如果需要一个 RE2 语法的正则表达式描述分隔符，则要使用 `Regexp.Split` 如下所示：

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile(`\d\. `)
	fmt.Println(r.Split("my day: 1. get up 2. get dressed 3. eat toast", -1))
}
```

输出如下：

```bash
[my day:  get up  get dressed  eat toast]
```

对于不区分大小写的分隔符匹配，可以对 `regexp.MustCompile` 使用 `(?i)` 标志：

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile(`(?i) x `)
	fmt.Println(r.Split("31 inches x 22 inches X 9 inches", -1))
}
```

输出如下：

```bash
[31 inches 22 inches 9 inches]
```

## 讨论

这里最简单的解决方法是使用 `strings.Split` 。为这个函数传入你的分隔符字符串和要分解的字符串。另外使用 `strings.SplitN` 还可以指定一个可选的限制，指出应当返回多少个元素，如果指定的限制小于可能的字符块个数，最后一个字符块将包含所有余下的字符。如下所示：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.SplitN("dopey,sleepy,happy,grumpy,sneezy,bashful,doc", ",", 5))
}
```

输出如下：

```bash
[dopey sleepy happy grumpy sneezy,bashful,doc]
```

分隔符将由 `strings.Split` 作为直接量来处理。如果指定一个逗号和一个空格作为分隔符，只有在遇到一个逗号后面接一个空格时才会分解字符串，只有逗号或者空格时不会分解。

利用 `regexp` 可以有更大的灵活性。不是以一个字符串直接量作为分隔符，它使用了一个 RE2 语法正则表达式引擎。通过使用 `regexp` 可以充分利用正则表达式的大量技巧。

## 参见

`strings.Split` 的有关文档（https://pkg.go.dev/strings/#Split）。

`strings.SplitN` 的有关文档（https://pkg.go.dev/strings/#SplitN）。

`Regexp.Split` 的有关文档（https://pkg.go.dev/regexp#Regexp.Split）。