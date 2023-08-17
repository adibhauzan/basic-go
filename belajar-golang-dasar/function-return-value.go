package main

import "fmt"

func main()  {
	result := getHello("Adib")
	fmt.Println(result)
	fmt.Println(getHello(""))
}

func getHello(name string)string  {
	if name== ""{
		return "Haloo bro"
	}else {
		return "Hallo " + name
	}
}