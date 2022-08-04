package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim(" zipcode ", "  \t\n\r\000\x0B"))
	fmt.Println(strings.TrimRight("text ", "  \t\n\r\000\x0B"))
	fmt.Println(strings.TrimLeft(" name", "  \t\n\r\000\x0B"))
}
