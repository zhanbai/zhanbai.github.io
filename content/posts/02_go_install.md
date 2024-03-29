+++ 
draft = false
date = 2021-12-02T11:06:51+08:00
title = "搭建 Go 开发环境"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

本文介绍 Go 开发环境的搭建。

> 本教程基于 MacOS 系统

## 关于 Go

Go 是一门具有表现力、简洁、干净和高效的，支持高并发的开源静态编译型语言。

## 安装

首先，下载 [Go for Mac](https://go.dev/dl/) 软件包，然后打开按照提示进行安装。Go 发行版将会被安装到 `/usr/local/go` 目录，`/usr/local/go/bin` 也将被加入 `PATH` 环境变量，然后重启终端将会生效。

验证是否安装成功：

```bash
$ go version
```

## 编写程序

推荐使用 [VSCode](https://code.visualstudio.com/) 编写程序，并安装 [Go 扩展](https://marketplace.visualstudio.com/items?itemName=golang.go)。

进入家目录，并创建 `hello` 项目文件夹：

```bash
$ cd
$ mkdir hello
$ cd hello
```

启用依赖项跟踪：

```bash
$ go mod init example/hello
go: creating new go.mod: module example/hello
```

通过编辑器创建 `hello.go` 文件，并编写以下代码：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

运行这个程序：

```bash
$ go run .
Hello, World!
```

到此，你已经完成了 Go 开发环境的搭建。

更多参考官方教程：

* [下载并安装](https://go.dev/doc/install)
* [教程：Go 入门](https://go.dev/doc/tutorial/getting-started)