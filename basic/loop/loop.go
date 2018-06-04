package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(v int) string {
	result :=""
	for ; v>0; v /=2 {
		result = strconv.Itoa(v%2) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever(){
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(72387885),
		convertToBin(0),
	)
    printFile("abc.txt")
	//forever()
}
