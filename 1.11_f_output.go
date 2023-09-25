package main

import "fmt"

func main111() {
	var (
		a   rune    = 'Ы'
		b   int     = 33
		c   float32 = 432.2222
		a1          = 3.14
		a11 string  = "123"
		a12 string  = "1234"
	)

	fmt.Printf("%q %T \n", a, b)
	fmt.Printf("%v %#v %v \n", a, a, a) //Также можно применять универсальный спецификатор %v, который для типа boolean аналогичен %t, для целочисленных типов - %d, для чисел с плавающей точкой - %g, для строк - %s.
	fmt.Printf("%10.3f \n", c)

	fmt.Printf("|%8f|\n", a1)
	fmt.Printf("|%8.1f|\n", a1)
	fmt.Printf("|%8.3f|\n", a1)
	fmt.Printf("|%-8.3f|\n", a1)

	fmt.Printf("%q \n%s \n", a11, a12)
}

func main11() {
	var a float64
	fmt.Scan(&a)
	if a <= 0 {
		fmt.Printf("число %2.2f не подходит", a)
	} else if a > 10000 {
		fmt.Printf("%e", a)
	} else {
		a *= a
		fmt.Printf("%.4f", a)
	}
}

/*

вывод:
|3.140000|
|     3.1|
|   3.140|
|3.140   |
*/

/*
	var a float64 = 100.123456
	fmt.Printf("это число %f типа %T \n", a, a)
	// вывод: это число 100.123456 типа float64

	var a1 byte = 's'
	var a2 int = 1234
	fmt.Printf("%q %b\n", a1, a2)
	// вывод: 's' 10011010010


	// использование \n позволяет сделать перенос строки
	var a1 string = "123"
	var a2 string = "1234"
	fmt.Printf("%q \n%s\n", a1, a2)
	// вывод:
	// "123"
	// 1234

	/*
	%t: для вывода значений типа boolean (true или false)
%b: для вывода целых чисел в двоичной системе
%c: для вывода символов, представленных числовым кодом
%d: для вывода целых чисел в десятичной системе
%o: для вывода целых чисел в восьмеричной системе
%q: для вывода символов в одинарных кавычках
%x: для вывода целых чисел в шестнадцатеричной системе, буквенные символы числа имеют нижний регистр a-f
%X: для вывода целых чисел в шестнадцатеричной системе, буквенные символы числа имеют верхний регистр A-F
%U: для вывода символов в формате кодов Unicode, например, U+1234
%e: для вывода чисел с плавающей точкой в экспоненциальном представлении, например, -1.234456e+78
%E: тоже самое что %e но в верхнем регистре, например, -1.234456E+78
%f: для вывода чисел с плавающей точкой, например, 123.456
%F: то же самое, что и %f
%g   %e для огромных экспонент, %f в противном случае
%G    %E для огромных экспонент, %F в противном случае
%s: для вывода строки
%p: для вывода значения указателя - адреса в шестнадцатеричном представлении (указатели мы пройдем на следующих

	%f: точность и ширина значения по умолчанию
%9f: ширина - 9 символов и точность по умолчанию
(число с плавающей точкой будет занимать как минимум 9 позиций. Если ширина больше, чем требуется значению, то заполняется пробелами.)
%.2f: ширина по умолчанию и точность - 2 символ
%9.2f: ширина - 9 и точность - 2
%9.f: ширина - 9 и точность - 0


Упарвляющие символы

\\
\"
\f - подача страницы!
\v - вертикальный таб!
\r
\b  - возврат (backspace U+0008)
\t - tab
\n



*/
