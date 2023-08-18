package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi Tidak Boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}

}

func main() {

	hasil, err := Pembagi(100, 0)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hasil", hasil)
	}
}
