package basic

import (
	"fmt" 
	"time"
	// "go-module/c4f"
)

func basic() {
	// Declare variable
	var email string = "Helloyeah";

	// Print with format
	fmt.Printf("%v\n\n", email);
	// Print value
	fmt.Printf("value: %v\n", email);
	// Print type
	fmt.Printf("type: %T\n\n", email);

 
	// Declare using shorthand

	fullName := "Ryan Nguyen"
	fmt.Println(fullName)
	fmt.Printf("%T\n\n", fullName)

	var number = 10
	var float = 10.123
	fmt.Printf("%T\n", number)
	fmt.Printf("%T\n\n", float)

	total := sum(1, 2)
	fmt.Printf("%v\n\n", total)

	// slice: is a dynamic version of array
	s := make([]string, 0)
	s = append(s, "a")
	s = append(s, "b")
	fmt.Println(s)
	fmt.Println()

	// map: key - value
	m := make(map[string]int)
	m["age"] = 90
	m["count"] = 1

	fmt.Println(m["age"])
	fmt.Println(m["count"])
	fmt.Println()

	// array
	arr := [1] string{}
	arr[0] = "value1"
	// arr[1] = "value2"
	fmt.Println(arr)
	fmt.Println()

	// OOP
	st := Student {
		firstName: "Ryan",
		lastName: "Nguyen",
		Email: "ryan@",
	}

	st.firstName = "Alice"
	fmt.Println(st)

	e1 := st.getEmail()
	fmt.Println(e1)
	fmt.Println()

	e2 := st.setAndGetEmail("alice@")
	fmt.Println(e2)
	fmt.Println()

	// check again
	// value not changed
	// due to value receiver
	fmt.Println(st.Email)
	fmt.Println()

	e3 := st.setAndGetEmailByPointer("alice_new@")
	fmt.Println(e3)
	fmt.Println()

	// check again
	// value not changed
	// due to value receiver
	fmt.Println(st.Email)
	fmt.Println()
	fmt.Println()

	// goroutine
	c := make(chan int)

	go func() {
		fmt.Println("Go 1")
		c <- 100
	}()
	
	go func() {
		fmt.Println("Go 2")
		v := <- c
		fmt.Printf("Receive %d", v)
		fmt.Println()
	}()

	time.Sleep(time.Second)
	fmt.Println("End of main")

}

// OOP
// struct
type Student struct {
	firstName string
	lastName string
	Email string
}

// receiver method
func (s Student) getEmail() string {
	return s.Email
}

// value receiver
func (s Student) setAndGetEmail(email string) string {
	s.Email = email
	return s.Email
}

// pointer receiver
func (s *Student) setAndGetEmailByPointer(email string) string {
	s.Email = email
	return s.Email
}

// End of OOP


// Function
func sayHello() {
	fmt.Println("sayHello\n")
}

func sum(a int, b int) int {
	return a + b
}

func demoVariable() {
	var email = "ryan@"

	email = "newOne"

	fmt.Println(email)
}
