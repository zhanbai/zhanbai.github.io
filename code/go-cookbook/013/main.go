package main

import "fmt"

// 返回字符串的子串
func substr(s string, start int, length int) string {
	if length > len([]rune(s)) {
		return s
	}
	return s[start : start+length]
}

// 使用另一个字符串填充字符串为指定长度
func strPad(input string, padLength int, padString string) string {
	output := ""
	inputLen := len(input)
	if inputLen >= padLength {
		return input
	}
	for i := 0; i < (padLength - inputLen); i = i + len(padString) {
		output += padString
	}
	return input + output
}

func main() {
	books := [][]string{
		{"Elmer Gantry", "Sinclair Lewis", "1927"},
		{"The Scarlatti Inheritance", "Robert Ludlum", "1971"},
		{"The Parsifal Mosaic", "William Styron", "1979"},
	}

	for _, book := range books {
		title := strPad(substr(book[0], 0, 25), 25, ".")
		author := strPad(substr(book[1], 0, 15), 15, ".")
		year := strPad(substr(book[2], 0, 4), 4, ".")
		fmt.Println(title + author + year)
	}
}
