package main

import "fmt"

/*
Из натурального числа удалить заданную цифру.
*/
func main() {
	var N, t, u int
	fmt.Scanln(&N)
	fmt.Scanln(&u)
	a := make([]int, 0, 100)
	//	fmt.Printf("Тип: %T, Значени: %v", a, a)
	for N > 0 {
		t = N % 10
		if t != u {
			a = append(a, t)
		}
		N = N / 10
	}
	for i := len(a); i >= 1; i-- {
		fmt.Printf("%v", a[i-1])
	}

}
