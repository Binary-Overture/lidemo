package main

import "fmt"

func main() {
	str1 := "E`}J]OrQF[V8zV:hzpV}fVF[t"
	//var shellcode = ""
	for _, v := range str1 {
		fmt.Printf(string(v ^ 9))
	}
}
