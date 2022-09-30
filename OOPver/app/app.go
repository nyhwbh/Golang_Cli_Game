package app

import (
	"fmt"
	"strconv"

	"cligame/app/component"

	. "github.com/logrusorgru/aurora"
)

func RunGame() {

	selection := ""

	fmt.Printf("게임을 시작하시겠습니까? (1.Yes/2.No)  ")
	fmt.Scanln(&selection)
	fmt.Println(" ")

	for {

		userSelection := checkSelection(selection, 2)

		if userSelection == 1 {
			fmt.Println(Bold("게임을 시작합니다."))

			//캐릭터 생성
			char := component.CharacterStatus{}

			char.BasicSetup()
			fmt.Println("케릭터가 생성 되었습니다.")

			//종족 선택
			shortExplainSpecies()
			fmt.Scanln(&selection)

			userSelection = checkSelection(selection, 3)

			char.SetSpecies(userSelection)

			char.PrintCharterStatus()

			//무기 선택
			shortExplainWeapons(userSelection)
			fmt.Scanln(&selection)

			userSelection = checkSelection(selection, 3)
			char.SetWeapon(userSelection)
			char.PrintCharterStatus()

			//스킬 생성

			//게임 진행

			break
		} else if userSelection == 2 {
			fmt.Println(Bold("게임을 종료합니다."))
			break
		}
	}

}

func checkSelection(selection string, lastNum int) int {
	returnValue := 0
	for {
		userSelection, err := strconv.Atoi(selection)
		if err != nil {
			fmt.Printf("선택지에 해당하는 숫자를 입력해 주세요  ")
			fmt.Scanln(&selection)
		} else {
			if 1 <= userSelection && userSelection <= lastNum {
				returnValue = userSelection
				break
			} else {
				fmt.Printf("선택지에 해당하는 숫자를 입력해 주세요  ")
				fmt.Scanln(&selection)
			}
		}
	}
	return returnValue
}

func shortExplainSpecies() {
	fmt.Println("1.인간 : 무기 - 검, 특징 - 공격력과 방어력이 높다")
	fmt.Println("2.엘프 : 무기 - 활, 특징 - 공격속도와 회피율이 높다")
	fmt.Println("3.오크 : 무기 - 둔기, 특징 - 공격속도가 느리지만 공격력이 매우 높다")
	fmt.Printf("케릭터의 종족을 선택해 주세요(1.인간 2.엘프 3.오크) :  ")
}

func shortExplainWeapons(selection int) {
	switch selection {
	case 1:
		fmt.Println("1.Fist        : 기본적인 무기, 특별한 효과가 없다.")
		fmt.Println("2.Short sword : 인간전용 무기, 공격력 5% 상승")
		fmt.Println("3.Long sword  : 인간전용 무기, 공격력 10% 상승")
		fmt.Printf("무기를 선택해 주세요(1.Fist 2.Short sword 3.Long sword) : ")
	case 2:
		fmt.Println("1.Fist      : 기본적인 무기, 특별한 효과가 없다.")
		fmt.Println("2.Short bow : 엘프전용 무기, 공격속도 5% 상승")
		fmt.Println("3.Iron bow  : 엘프전용 무기, 공격속도 10% 상승")
		fmt.Printf("무기를 선택해 주세요(1.Fist 2.Short bow 3.Iron bow) : ")
	case 3:
		fmt.Println("1.Fist        : 기본적인 무기, 특별한 효과가 없다.")
		fmt.Println("2.Short Axe   : 오크전용 무기, 공격속도 5% 감소, 공격력 10% 증가")
		fmt.Println("3.Iron Hammer : 오크전용 무기, 공격속도 10% 감소, 공격력 20% 증가")
		fmt.Printf("무기를 선택해 주세요(1.Fist 2.Short Axe 3.Iron Hammer) : ")
	}

}
