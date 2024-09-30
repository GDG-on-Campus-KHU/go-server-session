package study

import "fmt"

func PrintVariablesTypes() {
	// 정수형 변수
	var intVar int = 42
	fmt.Println("정수형 변수 (int):", intVar)

	// 부동 소수점 변수
	var floatVar float64 = 3.14
	fmt.Println("부동 소수점 변수 (float64):", floatVar)

	// 문자열 변수
	var stringVar string = "Hello, Go!"
	fmt.Println("문자열 변수 (string):", stringVar)

	// 불리언 변수
	var boolVar bool = true
	fmt.Println("불리언 변수 (bool):", boolVar)

	// 상수
	const pi float64 = 3.14159
	fmt.Println("상수 (const float64):", pi)

	// 타입 추론을 사용한 변수 선언
	autoIntVar := 100
	autoStringVar := "자동 타입 추론"
	fmt.Println("자동 타입 추론 정수형 변수 (int):", autoIntVar)
	fmt.Println("자동 타입 추론 문자열 변수 (string):", autoStringVar)

	// 여러 변수 한 번에 선언
	var x, y, z int = 1, 2, 3
	fmt.Println("여러 변수 한 번에 선언 (int):", x, y, z)

	// 여러 상수 한 번에 선언
	const (
		constA = "상수 A"
		constB = "상수 B"
	)
	fmt.Println("여러 상수 한 번에 선언 (const string):", constA, constB)
}
