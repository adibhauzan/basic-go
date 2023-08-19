package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])a")

	fmt.Println(regex.MatchString("ada"))
	fmt.Println(regex.MatchString("ama"))
	fmt.Println(regex.MatchString("uda"))

	
}
