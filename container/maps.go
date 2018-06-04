package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) //m2 = empty map
	var m3 map[string]int //m3 =nil

	fmt.Println(m,m2,m3)

	fmt.Println("Travelling map")
	for k,v := range m {
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	courceName, ok := m["course"]
	fmt.Println(courceName,ok)

	if causeName, ok :=m["curse"]; ok {
		fmt.Println(causeName,ok)
	}else {
		fmt.Println("key doesn't exist")
	}

	fmt.Println("Deleting values")
	name , ok := m["name"]
	fmt.Println(name, ok)

	delete(m,"name")
	name, ok = m["name"]
	fmt.Println(name,ok)








}
