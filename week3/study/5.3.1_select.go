// < 5.3.1 select 문 >
package study

import (
	"fmt"
	"time"
)

// 코드 5.3 : 별도의 고루틴에서 함수의 생성과 호출을 동시에 진행하기
func Code53() {
	inputChan := make(chan int, 10)  // 버퍼 크기가 10인 inputChan 채널을 생성한다.
	finishChan := make(chan int)     // 종료 신호를 받을 finishChan 채널을 생성한다.
	outputChan := make(chan int, 10) // 버퍼 크기가 10인 outputChan 채널을 생성한다.

	// 고루틴을 사용하여 입력을 처리하고 종료 신호를 받는 함수를 실행한다.
	go func(inputChan, finishChan chan int) {
		for {
			select {
			// inputChan에서 데이터를 받으면 제곱하여 outputChan에 보낸다.
			case x := <-inputChan:
				outputChan <- x * x
			// finishChan에서 종료 신호를 받으면 함수를 종료한다.
			case _ = <-finishChan:
				return
			}
		}
	}(inputChan, finishChan)

	// 0부터 9까지의 숫자를 inputChan에 보낸다.
	for i := 0; i < 10; i++ {
		inputChan <- i
	}

	// outputChan에서 제곱된 값을 10개 받는다.
	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChan)
	}

	// 고루틴을 종료시키기 위해 finishChan에 값을 보낸다.
	finishChan <- 1
}

// 코드 5.4 : 클로저에서 Go의 변수 캡처 (변수 캡처 문제 발생)
func Code54() {
	// 고루틴을 실행하여 i 값을 출력한다.
	for i := 0; i < 10; i++ {
		go func() {
			// 1 밀리초 동안 대기한다.
			time.Sleep(1 * time.Millisecond)
			// i 값을 출력한다.
			fmt.Println(i)
		}()
	}
	// 메인 함수가 종료되지 않도록 100 밀리초 동안 대기한다.
	time.Sleep(100 * time.Millisecond)
}

/*
// 변수 캡처 문제 해결 방법
// 각 고루틴이 실행될 때 i는 for 루프가 끝난 후의 값을 참조할 가능성이 크다. 이 문제를 해결하려면 i 값을 고루틴 안으로 안전하게 전달해야 한다.
func Code54() {
	for i := 0; i < 10; i++ {
		// 고루틴에 i 값을 인자로 전달하여 캡처한다.
		go func(i int) {
			time.Sleep(1 * time.Millisecond)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
}
*/

// 코드 5.5 : 클로저에 전달할 변수의 복사본 만들기
func Code55() {
	for i := 0; i < 10; i++ {
		// 고루틴에 i 값을 인자로 전달하여 캡처한다.
		go func(nonCapturedI int) {
			// 1 밀리초 동안 대기한다.
			time.Sleep(1 * time.Millisecond)
			// 캡처된 nonCapturedI 값을 출력한다.
			fmt.Println(nonCapturedI)
		}(i) // i 값을 고루틴의 인자로 전달한다.
	}
	// 고루틴이 모두 실행될 시간을 주기 위해 100 밀리초 동안 대기한다.
	time.Sleep(100 * time.Millisecond)
}
