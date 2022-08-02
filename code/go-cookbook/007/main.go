package main

import (
	"fmt"
	"regexp"
	"strings"
)

func tabExpand(t string) string {
	re := regexp.MustCompile(`^([^\t\n]*)(\t+)`)
	for strings.Index(t, "\t") != -1 {
		t = re.ReplaceAllStringFunc(t, tabExpandHelper)
	}
	return t
}

func tabExpandHelper(s string) string {
	ts := 8
	re := regexp.MustCompile(`^([^\t\n]*)`)
	s1 := re.FindString(s)
	re = regexp.MustCompile(`(\t+)`)
	s2 := re.FindString(s)
	return s1 + strings.Repeat(" ", len(s2)*ts-(len(s1)%ts))
}

func tabUnexpand(t string) string {
	ts := 8
	lines := strings.Split(t, "\n")
	for k, v := range lines {
		// 将制表符扩展为空格
		v = tabExpand(v)
		chunkCount := len([]rune(v)) / ts
		line := ""
		// 扫描除最后一个字符块以外的所有其它字符块
		for i := 0; i < chunkCount; i++ {
			re := regexp.MustCompile(` {2,}$`)
			line += re.ReplaceAllString(string([]rune(v)[i*8:i*8+8]), "\t")
		}
		// 如果最后一个字符块是相当于制表符的空格
		// 将它转换为一个制表符；否则，保持不变
		if string([]rune(v)[(chunkCount)*8:chunkCount*8+8]) == strings.Repeat(" ", ts) {
			line += "\t"
		} else {
			line += string([]rune(v)[(chunkCount)*8 : chunkCount*8+8])
		}
		lines[k] = line
	}
	// 重新组合文本行
	return strings.Join(lines, "\n")
}

func main() {
	msg := "白日依山尽  黄河入海流	欲穷千里目	更上一层楼"
	fmt.Println(tabExpand(msg))
	fmt.Println(tabUnexpand(msg))
}
