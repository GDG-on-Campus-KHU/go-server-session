// < 3.2 서드파티 패키지 사용하기 >
package study

import (
	"fmt"
	"os"
	"strconv"

	"github.com/otiai10/primes"
)

// 코드 3.2 : otiai10 패키지를 이용하여 소수인지 확인하기
// go get github.com/otiai10/primes 명령어를 통해 패키지를 설치한 후 실행
func Code32() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage:", os.Args[0], "<number>")
		os.Exit(1)
	}
	number, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	f := primes.Factorize(int64(number))
	fmt.Println("primes:", len(f.Powers()) == 1)
}
