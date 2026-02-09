package main

import "fmt"

func main() {
	type Student struct {
		Name string
		Age  int
	}
	var std Student = Student{"Prakash Joshi", 30}
	var arr [3]int = [3]int{10, 22, 31}
	fmt.Println(arr)
	fmt.Println(std.Age)
	fmt.Println("Student Name", std.Name)
	var u interface{} = "Prakash Joshi"
	fmt.Println(u)
	u = 30
	fmt.Println(u)
	u = "Teja"
	fmt.Println(u)
	var AddTwoNums = func(a, b int) int {
		return a + b
	}
	fmt.Println(AddTwoNums(5, 7))
}
