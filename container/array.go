package main

import "fmt"

func printArr(arr [5]int){
	arr[0] = 100
	for i,v :=range arr {
		fmt.Println(i,v)
	}
}

func printArrPoint(arr *[5]int){
	arr[0] = 100
	for i,v :=range arr {
		fmt.Println(i,v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int {2,4,6,8,10}
	var grid [4][5]int
	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	//遍历
	for i:=0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}
	//method2 遍历
	for i,v :=range arr3 {
		fmt.Println(i,v)
	}
	//不要下标只要数值
	for _,v :=range arr3 {
		fmt.Println(v)
	}

	fmt.Println("printArray(arr1)")
	printArr(arr1)

	fmt.Println("printArray(arr3)")
	printArr(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1,arr3)

	fmt.Println("printArrayPoint(arr1)")
	printArrPoint(&arr1)

	fmt.Println("printArrayPoint(arr3)")
	printArrPoint(&arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1,arr3)
}
