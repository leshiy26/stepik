package main

import "fmt"

func main2() {
	var a, x, y, z int
	fmt.Scanln(&a)
	x = a % 10
	y = a / 10 % 10
	z = a / 100 % 10
	fmt.Printf("x=%d,y=%d,z=%d", x, y, z)
	if x != y && x != z && z != y {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// вывести первую цифру
func main1() {
	var a int
	fmt.Scan(&a)
	switch {
	case a == 10000:
		fmt.Println("1")
	case a < 10000 && a > 1000:
		fmt.Println(a / 1000)
	case a < 1000 && a > 100:
		fmt.Println(a / 100)
	case a < 100 && a > 10:
		fmt.Println(a / 10)
	default:
		fmt.Println(a)
	}

}

// счастливый билет
func main3() {
	var a, res1, res2, b1, b2 int
	fmt.Scan(&a)
	b1 = a / 1000
	b2 = a % 1000
	res1 = b1/100 + b1%10 + b1%100/10
	res2 = b2/100 + b2%10 + b2%100/10
	if res1 == res2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

// определить високосный год
func main() {
	var y int
	fmt.Scan(&y)
	if y%400 == 0 || y%4 == 0 && y%100 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
