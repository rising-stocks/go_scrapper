package main

import "fmt"

// Variable expression
func canIDrink(age int) bool {

	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
