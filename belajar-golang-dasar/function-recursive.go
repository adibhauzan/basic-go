package main

import "fmt"

func main()  {
	

	loop := factorialLoop
	loopFactorial := factorialRecursive
	fmt.Println(loop(5))
	fmt.Println(5*4*3*2*1)
	fmt.Println(loopFactorial(5))
}


func factorialLoop(value int) int {
	
	result := 1

	for i:= value; i > 0; i--{
		result *= i
	} 
	
	return result
}

func factorialRecursive(value int)int  {
	if value == 1{
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
	
}

