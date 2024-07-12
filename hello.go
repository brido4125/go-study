package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

// 이름이 main 으로 시작하면 전부 시작 가능
func main() {
	randomNum := rand.Intn(100)
	tryCount := 0

	for {
		fmt.Printf("숫자값을 입력하세요")
		inputNumber, err := InputIntValue()
		tryCount += 1
		if err != nil {
			fmt.Println("입력 오류발생")
		}
		if inputNumber < randomNum {
			fmt.Println("up")
		} else if inputNumber > randomNum {
			fmt.Println("down")
		} else if inputNumber == randomNum {
			fmt.Println("정답 = ", inputNumber)
			fmt.Println("Try count: ", tryCount)
			break
		}
	}

}

func InputIntValue() (int, error) {
	var stdin = bufio.NewReader(os.Stdin)
	var n int
	_, err := fmt.Scanln(&n) // 값 변경하려면 포인터 값 넘겨준다. -> 실제 값 넣으면 그 크기만큼 복사되고 해당 메서드 내에서만 변경됨
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
