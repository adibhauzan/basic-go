package main

import "fmt"

func main() {
	
	name := "asd"

	switch name {
	case "Adib":
		fmt.Println("Hello Adib")
	case "Mala":
		fmt.Println("Hello Mala")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}
}