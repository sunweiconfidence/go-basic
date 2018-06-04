package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"  //UTF-8
	fmt.Println(s)
	for _,b := range []byte(s) {
		fmt.Printf("%X ",b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d %X) ",i,ch)
	}

	fmt.Println("Rune Count:",utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes)>0 {
		ch,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Println("c ",ch)
	}
	fmt.Println() //换行

	for i ,ch := range []rune(s) {
		fmt.Printf("(%d %c)",i,ch)
	}
	fmt.Println()

}
