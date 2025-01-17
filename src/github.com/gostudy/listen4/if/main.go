package main

import "fmt"

func testF1() {
	//var num int
	//num = 10
	num := 10
	if num %2 ==0 {    	//if写法1，变量生命周期在if块内
		fmt.Printf("num:%d is even\n",num)
	} else {
		fmt.Printf("num:^d is odd\n",num)
	}
}

func testF2()  {
	num := 10
	if num >5 && num <10 {
		fmt.Printf("num:%d is >5 and <10\n",num)
	} else if num >=10 && num <20 {
		fmt.Printf("num:%d is >10 and <20\n",num)
	} else if num > 20 && num <30 {
		fmt.Printf("num:%d is >20 and <30\n",num)
	}else {
		fmt.Printf("num:%d is >30\n",num)
	}
}


func testF3() {    //if写法2
	//var num int
	//num = 10

	if num := 11; num %2 ==0 {
		fmt.Printf("num:%d is even\n",num)
	} else {
		fmt.Printf("num:%d is odd\n",num)
	}
}

func getNum() int  {
	return  10
}


func testF4() {    //if写法2
	//var num int
	//num = 10

	if num := getNum(); num %2 ==0 {    //if内调用函数
		fmt.Printf("num:%d is even\n",num)
	} else {
		fmt.Printf("num:%d is odd\n",num)
	}
}




func main()  {
	testF1()
	testF2()
	testF3()
	testF4()

}