package main

import "fmt"

func slice() {
	//func1()
	func2()
}

func func1() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var myslice []int = array[:3]

	fmt.Println("Element of Array:")
	for _, v := range array {
		fmt.Println(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")
	for _, v := range myslice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// slicecopy
func func2() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Print(slice2)
	fmt.Print(slice1)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Print(slice2)
	fmt.Print(slice1)
}
