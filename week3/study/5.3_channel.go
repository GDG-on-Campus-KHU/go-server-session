// < 5.3 채널 >
package study

import "fmt"

// 코드 5.2: 숫자 제곱 프로그램을 통한 동시성 구현
func squareItChannel(inputChan, outputChan chan int) {
	// inputChan에서 값을 받아 처리
	for x := range inputChan {
		outputChan <- x * x
	}
}
func Code52() {
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// 고루틴 실행
	go squareItChannel(inputChannel, outputChannel)

	// 입력값을 inputChannel에 전송
	for i := 0; i < 10; i++ {
		inputChannel <- i
	}
	// 입력이 끝났으므로 inputChannel을 닫는다.
	close(inputChannel)

	// outputChannel에서 값을 받아 출력함.
	for i := range outputChannel {
		fmt.Println(i)
	}

}

// 추가 코드 : 버퍼 개수를 지정하여 데드락을 방지하는 방법
func squareItChannel2(inputChan, outputChan chan int) {
	// inputChan에서 값을 받아 처리
	for x := range inputChan {
		outputChan <- x * x
	}
}
func Code520() {
	inputChannel := make(chan int)
	outputChannel := make(chan int, 10) // 버퍼를 10으로 설정하면 데드락 발생하지 않음

	// 고루틴 실행
	go squareItChannel2(inputChannel, outputChannel)

	// 입력값을 inputChannel에 전송
	for i := 0; i < 10; i++ {
		inputChannel <- i
	}

	// 입력이 끝났으므로 inputChannel을 닫는다.
	// close(inputChannel)

	// outputChannel의 버퍼 크기만큼 반복하여 값을 출력
	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChannel)
	}
}

// 추가 코드 : Close 함수를 사용하여 채널을 닫는 방법
func squareItChannel3(inputChan, outputChan chan int) {
	// inputChan에서 값을 받아 처리하고 outputChan으로 전송
	for x := range inputChan {
		outputChan <- x * x
	}
	// 모든 작업이 끝났음을 알리기 위해 outputChan을 닫음
	close(outputChan)
}

func Code521() {
	inputChannel := make(chan int)
	outputChannel := make(chan int, 10) // 버퍼를 사용하지 않음

	// 고루틴 실행
	go squareItChannel3(inputChannel, outputChannel)

	// 입력값을 inputChannel에 전송
	for i := 0; i < 10; i++ {
		inputChannel <- i
	}
	// inputChannel에 데이터 전송 완료 후 채널을 닫음
	close(inputChannel)

	// outputChannel에서 값을 읽어 출력
	for result := range outputChannel {
		fmt.Println(result)
	}
}
