package main

import (
	"fmt"
	"strconv"
)

func main() {

	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println(nilai16)

	var name = "adib hauzan sofyan"
	var e = name[5]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)

	var umur int = 21
	z := strconv.Itoa(umur)
	fmt.Println(z)
}

// iseng

// package main

// import (
// 	"fmt"
// 	"strconv"
// )
// func main() {
// 	var nilai32 int32 = 100000
// 	var nilai8 int8 = int8(nilai32)

// 	fmt.Println(nilai32)
// 	fmt.Println(nilai8)
// 	fmt.Println(kesimpulan() + " " + name() + ", umur " + strconv.Itoa(umur())+ ", dengan nilai " + strconv.Itoa(int(nilai8)))
// }

// func name() string{
// 	const name = "adib hauzan sofyan"
// 	return name
// }

// func umur() int{
// 	const umur = 21
// 	return umur
// }

// func kesimpulan() string {
// 	return "kamu adalah :"
// }

//end iseng
