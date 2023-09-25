package main

import "fmt"

//Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.
func main1() {
	var a, b, sum int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	for ; a <= b; a++ {
		sum += a
	}
	fmt.Println(sum)
}

//Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.
func main2() {
	var a, b, result int
	for fmt.Scan(&a); a > 0; a--{
		fmt.Scan(&b)
		if b % 8 == 0 && b < 100 && b >= 10 {
			result += b 
		}
	}
	fmt.Println(result)
}

func main4() {
	var a, a_max, q_max int
	for fmt.Scan(&a); a != 0; fmt.Scan(&a){
		if a > a_max{
			a_max = a
			q_max = 1
		}else if a == a_max {
			q_max += 1
		}
	}
	fmt.Print(q_max)
}

func main5() {
	var  n, c, d int
	fmt.Scanln(&n)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	for i:=1; i<=n; i++ {
		if i % c == 0 && i % d != 0 {
			fmt.Print(i, " ")
			break
		}
	}
}


/*
Напишите программу, которая считывает целые числа с консоли по одному числу в строке.

Для каждого введённого числа проверить:

если число меньше 10, то пропускаем это число;
если число больше 100, то прекращаем считывать числа;
в остальных случаях вывести это число обратно на консоль в отдельной строке.
*/
func main6() {
	var a int

	for fmt.Scanln(&a);a>0;fmt.Scanln(&a) {
        
        if a < 10  {
            continue
        } else if a > 100 {
            break
        } else {
            fmt.Println(a)            
		}
    }
}


//Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
func main7() {
	var  x, p, y, n int
	fmt.Scanln(&x)
	fmt.Scanln(&p)
	fmt.Scanln(&y)
	for x < y {
		n++ 
		x += x * p / 100
	} 
	fmt.Println(n)
}

//Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
func main() {
	var  x, y, b, start int
	fmt.Scan(&x)
	fmt.Scan(&y)
	for d := 10000; d > 0; d /= 10{
		b = x / d
		x %= d
		if b == 0 && start == 0 {
			continue
		}else {
			start = 1
		}
		for c := y; c > 0; c /= 10{
			if c % 10 == b {
				fmt.Printf("%d ",b)
			}
		}
	}
}