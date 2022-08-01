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
		v = tabExpand(v)
		chunkCount := len([]rune(v)) / ts
		line := ""
		for i := 0; i < chunkCount; i++ {
			re := regexp.MustCompile(` {2,}$`)
			line += re.ReplaceAllString(string([]rune(v)[i*8:i*8+8]), "\t")
		}
		if string([]rune(v)[(chunkCount)*8:chunkCount*8+8]) == strings.Repeat(" ", ts) {
			line += "\t"
		} else {
			line += string([]rune(v)[(chunkCount)*8 : chunkCount*8+8])
		}
		lines[k] = line
	}
	return strings.Join(lines, "\n")
}

func main() {
	msg := "白日依山尽  黄河入海流	欲穷千里目	更上一层楼"
	fmt.Println(tabExpand(msg))
	fmt.Println(tabUnexpand(msg))
}
