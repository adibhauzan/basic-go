package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bandung", "West Java", "Indonesia"}
	address2 := address1

	address2.City = "Padang"
	address2.Province= "West Sumatra"

	fmt.Println(address1)
	fmt.Println(address2)
}