package main

import "fmt"

//Дано трехзначное число. Найдите сумму его цифр.
func main131() {
	var a, r int
	fmt.Scan(&a)
	r = a/100 + a/10%10 + a%10
	fmt.Print(r)
}

//Дано трехзначное число. Переверните его, а затем выведите.
func main132() {
	var a, r int
	fmt.Scan(&a)
	r = a/100 + a/10%10*10 + a%10*100
	fmt.Print(r)
}

/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток. Например, если
k=13257=3*3600+40*60+57,
то h=3 и m=40.
*/
func main133() {
	var mm, hh, ss int
	fmt.Scan(&ss)
	hh = ss / 60 / 60
	mm = (ss - hh*60*60) / 60
	fmt.Printf("It is %d hours %d minutes.", hh, mm)
}

/*
Заданы три числа
a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить,
 является ли треугольник прямоугольным. Если является,
  вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/
func main134() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a*a+b*b == c*c {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}

}

func main135() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	r := "Существует"
	if a+b <= c || a+c <= b || b+c <= a {
		r = "Не существует"
	}
	fmt.Printf("%s", r)
}

//Даны два числа. Найти их среднее арифметическое.
func main136() {
	var (
		a, b int
		//c float32
	)
	fmt.Scan(&a)
	fmt.Scan(&b)
	c := float32(a+b) / 2
	fmt.Printf("%v", c)
}

//По данным числам, определите количество чисел, которые равны нулю.
func main13() {
	var a, b int
	c := 0
	for fmt.Scanln(&a); a >= 0; fmt.Scanln(&b) {
		if b == 0 {
			c++
		}
		a--
	}
	fmt.Printf("%v", c)

}
