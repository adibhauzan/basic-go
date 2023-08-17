package main

import (
	"fmt"

)

func main() {

	total := sumALl(10, 90, 30, 50, 40)
	fmt.Println(total)

	slice := []int{10,10,10,10}
	total = sumALl(slice...)
	fmt.Println(total)

}

func sumALl(numbers ...int) int {

	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}