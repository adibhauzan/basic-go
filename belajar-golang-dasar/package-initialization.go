package main

import (
	"belajar-golang-dasar/database"
	_"belajar-golang-dasar/helper" // blank indentifier 
	"fmt"
)

func main() {

	result := database.GetDatabase()
	fmt.Println(result)
}