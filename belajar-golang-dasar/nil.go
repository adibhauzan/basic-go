package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	var orang map[string]string = NewMap("Adib")

	if orang == nil {
		fmt.Println("Data Kosong")
	}else {

		fmt.Println(orang)
	}
	

	person := NewMap("Adib")
	fmt.Println(orang)
	fmt.Println(person)
	fmt.Println(NewMap("adib"))

}