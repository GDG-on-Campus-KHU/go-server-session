package lab1

import "fmt"

// 실습1
// 더하기 함수
func sum(input, output chan int) {
	// 채널이 열려있기에
	num1 := <-input
	num2 := <-input

	output <- num1 + num2
}

func mul(input, output chan int) {
	num1 := <-input
	num2 := <-input

	output <- num1 + num2
}

func main() {
	sumInputChannel := make(chan int)
	mulInputChannel := make(chan int)
	sumOutputChannel := make(chan int)
	mulOutputChannel := make(chan int)

	go sum(sumInputChannel, sumOutputChannel)
	go mul(mulInputChannel, mulOutputChannel)

	inputNum1 := 0
	inputNum2 := 0
	fmt.Print("첫 번째 정수를 입력하세요: ")
	fmt.Scan(&inputNum1)
	fmt.Print("두 번째 정수를 입력하세요: ")
	fmt.Scan(&inputNum2)

	sumInputChannel <- inputNum1
	sumInputChannel <- inputNum2
	mulInputChannel <- inputNum1
	mulInputChannel <- inputNum2

	fmt.Println("덧셈 결과는: ", <-sumOutputChannel)
	fmt.Println("곱셈 결과는: ", <-mulOutputChannel)

}
