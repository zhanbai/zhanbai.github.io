package main

import (
	"fmt"
	"net/http"
	"strings"
)

func replace(s string, new string, start int, length int) string {
	return strings.Replace(s, string([]rune(s)[start:start+length]), new, 1)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	id := 1
	message := "希望用另外一个不同的字符串来替换一个子串。例如，打印一个信用卡号之前，想要对除了后四位以外的部分模糊处理。"
	fmt.Fprintf(w, "<a href=\"more-text?id=%d\">%s</a>", id, replace(message, " ...", 25, len([]rune(message))-25))
}

func main() {
	s := "My pet is a blue dog."
	fmt.Println(replace(s, "fish.", 12, len([]rune(s))-12))
	fmt.Println(replace(s, "green", 12, 4))
	creditCard := "4111 1111 1111 1111"
	fmt.Println(replace(creditCard, "xxxx ", 0, len([]rune(creditCard))-4))
	fmt.Println(replace(s, "Title: ", 0, 0))

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
