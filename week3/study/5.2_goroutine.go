// < 5.2 고루틴 >
package study

import "fmt"

// 코드 5.1: 숫자를 제곱하는 간단한 싱글 스레드 프로그램
func squareIt(x int) {
	fmt.Println(x * x)
}
func Code51() {
	squareIt(2)
}

// 추가 코드 : 고루틴을 사용하여 숫자를 제곱하는 프로그램
func squareItAsync(x int, done chan bool) {
	fmt.Println(x * x)
	done <- true
}

func GoRoutineMain() {
	done := make(chan bool)
	go squareItAsync(2, done)
	<-done
}
