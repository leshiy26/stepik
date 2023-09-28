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
	var a int
	c := 0
	b := 1

	fmt.Scanln(&a)

	for a > 0 {
		fmt.Scanln(&b)
		if b == 0 {
			c++
		}
		a--
	}
	fmt.Printf("%v", c)
}

//Найдите количество минимальных элементов в последовательности.
func main137() {
	var a, b int
	c := 0
	min := 0

	fmt.Scanln(&a)

	if a > 0 {
		fmt.Scanln(&min)
		a--
		c++
	}

	for a > 0 {
		fmt.Scanln(&b)
		if b < min {
			c = 1
			min = b
		} else if b == min {
			c++
		}
		a--
	}
	fmt.Printf("%v", c)
}

/*
//Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр,
на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации.
 Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
 Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
По данному числу определите его цифровой корень.
*/

func main138() {
	var a int
	fmt.Scanln(&a)
	b := 0

	for i := a; ; i = i / 10 {

		b += i % 10
		if i <= 1 {
			break
		}
	}

	c := b%10 + b/10
	fmt.Printf("%v", c)

}

/*
Найдите самое большее число на отрезке от a до b, кратное 7 .
*/

func main139() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	c := 0
	for i := b; i >= a; i-- {
		if i%7 == 0 {
			b = i
			c = 1
			break
		}
	}
	if c == 0 {
		fmt.Printf("NO")
	} else {
		fmt.Printf("%v", b)
	}

}

/*
По данному числу n закончите фразу "На лугу пасется..."
 одним из возможных продолжений: "n коров", "n корова", "n коровы",
  правильно склоняя слово "корова".
*/

func main1310() {
	var (
		a, b int
		c    string
	)

	fmt.Scanln(&a)
	b = a % 10

	if b == 0 || b >= 5 || (a >= 10 && a <= 14) {
		c = "korov"
	} else if b == 1 {
		c = "korova"
	} else {
		c = "korovy"
	}

	fmt.Printf("%v %s", a, c)

}

/*
Дано натуральное число N. Выведите его представление в двоичном виде.
*/
func main1311() {
	var N int
	fmt.Scanln(&N)
	fmt.Printf("%b", N)
}

/*
Из натурального числа удалить заданную цифру.
*/
func main1312() {
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
