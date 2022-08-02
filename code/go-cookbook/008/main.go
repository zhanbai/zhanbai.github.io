package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 将字符串的首字母转换为大写
func ucfirst(s string) string {
	for _, v := range s {
		u := string(unicode.ToUpper(v))
		return u + s[len(u):]
	}
	return ""
}

func main() {
	// s := "how do you do today?"
	// fmt.Println(ucfirst(s))
	// s = "the prince of wales"
	// fmt.Println(strings.Title(s))
	// fmt.Println(strings.ToUpper("i'm not yelling!"))
	// fmt.Println(strings.ToLower("<H1>HELLO WORLD</H1>"))
	// fmt.Println(ucfirst("monkey face"))
	// fmt.Println(ucfirst("1 monkey face"))
	// fmt.Println(strings.Title("1 monkey face"))
	// fmt.Println(strings.Title("don't play zone defense against the philadelphia 76-ers"))
	// fmt.Println(ucfirst("macWorld says I should get an iBook"))
	// fmt.Println(strings.Title("eTunaFish.com might buy itunaFish.Com!"))
	fmt.Println(strings.ToUpper(`"since feeling is first" is a poem by e. e. cummings.`))
	fmt.Println(strings.ToLower("I programmed the WOPR and the TRS-80."))
}
