package lab

import (
	"fmt"
)

// 은행 계좌
type BankAccount struct {
	Balance int
}

// 입금 함수
func (a *BankAccount) Deposit(amount int) error {
	if amount < 0 {
		return fmt.Errorf("음수 금액은 입금할 수 없습니다: %d원", amount)
	}
	a.Balance += amount
	return nil
}

// 출금 함수
func (a *BankAccount) Withdraw(amount int) error {
	if amount > a.Balance {
		return fmt.Errorf("잔액이 부족합니다: 현재 잔액 %d원, 요청한 금액: %d원", a.Balance, amount)
	} else if amount < 0 {
		return fmt.Errorf("음수 금액은 출금할 수 없습니다: %d원", amount)
	}
	a.Balance -= amount
	return nil
}

func Lab3() {
	account := BankAccount{Balance: 100}
	var amount int

	var command int
	for {
		fmt.Print("입금 (1), 출금 (2), 종료 (Others) : ")
		fmt.Scanln(&command)

		if command == 1 {
			fmt.Print("입금할 금액을 입력하세요: ")
			fmt.Scanln(&amount)

			if err := account.Deposit(amount); err != nil {
				fmt.Println("오류:", err)
			} else {
				fmt.Printf("입금 성공! 현재 잔액: %d원\n", account.Balance)
			}
		} else if command == 2 {
			fmt.Print("출금할 금액을 입력하세요: ")
			fmt.Scanln(&amount)

			if err := account.Withdraw(amount); err != nil {
				fmt.Println("오류:", err)
			} else {
				fmt.Printf("출금 성공! 현재 잔액: %d원\n", account.Balance)
			}
		} else {
			break
		}
		fmt.Println()
	}
}
