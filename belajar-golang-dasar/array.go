package main

import "fmt"

func main(){
	var names[3] string

	names[0] = "adib"
	names[1] = "Hauzan"
	names[2] = "Sofyan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])


	var nilai = [3] int{
		90,
		80,
		70,
	}

	fmt.Println(nilai)

	fmt.Println(len(names))
	fmt.Println(len(nilai))
}
