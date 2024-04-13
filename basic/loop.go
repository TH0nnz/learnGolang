package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func forever() {
	for {
		fmt.Println("abe")
	}
}
func main() {
	//	fmt.Println(
	//		convertBin(5),  //101 -> 4+1=5
	//		convertBin(13), //1101 -> 8+4+1=13
	//		convertBin(3424235434543234),
	//	)
	//readFile("abc.txt")
	forever()
}
