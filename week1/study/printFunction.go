package study

import "fmt"

// 함수 (리턴타입은 중괄호 앞에 위치)
func valueOfPi() float64 {
	return 3.141592
}

// 인자가 있는 함수 (인자 타입은 변수명 뒤에 위치)
func multiplyByPi(x float32) float32 {
	return x * 3.141592
}

// 인자가 여러 개인 함수
func add(x int, y int) int {
	return x + y
}

// 인자가 가변인 함수 (인자의 개수를 고정하지 않고, 가변적인 수의 인자를 받을 수 있는 함수)
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 리턴값이 여러 개인 함수
func swap(x, y string) (string, string) {
	return y, x
}

func nameAndAge(uid int) (string, int) {
	switch uid {
	case 0:
		return "Alice", 25
	case 1:
		return "Bob", 30
	default:
		return "Unknown", 0
	}
}

// 함수를 인자로 받는 함수
func runMathOp(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

// defer 키워드를 사용한 함수 (함수를 반환하기 전에 특정 코드를 실행하는 명령어)
func test(x int) int {
	defer fmt.Println(x)
	y := x + 1
	fmt.Println(y)
	return y
}

func PrintFunction() {
	fmt.Println("Value of Pi:", valueOfPi())
	fmt.Println("Multiply by Pi:", multiplyByPi(2))
	fmt.Println("Add 3 and 5:", add(3, 5))
	fmt.Println("Sum of 1, 2, 3, 4:", sum(1, 2, 3, 4))

	a, b := swap("hello", "world")
	fmt.Println("Swap:", a, b)

	name, age := nameAndAge(0)
	fmt.Printf("Name and Age (uid 0): %s, %d\n", name, age)
	name, age = nameAndAge(1)
	fmt.Printf("Name and Age (uid 1): %s, %d\n", name, age)
	name, age = nameAndAge(2)
	fmt.Printf("Name and Age (uid 2): %s, %d\n", name, age)

	fmt.Println("Run Math Op (3 + 4):", runMathOp(3, 4, add))

	fmt.Println("Test with defer:")
	test(5)
}
