package main

import "fmt"

func fibonacci() func() int {
	curr := 0
	prev := 0
	temp := 0
	return func() int {
		if curr == 0 && prev == 0 {
			prev++
			return 0
		} else {
			temp = curr
			curr += prev
			prev = temp
			return curr
		}
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	f2 := fibonacci()

	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())

	fmt.Println(f())
	fmt.Println(f())
}
