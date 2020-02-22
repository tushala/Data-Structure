package main

import "fmt"

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/*golang 中slice 即为动态数组*/
func main() {
	//var numbers []int
	//// 添加元素
	//PrintSlice(numbers)
	//numbers = append(numbers, 1)
	//PrintSlice(numbers)
	//numbers = append(numbers, 2, 3, 4)
	//PrintSlice(numbers)
	//fmt.Println(numbers[3])

	numbers1 := make([]int, 2)
	PrintSlice(numbers1)
	// 扩容
	numbers1 = append(numbers1, 2, 3, 4)
	PrintSlice(numbers1)
	// 扩容机制
	a := []int{}
	for i := 0; i < 16; i++ {
		a = append(a, i)
		fmt.Print(cap(a), " ")
	}
	// 1 2 4 4 8 8 8 8 16 16 16 16 16 16 16 16
}
