package app

import (
	"fmt"

	"cligame/app/component"

	. "github.com/logrusorgru/aurora"
)

func RunGame() {

	selection := ""

	fmt.Printf("게임을 시작하시겠습니까? (1.예/2.아니요)  ")
	fmt.Scanln(&selection)
	fmt.Println(" ")

	for {

		userSelection := component.CheckSelection(selection, 2)

		if userSelection == 1 {
			fmt.Println(Bold("게임을 시작합니다."))

			//캐릭터 생성
			process := component.Process{}

			process.CreateCharacter(selection)

			process.GameProcess()

			break
		} else if userSelection == 2 {
			fmt.Println(Bold("게임을 종료합니다."))
			break
		}
	}

}
