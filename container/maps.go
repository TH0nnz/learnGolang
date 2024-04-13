package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //m2 == empty map

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Print("Traversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting value")
	courseName := m["course"]
	fmt.Println(courseName) //golang

	causeName := m["cause"]
	fmt.Println(causeName) //沒有附值就是使用 zero value ，String  zero value 就是""

	//如何判斷有沒有這個值在不在map中
	courseName2, ok := m["course"]
	fmt.Println(courseName2, ok) //golang ,true

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key dose not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok) //ccmouse true
	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok) //false

}
