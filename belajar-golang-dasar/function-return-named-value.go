package main

import "fmt"

func main()  {
	fmt.Println(getFullName())
}


func getFullName()(firstName, middleName, lastName string, age int)  {
	firstName = "adib"
	middleName = "hauzan"
	lastName = "sofyan"
	age = 21
	return
}