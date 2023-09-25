package main

import "fmt"

func main() {
	//a := 2
	i := 4
	b := 3
	if a := 5; a < b {
		fmt.Println("a < b")
	} else if a > b {
		fmt.Println("Ветка Else if")
	} else {
		fmt.Println("Переменные идентичны")
	}

	switch i {
	case 0:
		fmt.Println("0")
	case 4:
		fmt.Println("4")
		fmt.Println("44")
		fallthrough
	default:
		fmt.Println("def")
	}

	var c float32 = 9
	switch {
	case 1 <= c && c <= 9:
		fmt.Print("от 1 до 9 ")
		c--
		fallthrough
	case c == 10.2:
		fmt.Print("пройден ")
		fmt.Print(c)
	case 100 <= c && c <= 250:
		fmt.Print("от 100 до 250 ")
		fmt.Print(c)
	case 1000 <= c && c <= 6000:
		c += 999
		fmt.Print("от 1000 до 6000 ")
		fallthrough
	default:
		fmt.Print("default ")
	}

	/*var (
		d int
	)
	fmt.Scanln(&d)
	switch {
	case d > 0:
		println("Число положительное")
	case d < 0:
		println("Число отрицательное")
	default:
		println("Ноль")

	}
	*/
	var (
		e int
	)
	fmt.Scanln(&e)
	if e > 0 {
		print("Число положительное")
	} else if e < 0 {
		println("Число отрицательное")
	} else {
		println("Ноль")
	}

}


