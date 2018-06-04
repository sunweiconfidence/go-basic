package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

//func eval(a,b int,op string) int {
//	switch op {
//	case "+":
//		return a + b
//	case "-":
//	 return a-b
//	case "*":
//		return a*b
//	case "/":
//		//return a/b
//		q, _ :=div(a,b)
//		return q
//	default:
//		//panic(fmt.Sprintf("invalid operator",op))
//		panic("unsupport operation:"+op)
//	}
//}
//two return value
func eval(a,b int,op string) (int,error) {
	switch op {
	case "+":
		return a + b,nil
	case "-":
	 return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		//return a/b
		q, _ :=div(a,b)
		return q,nil
	default:
		//panic(fmt.Sprintf("invalid operator",op))
		//panic("unsupport operation:"+op)
		return 0,fmt.Errorf("unsupport operation:%s",op)
	}
}

//func div(a,b int) (int,int){
//	return a/b, a%b
//}

func div(a,b int) (q,r int){
	//q= a/b
	//r = a%b
	//return

	//best write method
	return a/b, a%b
}

func apply(op func(int,int) int,a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d,%d)\n",opName,a,b)
	return op(a,b)
}

func sum(numbers ...int) int{
	s :=0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func pow(a,b int) int {
	return int(math.Pow(float64(a),float64(b)))
}

func swap(a,b *int){
	*b,*a = *a,*b
}

//method2
func swap2(a,b int) (int,int){
	return b,a
}

func main() {
  //fmt.Println(eval(1,10,"+"))
  //fmt.Println(eval(3,4,"/"))
  //fmt.Println(div(13,3))
  if result, err :=eval(3,4,"x"); err!=nil{
  	fmt.Println("Error:",err)
  }else {
  	fmt.Println(result)
  }
  q, r :=div(13,3)
  fmt.Println(q,r)

  //write method1
  //fmt.Println(apply(pow,3,4))

  //write method2
  fmt.Println(apply(func(a int,b int) int {
	  return int(math.Pow(float64(a), float64(b)))
  },3,4))

  fmt.Println(sum(1,2,3,4,5))

	a , b :=3,4
	swap(&a,&b)
	fmt.Println(a,b)

	c,d:=3,4
	c,d = swap2(c,d)
	fmt.Println(c,d)
}
