package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "do something"
	fmt.Println(s, len(s))         //do something 12
	s1 := "做些事情"                   //UTF-8
	fmt.Println(s1, len(s1))       //做些事情 12
	fmt.Printf("%s\n", []byte(s1)) //做些事情
	fmt.Printf("%X\n", []byte(s1)) //E5819AE4BA9BE4BA8BE68385
	fmt.Printf("中文字節 UTF-8編碼：\n")
	for _, b := range []byte(s1) {
		fmt.Printf("%X ", b)
		//E5 81 9A E4 BA 9B E4 BA 8B E6 83 85 -> UTF-8編碼
	}
	fmt.Printf("\n英文字節：\n")
	for _, b := range []byte("abc") {
		fmt.Printf("%X ", b)
		//61 62 63
	}
	fmt.Printf("\n中文字節 unicode 編碼：\n")
	for i, ch := range s1 { //ch us a rune -> int32 : 一個四字節整數，
		//s1 做utf-8 解碼，解出來每個字符再轉unicode，再放在rune這個四字節內
		//E5 81 9A UTF-8 轉 unicode 505A 從s1的0號位取出
		//E4 BA 9B UTF-8 轉 unicode 4E9B 從s1的3號位取出
		fmt.Printf("(%d %X) ", i, ch)
		//(0 505A) (3 4E9B) (6 4E8B) (9 60C5) -> unicode 編碼
	}
	fmt.Printf("\nUTF-8 下的長度：\n")
	fmt.Println("Rune count:", utf8.RuneCountInString(s1))
	bytes := []byte(s1)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //英文size=1 中文=3
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}
