package main

import (
	"os"
	//"fmt"
)

func main() {
	var buf [16]byte
	os.Stdin.Read(buf[:])
	//fmt.Println(string(buf[:]))

	os.Stdout.WriteString(string(buf[:]))
	//fmt.Scanf
}