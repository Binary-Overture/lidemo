package main

import (
	"fmt"
)

// slice根据数组array创建，与array共享数据存储空间
//
//	func main() {
//		var array [10]int
//		var slice = array[9:10]
//
//		fmt.Println("len of slice", len(slice))
//		fmt.Println("capacity of slice", cap(slice))
//		fmt.Println(&slice[0] == &array[6])
//		fmt.Println(&slice[0])
//
// }

func main() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	for i := 0; i < len(order); i++ {
		order[i] = uint16(i)
	}
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]
	fmt.Println(pollorder, lockorder)
	fmt.Println("len(pollorder) == ", len(pollorder))
	fmt.Println("cap(pollorder) == ", cap(pollorder))
	fmt.Println("len(lockorder) == ", len(lockorder))
	fmt.Println("cap(lockorder) == ", cap(lockorder))

}
