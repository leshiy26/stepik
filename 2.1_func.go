package main

import "fmt"

/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */
func main1111() {
	a, b := sumInt(1, 2, 3, 4)
	fmt.Printf("%v, %v", a, b)
}

func sumInt(a ...int) (n, sum int) {
	s := 0
	for _, elem := range a {
		s += elem
	}

	return len(a), s
}
