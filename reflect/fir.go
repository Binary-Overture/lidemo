package main

import (
	"strconv"
)

//func main() {
//	x := Spring(true)
//	fmt.Println(x)
//	return
//}

func Spring(x interface{}) string {
	type Stringer interface {
		String() string
	}
	switch x := x.(type) {
	case Stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "??"
	}
}
