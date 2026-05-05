package main

import "fmt"

func main() {
	//dataTypes()
	//validations()
	//labTest()
	scopeTest()
}

func dataTypes() {

	user := "Fer"
	age := 28
	balance := 100.1
	active := true

	//	Go infiere el tipo → más limpio, más usado. Solo usado dentro de func
	fmt.Printf("User: %s, Age: %d, Balance: %.4f, Active: %t\n", user, age, balance, active)
}
func validations() {
	age := 29

	if age >= 18 {
		fmt.Println("Age is greater than 18")
	} else {
		fmt.Println("Age is less than 18")
	}
}

func labTest() {
	user := "Fer"
	active := true
	balance := 50.0

	var result string

	if !active {
		result = "User is not active"
	} else if balance < 100 {
		result = "Balance is insufficient"
	} else {
		result = "User is valid"
	}
	fmt.Printf("User: %s -> %s (balance: %.2f)\n", user, result, balance)
}

func scopeTest() {
	status := "init"

	if true {
		status = "processing"
	}

	fmt.Println(status)
}
