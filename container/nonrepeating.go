package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	//for i, ch := range []byte(s) {
	//	if lastOccurred[ch] >= start {
	//		start = lastOccurred[ch] + 1
	//	}
	//	if i-start+1 > maxLength {
	//		maxLength = i - start + 1
	//	}
	//	lastOccurred[ch] = i
	//}

	for i, ch := range []byte(s) {

		lastI, ok := lastOccurred[ch]

		if ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}
	return maxLength
}
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcabcbb")) //3
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))       //1
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))      //3
	fmt.Println(lengthOfNonRepeatingSubStr(""))            //0
	fmt.Println(lengthOfNonRepeatingSubStr("b"))           //1
	fmt.Println(lengthOfNonRepeatingSubStr("asdfghj"))     //6

}
