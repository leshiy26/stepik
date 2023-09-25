package main
import "fmt"

func main1(){
var a [3]int = [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [...]int{1, 2, 3} // ... вместо указания длинны массива
d := [3]int{1: 12}
e := [5]int{1, 2, 3, 4, 5}
fmt.Println(a, b, c,d,e) // [1 2 3 4 5]

//итерация по всем элементам (range получает индекс и КОПИЮ элемента массива)
f := [5]int{1, 2, 3, 4, 5}
for idx, elem := range f {
    fmt.Printf("Элемент с индексом %d: %d\n", idx, elem)
}

}
func main2()  {
	var indA,indB,tmp,i uint8
	workArray := [10]uint8{}
	for idx := range workArray{
		fmt.Scan(&i)
		workArray[idx] = i
	}	
	for j:=0; j<3; j++{
		fmt.Scan(&indA, &indB)
		tmp = workArray[indA]
		workArray[indA] = workArray[indB]
		workArray[indB] = tmp
	}
	for _, num := range workArray{
		fmt.Print(num, " ")
	}
}

// Напишите программу, принимающая на вход число N(N≥4), 
// а затем N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу)
// элемент данного среза.
func main3() {
	var (N int
		 m int)
	fmt.Scan(&N)
	s := make([]int, N, N)
	for i:=0; i<N; i++{

		fmt.Scan(&m)
		s[i] = m
	}
	fmt.Print(s[3], ' ')
}

//Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
func main4()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
    // здесь ваш код
	max:=array[0]
	for _, val :=  range(array){
		if val > max  {
			max = val
		}
	}
	fmt.Print(max)



    // ...
}

//Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. 
//Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
func main5() {
	var n int
	fmt.Scan(&n)
	sl:= make([]int, n, n)
	for i:=0; i < n; i++{
		fmt.Scan(&sl[i])
	}

	for i:=0; i<=(len(sl)-1)/2; i+=1{
		fmt.Print(sl[i*2], " ")
	}

}

//Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.
func main() {
	var n int
	fmt.Scan(&n)
	sl:= make([]int, n, n)
	for i:=0; i < n; i++{
		fmt.Scan(&sl[i])
	}

	q:=0
	for _, val := range(sl) {
		if val > 0 {
			q++
		}
	}
	fmt.Print(q)

}


//Слайс (срез, slice) - это последовательность элементов одного типа переменной длины
//a := make([]int, 10, 10)initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"} // базовый массив
/*

//СРЕЗ 

users1 := initialUsers[2:6] // с 3-го по 6-й
users2 := initialUsers[:4] // с 1-го по 4-й
users3 := initialUsers[3:] // с 4-го до конца

fmt.Println(users1) // [Kate Sam Tom Paul]
fmt.Println(users2) // [Bob Alice Kate Sam]
fmt.Println(users3) // [Sam Tom Paul Mike Robert]

// добавить в срез
a := []int{1, 2, 3}
a = append(a, 4, 5)
fmt.Println(a) // [1 2 3 4 5]



*/

