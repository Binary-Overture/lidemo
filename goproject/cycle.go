package main

import "fmt"

func main() {
	var c []func()
	for _, v := range []int{1, 2, 3} {
		fmt.Println("v:", v)
		c = append(c, func() {
			fmt.Println(v)
		})
	}

	for _, vv := range c {
		vv()
	}

	return
}
