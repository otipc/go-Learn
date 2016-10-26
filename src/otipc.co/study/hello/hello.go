package main

import (
	"fmt"

	"otipc.co/cal"
	"otipc.co/cal/multi"
	"otipc.co/study/hellolib"
)

func main() {

	fmt.Print("你好，世界！\n")
	fmt.Printf("2 和 3中最大的是 %d！\n", hellolib.Max(2, 3))
	fmt.Printf("2 和 3中最小的是 %d！\n", hellolib.Min(2, 3))

	result := cal.Add(2, 3)
	fmt.Printf("%d\n", result)

	result = cal.Subtrack(30, 3)
	fmt.Printf("%d\n", result)

	result = multi.Multi(2, 4)
	fmt.Printf("%d\n", result)

}
