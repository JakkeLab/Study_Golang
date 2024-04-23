package main

import "fmt"

func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	//defer는 스택구조이므로, 맨처음에 fmt.Println(0) 부터 쌓여서 맨위에 fmt.Println(9)가 놓인다.
	//따라서 9 8 7 ... 1 0 으로 출력되게 된다.
}
