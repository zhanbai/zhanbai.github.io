package main

import (
	"fmt"
	"strings"
)

// 返回字符串的子串
func substr(s string, start int, length int) string {
	if length > len([]rune(s)) {
		return s
	}
	return s[start : start+length]
}

func main() {
	booklist := `Elmer Gantry             Sinclair Lewis 1927
The Scarlatti InheritanceRobert Ludlum  1971
The Parsifal Mosaic      William Styron 1979`
	bookArr := make(map[int]map[string]string)
	book := make(map[string]string)
	books := strings.Split(booklist, "\n")
	for i := 0; i < len(books); i++ {
		book["title"] = strings.Trim(substr(books[i], 0, 25), " ")
		book["author"] = strings.Trim(substr(books[i], 25, 15), " ")
		book["publication_year"] = strings.Trim(substr(books[i], 40, 4), " ")
		bookArr[i] = book
	}

	fmt.Println(bookArr)
}
