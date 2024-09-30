package study

import (
	"fmt"
)

func PrintFmt() {
	fmt.Println("Hello, World!")

	// Printf는 형식 지정자에 따라 출력
	name := "Alice"
	age := 30
	fmt.Printf("제 이름은 %s이고, 나이는 %d살입니다.\n", name, age)

	// Sprintf는 출력하지 않고 형식화된 문자열을 반환
	message := fmt.Sprintf("환영합니다, %s!", name)
	fmt.Println(message)

	// Print는 여러 인수를 공백 없이 출력
	fmt.Print("이것은 ", "한 줄로 ", "출력됩니다.\n")

	// Errorf는 형식화된 오류 메시지를 생성
	err := fmt.Errorf("오류 발생: %s", "무언가 잘못되었습니다")
	fmt.Println(err)

	// 사용자 입력을 받는 Scan 함수 사용
	var inputName string
	var inputAge int

	// 사용자에게 이름과 나이를 입력받음
	fmt.Print("이름을 입력하세요: ")
	fmt.Scan(&inputName)

	fmt.Print("나이를 입력하세요: ")
	fmt.Scan(&inputAge)

	// 입력된 내용을 출력
	fmt.Printf("입력된 이름: %s, 입력된 나이: %d\n", inputName, inputAge)
}
