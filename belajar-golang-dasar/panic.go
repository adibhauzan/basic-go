package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	fmt.Println("Terjadi Error",  message)
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("aplikasi error")
	}

	fmt.Println("aplikasi berjalan")
}