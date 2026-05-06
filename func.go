package main

import "fmt"

func main() {
	//res := sum(10, 90)
	//fmt.Println(res)

	//result := validateUser("Fer", true, 450.0)
	//fmt.Println(result)

	res := checkAccess(20, false)
	fmt.Println(res)
}

func sum(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func validateUser(user string, active bool, balance float64) string {
	if !active {
		return "user not active"
	} else if balance < 100 {
		return "user balance less than 100 (insufficient balance)"
	}
	return "user is valid"
}

func checkAccess(age int, isMember bool) string {
	if age < 18 {
		return "access denied: underage"
	} else if age > 18 && !isMember {
		return "access denied: not a member"
	}
	return "access granted"
}
