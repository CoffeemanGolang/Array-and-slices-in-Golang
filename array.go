package main

import "fmt"

func main() {
	//array
	arr:= [5]int{73,98,86,61,96}
	fmt.Println(arr)

	num := arr[2]
	fmt.Println(num)

	arr[2] = 100
	fmt.Println(arr)

	//slice
	s := []int{73,98,86,61,96}
	s[4] = 505
	//массив не может создаваться ячейку по индексу, если она не существует, а слайс может
	s = append(s, 100)
	fmt.Println(s)
	//длина слайса
	fmt.Println("длина:",len(s))
	//вместимость слайса
	fmt.Println("вместимость:",cap(s))

	//создание слайс с определенной длиной и вместимостью
	s1 := make([]int, 5, 100)
	fmt.Println(s1)
	fmt.Println("длина:",len(s1), "вместимость:",cap(s1))

	//получить подслайс из слайса
	s2 := s[0:3:10]//начало : конец : вместимость
	fmt.Println(s2)
	fmt.Println("длина:",len(s2), "вместимость:",cap(s2))

}