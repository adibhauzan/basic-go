package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments : ")
	fmt.Println(args)


	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])
}