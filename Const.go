package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday   = Monday + 1
	Wednesday = iota + 3
	Thursday
	Friday
	Saturday
	Add
)

func main() {
	fmt.Println(Sunday)                  // вывод 0
	fmt.Println("Tuesday:", Tuesday)     // вывод 6
	fmt.Println("Wednesday:", Wednesday) // вывод 6
	fmt.Println("Thursday:", Thursday)   // вывод 6
	fmt.Println("Tuesday:", Tuesday)     // вывод 6
	fmt.Println(Wednesday)               // вывод 6
	fmt.Println("Saturday:", Saturday)   // вывод 6
}
