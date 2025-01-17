package main

import "fmt"

func describe(a interface{})  {
	fmt.Printf("%T %v\n", a, a)
}

func testInterface()  {
	var a interface{}   //空接口可以存储任何类型
	var b int = 100
	a = b

	fmt.Printf("%T %v\n", a, b)

	var c string = "golang"
	a = c
	fmt.Printf("%T %v\n", a,a)

	var d map[string]int = make(map[string]int, 100)
	d["abc"] = 1111
	d["ddd"]= 122
	a = d
	fmt.Printf("%T %v\n", a,a)
}

type Student struct {
	Name string
	Sex  int
}

func main()  {
	a := 65
	describe(a)
	var stu Student = Student{
		Name:"user01",
		Sex: 1,
	}
	describe(stu)
}
