package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//슬라이스는 Call by Reference 방식처럼 슬라이스로 뽑아낸 배열속의 요소를 바꾸면 원본에서도 해당 요소가 변경된다.
}
