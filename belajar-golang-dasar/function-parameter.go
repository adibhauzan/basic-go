package main

import "fmt"

func main()  {
	sayHelloTo("adib", "hauzan")
	
}

func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}