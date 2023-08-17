package main

import "fmt"

func main() {
	name := "malaasdjkjasd"

	if name == "Adib" {
		fmt.Println("Hello " + name)
	}else if name == "mala"{
		fmt.Println("Hello " + name)
	} else {
		fmt.Println("Kamu Siapa")
	}


	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Pas")
	}
}