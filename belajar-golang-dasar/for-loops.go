package main

import "fmt"

func main(){
	
	counter1 := 1

	for counter1  <= 10{
		fmt.Println("Perulangan ke", counter1)
		counter1++
	}

	for counter2 := 1; counter2 <= 10; counter2++{
		fmt.Println("Perulangan ke", counter2)
	}

	slice := []string{"Adib", "Hauzan", "Sofyan"}

	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index ", i, "=", value)
	}
}