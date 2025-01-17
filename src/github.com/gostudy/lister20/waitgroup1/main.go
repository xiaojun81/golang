package main

import (
	"time"
	"fmt"
)

func process(i int, ch chan bool)  {
	fmt.Println("started Goroutime", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	ch <- true
}

func main()  {
	no := 3
	exitChan := make(chan bool, no)
	for i := 0; i < no; i++ {
		go process(i, exitChan)
	}
	for i :=0; i < no; i++ {
		<-exitChan
	}
	fmt.Printf("All go routines finished executing")
}