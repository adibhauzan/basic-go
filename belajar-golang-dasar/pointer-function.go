package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

func ChangeAddressToWakanda(alamat *Alamat)  {
	alamat.Country = "Wakanda"
}

func main() {

	var alamat Alamat = Alamat{"Bandung", "West Java", "Indonesia"}
	// ChangeAddressToWakanda(&alamat)
	// fmt.Println(alamat)

	var alamatPointer *Alamat = &alamat
	ChangeAddressToWakanda(alamatPointer)
	fmt.Println(alamat)
}