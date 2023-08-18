package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "Ups"
	}
}

func main() {

	var data interface{} = Ups(1)

	data2 := Ups(2)

	fmt.Println(data)  
	fmt.Println(data2)  
	fmt.Println(Ups(3))
}
