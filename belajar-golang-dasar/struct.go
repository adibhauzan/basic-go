package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var adib Customer
	adib.Name = "Adib Hauzan Sofyan"
	adib.Address = "Bandung"
	adib.Age = 21

	fmt.Println(adib.Name)
	fmt.Println(adib.Address)
	fmt.Println(adib.Age)
}