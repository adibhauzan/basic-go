package main

import "fmt"

func random() interface{} {
	return "Adib"
}

func main() {
	result := random()


	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	case bool:
		fmt.Println("Value", value, "is bool")
	default:
		fmt.Println("Unknown Type of" , value)
	}
	// resultString := result.(string)
	// fmt.Println(resultString)
}