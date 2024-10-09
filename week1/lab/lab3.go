package lab

import "fmt"

func SumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func FindMaxMin(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func Lab3() {
	defer fmt.Println("프로그램을 종료합니다.")

	arr := []int{3, 5, 1, 2, 0}

	totalSum := SumArray(arr)
	fmt.Println("배열의 총 합:", totalSum)

	max, min := FindMaxMin(arr)
	fmt.Println("배열의 최대값:", max)
	fmt.Println("배열의 최소값:", min)

	n := len(arr)
	switch {
	case n < 5:
		fmt.Println("배열의 길이가 5보다 작습니다.")
	case n == 5:
		fmt.Println("배열의 길이가 5입니다.")
	case n > 5:
		fmt.Println("배열의 길이가 5보다 큽니다.")
	}
}
