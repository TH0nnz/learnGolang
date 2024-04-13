package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {

	//定義arr的方法
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}          //:= 要給初始值
	arr3 := [...]int{2, 4, 6, 8, 19} //交給編譯器數
	var grid [4][5]bool              //二維陣列
	fmt.Println(arr1, arr2, arr3)    //四個長度為五的陣列
	fmt.Println(grid)

	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	printArray(&arr1)
	printArray(&arr3)
	fmt.Println(arr1, arr3)

}
