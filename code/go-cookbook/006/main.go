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
	charsLen := len(chars)
	str := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := l; i > 0; i-- {
		str += string(chars[r.Intn(charsLen)])
	}

	return str
}

func main() {
	fmt.Println(strRand(32, ""))
	fmt.Println(strRand(16, ".-"))
}
