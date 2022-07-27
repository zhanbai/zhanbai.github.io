package main

import "fmt"

func main() {
	// 用法：string([]rune(str)[start:start+length])
	username := "zhanbai展白"
	username = string([]rune(username)[:8])
	fmt.Println(username)

	s := "watch out for that tree"
	s = string([]rune(s)[17:])
	fmt.Println(s)
}
