package main

import "fmt"

func main(){

	type NoKtp string

	var ktpAdib NoKtp = "12354"
	fmt.Println(ktpAdib)
}

// iseng

// func main() {
// 	fmt.Println("Name : " + name())
// 	fmt.Println("No Ktp : " + ktp())
// }

// func name() string {
// 	type Name string

// 	var adib Name = "adib hauzan sofyan"
// 	return string(adib)
// }

// func ktp() string {
// 	type NoKtp string

// 	var ktpAdib NoKtp = "12354"
// 	return string(ktpAdib)
// }

// end iseng
