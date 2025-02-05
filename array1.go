package main

import "fmt"

func main() {
	//1
	s := []int{73,98,86,61,96,10}
	fmt.Println(s)
	//если привысить  слайса, то он увеличится в 2 раза

	//2 Удаление элемента по индексу
	s1 := []int{73, 98, 86, 61, 96, 10}
    fmt.Println(s)

    index := 2
    
    s1 = append(s1[:index], s1[index+1:]...)

    fmt.Println("Slice after deletion:", s1)

	//3 Как создать слайс с нулевой длиной, но вместимостью 10
	s2 := make([]int, 0, 10)
	fmt.Println(s2)
	fmt.Println("длина:",len(s2), "вместимость:",cap(s2))

	//4 Как скопировать элементы одного слайса в другой?
	s3 := s2[0:9:10]
	fmt.Println(s3)

	//5 Как получить последний элемент слайса, не зная его длину заранее?
	s4 := []int{73, 98, 86, 61, 96, 10, 15, 20, 25}
	len1 := len(s4)-1
	fmt.Println(s4[len1])

	




}