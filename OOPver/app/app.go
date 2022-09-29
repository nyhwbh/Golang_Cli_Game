package app

import (
	"fmt"
	"strconv"

	"github.com/gosuri/uilive"
	. "github.com/logrusorgru/aurora"
)

func RunGame() {

	selection := ""

	fmt.Printf("%s", Bold("게임을 시작하시겠습니까? (1.Yes/2.No)  "))
	fmt.Scanln(&selection)
	fmt.Println(" ")

	//terminal output controll
	writer := uilive.New()
	writer.Start()

	for {

		userSelection, err := strconv.Atoi(selection)

		if err != nil {
			fmt.Printf("선택지에 해당하는 숫자를 입력해 주세요 (1.Yes / 2.No)  ")
			fmt.Scanln(&selection)
		} else {
			if userSelection == 1 {
				fmt.Println("게임을 시작합니다.")
				break
			} else if userSelection == 2 {
				fmt.Println("게임을 종료합니다.")
				break
			} else {
				fmt.Printf("선택지에 해당하는 숫자를 입력해 주세요 (1.Yes / 2.No)  ")
				fmt.Scanln(&selection)
			}
		}

	}

	writer.Stop()
}
