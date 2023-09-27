package main

import (
	"fmt"
	"unsafe"
)

type structure1 struct {
	a bool   // 1
	b bool   // 1
	c string // 16
	d bool   // 1
}

type structure2 struct {
	a bool   // 1
	b bool   // 1
	c bool   // 1
	d string // 16
}

func main() {

	var sample1 = structure1{}
	addr1A := &sample1.a
	addr1B := &sample1.b
	addr1C := &sample1.c
	addr1D := &sample1.d
	fmt.Printf("addr1A = %p\n", addr1A) // 0x1400005e020
	fmt.Printf("addr1B = %p\n", addr1B) // 0x1400005e021
	fmt.Printf("addr1C = %p\n", addr1C) // 0x1400005e028
	fmt.Printf("addr1D = %p\n", addr1D) // 0x1400005e038
	fmt.Printf("size of structure1 = %v\n", unsafe.Sizeof(sample1))

	fmt.Println()
	var sample2 = structure2{}
	addr2A := &sample2.a
	addr2B := &sample2.b
	addr2C := &sample2.c
	addr2D := &sample2.d
	fmt.Printf("addr2A = %p\n", addr2A) // 0x140000b8000
	fmt.Printf("addr2B = %p\n", addr2B) // 0x140000b8001
	fmt.Printf("addr2C = %p\n", addr2C) // 0x140000b8008
	fmt.Printf("addr2D = %p\n", addr2D) // 0x140000b8002
	fmt.Printf("size of structure2 = %v\n", unsafe.Sizeof(sample2))
}
