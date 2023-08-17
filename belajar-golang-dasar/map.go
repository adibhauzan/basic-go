package main

import "fmt"

func main()  {
	
	person := map[string]string{
		"name": "adib hauzan sofyan",
		"address": "bandung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)

	book["title"] = "belajar golang"
	book["author"] = "adib"
	book["wrong"] = "ups"

	fmt.Println(book)
	fmt.Println(book["title"])
	fmt.Println(book["author"])
	fmt.Println(book["wrong"])

	delete(book, "wrong")
	fmt.Println(book)

}