// < 2.2.9 오류 >
package study

import (
	"errors"
	"fmt"
)

// 코드 2.28 : 간단한 DivisionByZero 오류
var DivisionByZero = errors.New("division by zero")

func Divide(number, d float32) (float32, error) {
	if d == 0.0 {
		return 0, DivisionByZero
	}
	return number / d, nil
}
func Code228() {
	n1, e1 := Divide(1, 1)
	fmt.Println(n1)
	if e1 != nil {
		fmt.Println(e1.Error())
	}
	n2, e2 := Divide(1, 0)
	fmt.Println(n2)
	if e2 != nil {
		fmt.Println(e2.Error())
	}
}

// 코드 2.29 : 패닉의 동작(단일 함수 호출)
var SampleError = errors.New("This is a test error")

// 다른 언어의 try-catch문과 비슷하게 go에서는 defer, panic, recover를 사용하여 예외처리를 할 수 있다.
func testRecover() {
	defer func() {
		if recover() != nil {
			fmt.Println("got an error!")
		} else {
			fmt.Println("no error")
		}
	}()
	// 패닉을 주석 처리하면 "Hello!"가 출력되고, 주석을 해제하면 "got an error!"가 출력된다.
	//panic(SampleError)
	fmt.Println("Hello!")
}
func Code229() {
	testRecover()
}

// 코드 2.30 : 패닉의 동작(여러 함수 호출)
func testPanic() {
	panic(SampleError)
	// 출력 메세지 수행하지 못하고 바로 반환됨.
	fmt.Println("Hello from testPanic!")
}
func testRecover2() {
	defer func() {
		if recover() != nil {
			fmt.Println("got an error!")
		} else {
			fmt.Println("no error")
		}
	}()
	testPanic()
	fmt.Println("Hello from testRecover2!")
}
func Code230() {
	testRecover2()
}
