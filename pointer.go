package main1

import (
	"fmt"
)


func main1() {

	type Student struct {
		id int
		name string
		email string
		age int
	}

	st4 := &Student {
		id: 999,
		name: "Robin",
	}

	fmt.Printf("%T", st4)
	fmt.Println()

	
}

/*
	var a string = "origin"
		passByValue(a)
		fmt.Println(a)
		passByRefer(&a)
		fmt.Println(a)

	func passByValue(a string) {
		a = "changed"
	}

	func passByRefer(a *string) {
		*a = "changed"
	}
*/