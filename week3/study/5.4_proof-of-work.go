// < 5.4 작업 증명 2 >
package study

import (
	"crypto/sha256"
	"fmt"
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
)

// 작업 증명 알고리즘은 블록체인의 핵심 구성 요소 중 하나이다.
// 동시성을 통해 작업 증명 최적화를 구현할 수 있다.

// 코드 5.6 : 새롭게 개선된 고루틴 기반 pow 함수
func pow(prefix string, bitLength int) {
	start := time.Now()      // 함수 시작 시간을 기록한다.
	hash := []int{}          // 각 CPU별 해시 계산 횟수를 저장할 배열을 선언한다.
	totalHasesProcessed := 0 // 전체 해시 처리 횟수를 저장할 변수를 선언한다.

	numberOfCPU := runtime.NumCPU()      // 현재 시스템의 CPU 코어 수를 가져온다.
	closeChan := make(chan int, 1)       // 고루틴을 종료하기 위한 채널을 생성한다.
	solutionChan := make(chan []byte, 1) // 해시 계산이 성공한 결과를 전달할 채널을 생성한다.

	// 각 CPU 코어별로 고루틴을 생성하여 병렬로 해시를 계산한다.
	for idx := 0; idx < numberOfCPU; idx++ {
		hash = append(hash, 0) // 각 CPU 코어에 대응하는 해시 계산 횟수 저장을 위한 초기값 추가.

		// 고루틴 생성, 인자로 idx 값을 넘겨주어 고루틴 내에서 고유한 값으로 사용하게 한다.
		go func(hasIndex int) {
			seed := uint64(time.Now().UnixNano())                // 난수를 생성하기 위한 시드를 현재 시간으로 설정한다.
			randomBytes := make([]byte, 20)                      // 난수를 저장할 20바이트 배열을 생성한다.
			randomBytes = append([]byte(prefix), randomBytes...) // prefix를 랜덤 바이트 배열에 추가하여 새로운 배열을 만든다.

			for {
				select {
				// closeChan에서 값이 들어오면 고루틴을 종료한다.
				case <-closeChan:
					closeChan <- 1
					return
				// 매 나노초마다 작업을 수행한다.
				case <-time.After(time.Nanosecond):
					count := 0         // 해시 시도 횟수를 기록하기 위한 변수.
					for count < 5000 { // 5000번의 해시를 시도한다.
						count++
						seed = RandomString(randomBytes, len(prefix), seed) // 랜덤 문자열을 생성하여 seed를 갱신한다.
						// 해시 계산이 성공하면 결과를 solutionChan에 전송하고, 고루틴을 종료한다.
						if Hash(randomBytes, bitLength) {
							hash[hasIndex] += count     // 해당 CPU 코어에서의 해시 처리 횟수를 기록한다.
							solutionChan <- randomBytes // 성공한 해시 결과를 solutionChan에 전송한다.
							closeChan <- 1              // 다른 고루틴들을 종료하기 위해 closeChan에 값을 전송한다.
							return
						}
					}
					hash[hasIndex] += count // CPU 코어별 해시 처리 횟수를 업데이트한다.
				}
			}
		}(idx) // idx 값을 고루틴에 넘겨주어 고유한 값으로 유지한다.
	}

	<-solutionChan // solutionChan에서 성공한 해시 결과를 받는다.

	// 각 CPU 코어에서 처리한 해시 횟수를 모두 합산한다.
	for _, v := range hash {
		totalHasesProcessed += v
	}

	end := time.Now() // 함수 종료 시간을 기록한다.

	// 총 실행 시간을 출력한다.
	fmt.Println("time:", end.Sub(start).Seconds())
	// 총 해시 처리 횟수를 출력한다.
	fmt.Println("processed:", humanize.Comma(int64(totalHasesProcessed)))
	// 초당 처리한 해시 수를 계산하고 출력한다.
	fmt.Printf("processed/sec: %s\n", humanize.Comma(int64(float64(totalHasesProcessed)/end.Sub(start).Seconds())))
}

// RandomString generates a random string of the given length using the provided seed.
func RandomString(randomBytes []byte, prefixLen int, seed uint64) uint64 {
	for i := prefixLen; i < len(randomBytes); i++ {
		randomBytes[i] = byte(seed & 0xFF)
		seed >>= 8
	}
	return seed
}

// Hash checks if the hash of the given data meets the required bit length.
func Hash(data []byte, bitLength int) bool {
	hash := sha256.Sum256(data)
	for i := 0; i < bitLength/8; i++ {
		if hash[i] != 0 {
			return false
		}
	}
	if bitLength%8 != 0 {
		if hash[bitLength/8]>>(8-bitLength%8) != 0 {
			return false
		}
	}
	return true
}
