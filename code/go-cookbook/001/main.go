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
