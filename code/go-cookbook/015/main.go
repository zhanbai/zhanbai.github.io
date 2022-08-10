package main

import (
	"fmt"
	"regexp"
)

func main() {
	// fmt.Println(strings.Split("My sentence is not very complicated", " "))
	// r := regexp.MustCompile(`\d\. `)
	// fmt.Println(r.Split("my day: 1. get up 2. get dressed 3. eat toast", -1))
	// r = regexp.MustCompile(`(?i) x `)
	// fmt.Println(r.Split("31 inches x 22 inches X 9 inches", -1))
	// fmt.Println(strings.Split("dopey,sleepy,happy,grumpy,sneezy,bashful,doc", ","))
	// fmt.Println(strings.SplitN("dopey,sleepy,happy,grumpy,sneezy,bashful,doc", ",", 5))
	r := regexp.MustCompile(`([+\-\/\d])`)
	fmt.Println(r.FindAllString("3 + 2 / 7 - 9", -1))
	// s := r.FindAllString("3 + 2 / 7 - 9", -1)
	// for _, v := range s {
	// 	fmt.Println("/" + v + "/")
	// }
}
