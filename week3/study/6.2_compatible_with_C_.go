// < 6.2 C 코드와 상호 호환하기 >
package study

// 코드 6.3 : Go에서 C "Hello World" 함수 호출

// #include <stdio.h>
// void printHelloWorld() {
//     printf("Hello, World from C!\n");
// }
import "C"

func Code63() {
	C.printHelloWorld() // C의 printHelloWorld 함수 호출
}

// // 코드 6.4 : Go에서 C로 케릭터 포인터 전달하기

// // #include <stdio.h>
// // #include <stdlib.h>
// // void printString(char* str) {
// //	 printf("%s\n", str);
// // }
// import "C"
// import "unsafe"

// func Code64(){
// 	a := C.CString("This is from Golang")
// 	C.printString(a)
// 	C.free(unsafe.Pointer(a))
// }
