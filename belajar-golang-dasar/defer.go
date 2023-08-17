package main

import "fmt"

func main() {
	runApplication()
}

func logging() {
	fmt.Println("Berhasil Memanggil Function")
}

func runApplication(){
	defer logging()
	fmt.Println("Running application")
}