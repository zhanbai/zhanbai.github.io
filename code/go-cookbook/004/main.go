package main

import (
	"fmt"
	"strconv"
	"strings"
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
	s := "This weekend, I'm going shopping for a pet chicken."
	vowels := 0
	for _, v := range s {
		if strings.Index("aeiouAEIOU", string(v)) != -1 {
			vowels++
		}
	}
	fmt.Println(vowels)

	// s := "1"
	// for i := 0; i < 10; i++ {
	// 	s = lookandsay(s)
	// 	fmt.Println(s)
	// }
}
