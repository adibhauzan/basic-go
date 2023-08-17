package main

import (
	"fmt"
)

func main()  {

	// pembuatan constant yang salah
	// const salah string
	// salah = "salah"

	const firstName = "Adib"
	const middleName = "Hauzan"
	const lastName = "Sofyan"
	const age = 21

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(age)

	const (
		fullName = "adib Hauzan Sofyan"
		umur = 21
		country = "indonesia"
		
	)

	fmt.Println(fullName)
	fmt.Println(umur)
	fmt.Println(country)
	
}