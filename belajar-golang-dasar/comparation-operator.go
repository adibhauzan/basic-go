package main

import "fmt"

func main()  {
	
	// operasi perbandingan string
	var name1 = "adib"
	var name2 = "adib"
	var result bool = name1 == name2

	fmt.Println(result)

	// end operasi perbandingan string

	// operasi perbadingan integer

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 == value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 != value2)


	// end operasi perbadingan integer

}