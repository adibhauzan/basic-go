package main

import "fmt"

func main(){

	counter :=0
	name := "adib"

	increment := func () {
		// name = "hauzan" // contoh salah
		name := "hauzan"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}