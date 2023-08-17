package main

import "fmt"

func main() {
	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("adib"))

}

func getGoodBye(name string) string {
	return "GoodBye " + name
}
