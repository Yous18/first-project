package main

import "fmt"

func main() {
	var UserNumber int
	fmt.Println("Vitya, print your number: ")
	fmt.Scan(&UserNumber)
	

}
func getNearestFibo (UserNumber int) {
	var n1 int = 0
	var n2 int = 1
	var n3 int = n2
	fmt.Printf("FiboList: %v", n1)
	for n := range UserNumber {
		n3 = n2
		n2 = n1 + n2
		if n2 >= x {
			break
		}
		n1 = n3
		fmt.Printf(", %v", n3)
	}

}
