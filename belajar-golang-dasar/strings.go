package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Adib Hauzan Sofyan", "Adib"))
	fmt.Println(strings.Contains("Adib Hauzan Sofyan", "Mala"))

	fmt.Println(strings.Split("Adib Hauzan Sofyan", ""))

	fmt.Println(strings.ToLower("ADIB HAUZAN Sofyan"))
	fmt.Println(strings.ToUpper("adib Hauzan Sofyan"))
	fmt.Println(strings.ToTitle("Adib Hauzan Sofyan"))

	fmt.Println(strings.Trim("Budi  Adib Hauzan Sofyan  Budi", " Budi "))
	fmt.Println(strings.ReplaceAll("Adib Adib Adib", "Adib", "Budi"))
	

}