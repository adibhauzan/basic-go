package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	adib := Customer{
		Name:    "Adib Hauzan Sofyan",
		Address: "Bandung",
		Age:     21,
	}

	fmt.Println(adib)

	mala := Customer{"Wati Komala Sari", "Bandung",21}
	fmt.Println(mala)
}