package main

import (
	"fmt"
	"strings"
)

// 按字节反转字符串
func strrev(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

// 反转单词数组
func arrayReverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	fmt.Println(strrev("This is not a chinese."))

	s := "Once upon a time there was a turtle."
	fmt.Println(strings.Join(arrayReverse(strings.Split(s, " ")), " "))
	// 将字符串分解为单词
	// words := strings.Split(s, " ")
	// 反转单词数组
	// words = arrayReverse(words)
	// 重新构建字符串
	// s = strings.Join(words, " ")
	// fmt.Println(s)
}
