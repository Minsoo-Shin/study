package main

import (
	"fmt"
)

func main() {
	var arrayInt = [5]int{1, 2, 3, 4, 5}

	addr0 := &arrayInt[0]
	addr1 := &arrayInt[1]
	addr2 := &arrayInt[2]
	// 데이터 크기가 8byte 떄문에 메모리 주소가 8의 배수
	fmt.Printf("arrayInt[0] = %p\n", addr0) // 0x140000b2000
	fmt.Printf("arrayInt[1] = %p\n", addr1) // 0x140000b2008
	fmt.Printf("arrayInt[2] = %p\n", addr2) // 0x140000b2010

	fmt.Println()

	var arrayStr = [3]string{"가", "나", "다"}

	addrStr0 := &arrayStr[0]
	addrStr1 := &arrayStr[1]
	addrStr2 := &arrayStr[2]
	// 데이터 크기가 16byte 떄문에 메모리 주소가 16의 배수
	fmt.Printf("arrayStr[0] = %p\n", addrStr0) // 0x1400009a0c0
	fmt.Printf("arrayStr[1] = %p\n", addrStr1) // 0x1400009a0d0
	fmt.Printf("arrayStr[2] = %p\n", addrStr2) // 0x1400009a0e0
	// string
	/* string의 크기는 StringHeader 구조체의 크기와 문자열의 길이에 따라 결정됩니다.
	type StringHeader struct {
			Data uintptr
			Len  int
		}
	*/

}
