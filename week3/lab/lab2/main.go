package lab2

import "fmt"

// 실습2
func main() {
	Input1Channel := make(chan int)
	Input2Channel := make(chan int)
	Output1Channel := make(chan int)
	Output2Channel := make(chan int)

	exitChannel := make(chan bool)

	go func(Input1Channel, Input2Channel, Output1Channel, Output2Channel chan int, exitChannel chan bool) {
		sum := 0
		mul := 1
		for {
			select {
			case x := <-Input1Channel:
				sum += x
				mul *= x
				Output1Channel <- sum
				Output2Channel <- mul
			case y := <-Input2Channel:
				sum += y
				mul *= y
				Output1Channel <- sum
				Output2Channel <- mul
			case <-exitChannel:
				return
			}
		}
	}(Input1Channel, Input2Channel, Output1Channel, Output2Channel, exitChannel)

	var inputNum1, inputNum2 int
	fmt.Printf("첫 번째 정수를 입력하세요: ")
	fmt.Scan(&inputNum1)
	fmt.Printf("두 번째 정수를 입력하세요: ")
	fmt.Scan(&inputNum2)

	Input1Channel <- inputNum1
	Input2Channel <- inputNum2

	fmt.Printf("덧셈 결과: %d\n", <-Output1Channel)
	fmt.Printf("곱셈 결과: %d\n", <-Output2Channel)

	exitChannel <- true
}
