+++ 
draft = false
date = 2022-07-28T06:27:51+08:00
title = "Go 实例：逐字节处理字符串"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

## 问题

需要分别处理字符串中的各个字节。

## 解决方案

使用 `for` 循环处理字符串中的各个字节，如下所示可以统计一个字符串中的元音字母个数。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This weekend, I'm going shopping for a pet chicken."
	vowels := 0
	for _, v := range s {
		if strings.Index("aeiouAEIOU", string(v)) != -1 {
			vowels++
		}
	}
	fmt.Println(vowels)
}
```

输出如下：

```bash
14
```

## 讨论

通过一次处理字符串中的一个字符，可以很容易的计算 “Look and Say” 序列（J.H.Conway 提出的一个整数序列），如下所示：

```go
package main

import (
	"fmt"
	"strconv"
)

func lookandsay(s string) string {
	// 将返回值初始化为一个空串
	r := ""
	// m 包含要统计的字符，初始化为字符串中的第一个字符
	m := s[0]
	// n 是已经查看过的 m 个数，初始化为 1
	n := 1
	for i := 1; i < len([]rune(s)); i++ {
		// 如果这个字符与上一个相同
		if s[i] == m {
			// 将这个字符的个数增加 1
			n++
		} else {
			// 否则，将字符个数和字符本身增加到返回值
			r += strconv.Itoa(n) + string(m)
			// 将要查找的字符设置为当前字符
			m = s[i]
			// 并将字符个数重置为 1
			n = 1
		}

	}
	return r + strconv.Itoa(n) + string(m)

}

func main() {
	s := "1"
	for i := 0; i < 10; i++ {
		s = lookandsay(s)
		fmt.Println(s)
	}
}

```

输出如下：

```bash
11
21
1211
111221
312211
13112221
1113213211
31131211131221
13211311123113112211
11131221133112132113212221
```

这称为 “Look and Say” 序列，因为每个元素是通过查看前一个元素并说出其中是什么来得到的。例如，查看第一个元素 1，你会说 “一个 1”，所以下一个元素是 “11”。这就是两个 1，所以下一个元素是 “21”。类似的，这里有一个 2 和一个 1，所以下一个元素是 “1211”，以此类推。

## 参见

`for` 的有关文档（https://golang.google.cn/ref/spec#For_statements）。