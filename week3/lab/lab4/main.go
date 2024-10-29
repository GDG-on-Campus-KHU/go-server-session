package main

/*
#cgo CFLAGS: -O0
#include <stdlib.h>
#include "string_concat.c"

char* string_concat(const char* str1, const char* str2);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	str1 := C.CString("Hello, ")
	str2 := C.CString("World!")

	result := C.string_concat(str1, str2)

	goResult := C.GoString(result)

	fmt.Println("Concatenated string:", goResult)

	defer C.free(unsafe.Pointer(str1)) 
	defer C.free(unsafe.Pointer(str2))
	C.free(unsafe.Pointer(result))
}