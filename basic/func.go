package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation : " + op)
	}
}

func div(a, b int) (q, r int) {
	//return a / b, a % b
	q = a / b
	r = a % b
	return q, r
}

// 將函式傳入 名為'op' 的函式 func(int,int) 返回值int 跟 a b int 傳入apply函式內，再轉傳道op內，回傳值為int
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling fuction %s with args (%d,%d)\n", opName, a, b)
	return op(a, b)
}

// a^b
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}
func main() {
	a, b := 3, 4
	a, b = swap(a, b)
	println(a, b)
	//fmt.Println(sum(1, 2, 3, 4, 5))
	//直接進到最裡面看，先將pow整串傳到apply內
	//fmt.Println(apply(func(a, b int) int {
	//	return int(math.Pow(float64(a), float64(b)))
	//}, 3, 4))

	//if result, err := eval(3, 4, "x"); err != nil {
	//	fmt.Println(err)
	//} else {
	//	println(result)
	//}

	//fmt.Println(eval(7, 3, "/"))
	//println(div(7, 3))
	//q, r := div(7, 3)
	//println(q, r)
}
