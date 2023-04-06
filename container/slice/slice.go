package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] // [2,3,4,5]
	s2 := s1[3:5]  // [5,6]
	fmt.Println(s2)

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5, arr)
	fmt.Printf("addr=%p\n", s3)
	fmt.Printf("addr=%p\n", s2)

	//a1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s := a1[2:3:7] //[low:high:max] len=high-low,cap=max-low
	//fmt.Println(s)
	//fmt.Printf("len(s) = %d, cap(s)=%d, addr=%p\n", len(s), cap(s), s)
	//fmt.Printf("addr=%p\n", a1)

}
