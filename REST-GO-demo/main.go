package main

import (
	"fmt"
)

func main() {
	fmt.Println("火币")
	
	// 测试 - second commit

	// test on Code with sourcetree
	// test on dir Doc


	/*
	 * edit with golang,
	 * commit use source tree
	 * dir in Code
	 */

	 fmt.Printf("测试\n")

	 // 创建数组
	 var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}

	 // 基于数组创建数组切片
	 var mySlice []int = myArray[:4]

	 fmt.Println("Elements of myArray: ")
	 for _,v := range myArray{
	 	fmt.Print(v, " ")
	 }

	 fmt.Println("\nElements of mySlice: ")
	 for _,v := range mySlice{
	 	fmt.Print(v, " ")
	 }

	 fmt.Println(cap(mySlice))
}

