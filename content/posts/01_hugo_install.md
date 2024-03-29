+++ 
draft = false
date = 2021-12-01T14:22:44+08:00
title = "使用 Hugo 和 GitHub Pages 构建免费博客网站"
description = ""
slug = ""
authors = []
tags = []
categories = []
externalLink = ""
series = []
+++

本文介绍如何使用 Hugo 和 Github Pages 来构建、维护和管理免费的博客网站。

> 基于 MacOS 系统，另外前提需要安装 [Git](https://git-scm.com/downloads) 工具。

## 关于 Hugo

Hugo 是基于 Go 语言开发的开源静态站点生成器，具有快速、灵活的特点。

## 安装 Hugo

```bash
$ brew install hugo
```

验证是否安装成功：

```bash
$ hugo version
```

## 创建一个新的网站

```bash
$ hugo new site myblog
```

## 配置主题

本文使用了简单、轻量的 [Coder](https://github.com/luizdepra/hugo-coder/) 主题。

```bash
$ cd myblog
$ git init
$ git submodule add https://github.com/luizdepra/hugo-coder.git themes/hugo-coder
```

编辑根目录下配置文件 `config.toml`，具体参数可以参考 `themes/hugo-coder/exampleSite/config.toml` 例子中的配置。

以下是我使用的比较精简的配置参数：

```toml
baseURL = "http://blog.zhanxiaobai.com/"
defaultContentLanguage = "zh-cn"
title = "展白的博客"
theme = "hugo-coder"

[[languages.zh-cn.menu.main]]
name = "文章"
weight = 1
url = "posts/"

[[languages.zh-cn.menu.main]]
name = "关于"
weight = 2
url = "about/"

[params]
avatarURL = "images/avatar.jpeg"
author = "展白"
info = ["互联网开发，个人技术随想"]
description = "展白的个人博客"
keywords = "blog,developer,personal"
dateFormat = "2006-01-02"
colorScheme = "auto"

[[params.social]]
name = "Github"
icon = "fa fa-2x fa-github"
weight = 1
url = "https://github.com/zhanbai/"

[[params.social]]
name = "RSS"
icon = "fa fa-2x fa-rss"
weight = 2
url = "https://blog.zhanxiaobai.com/index.xml"
rel = "alternate"
type = "application/rss+xml"
```

其中，`about` 页面，参考 `themes/hugo-coder/exampleSite/content/about.md` 文件。

## 创建第一篇文章

```bash
$ hugo new posts/1.md
```

## 启动 Hugo 服务

```bash
$ hugo server -D
```

默认文章中 `draft` 参数是 `true`，代表是草稿状态，并不会部署，所以会使用 `-D` 参数。当完成文章编辑，可以将其改为 `false`。

## 生成待托管的静态文件

```bash
$ hugo -D
```

## 创建 GitHub 仓库

创建基于你的用户名的 `<USERNAME>.github.io` 格式的公开仓库。

## 使用 GitHub Action 自动构建博客

创建 `.github/workflows/gh-pages.yml` 配置文件，内容如下：

```yml
name: github pages

on:
  push:
    branches:
      - main  # Set a branch to deploy
  pull_request:

jobs:
  deploy:
    runs-on: ubuntu-20.04
    steps:
      - uses: actions/checkout@v2
        with:
          submodules: true  # Fetch Hugo themes (true OR recursive)
          fetch-depth: 0    # Fetch all history for .GitInfo and .Lastmod

      - name: Setup Hugo
        uses: peaceiris/actions-hugo@v2
        with:
          hugo-version: 'latest'
          # extended: true

      - name: Build
        run: hugo --minify

      - name: Deploy
        uses: peaceiris/actions-gh-pages@v3
        if: github.ref == 'refs/heads/main'
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./public
```

## 设置 Github pages

设置仓库源分支，通过仓库 Settings > Pages 选项设置 `gh-pages` 为源分支。

## 使用自定义域名

创建文件 `static/CNAME`，内容是你的自定义域名，我这里是 `blog.zhanxiaobai.com`。域名需要做好 CNAME 解析。

通过以上一系列操作，你将完成使用 Hugo 和 GitHub Pages 构建免费博客网站。

更多参考官方教程：

* [快速开始](https://gohugo.io/getting-started/quick-start/)
* [在 GitHub 上托管](https://gohugo.io/hosting-and-deployment/hosting-on-github/)