package main

import (
	"fmt"
	"unsafe"
)

type Test struct {
	age  int
	name string
}

/*
//Arbitrytype is used only for documentation purposes and is not actually part of the unsafe package, which represents the type of any go expression.
type ArbitraryType int
//Pointer of any type, similar to * void of C
type Pointer *ArbitraryType
//Determine the exact size of the structure in memory
func Sizeof(x ArbitraryType) uintptr
//Returns the offset of a field in a structure
func Offsetof(x ArbitraryType) uintptr
//Returns the value of a field in a structure (the reason for byte alignment)
func Alignof(x ArbitraryType) uintptr
*/
func main() {

	i := 10
	ip := &i
	fp := (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(int(*fp))
	fmt.Println(i)

	test := Test{}
	test.age = 56
	test.name = "David"

	pointerTest := unsafe.Pointer(&test)
	Age := (*int)(unsafe.Pointer(uintptr(pointerTest) + unsafe.Offsetof(test.age)))
	Name := (*string)(unsafe.Pointer(uintptr(pointerTest) + unsafe.Offsetof(test.name)))
	*Age = 9654
	*Name = "hello wolrd"
	fmt.Println((test))

	b := Test{}
	b.age = 5822
	b.name = "hesjdasdja"
	bPointer := unsafe.Pointer(&b)
	bName := (*int)(unsafe.Pointer(uintptr(bPointer)))
	*bName = 55
	fmt.Println(b.age)
	//Get and modify according to sizeof function
	arr := []int{1, 2, 5, 2, 5, 5, 222, 1, 2, 0, 22}
	pointer := &arr[0]
	fmt.Println(pointer)
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(arr[0])
	for i := 0; i < len(arr)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(arr[0])
	}
}
