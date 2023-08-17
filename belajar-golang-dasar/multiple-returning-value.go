package main

import "fmt"

func main()  {
	fullname,_ := getFullName()
	fmt.Println(fullname)
}

func getFullName()(string, int){
	return "adib hauzan", 21
}