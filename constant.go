// contants was variabel who cant be changed value
package main

import "fmt"

// false
// func main() {
// 	const name string = "Fachry"
// 	fmt.Println(name)
// 	name = "Alifian" // will error because constant cant be changed value
// 	fmt.Println(name)
// }

// true
func main() {
	const name string = "Fachry"
	const age int = 20

	fmt.Println(name + " is " + fmt.Sprint(age) + " years old")
}
