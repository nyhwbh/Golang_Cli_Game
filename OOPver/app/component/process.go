package component

import (
	"fmt"
	"strconv"

	"github.com/gosuri/uilive"
	. "github.com/logrusorgru/aurora"
)

type Process struct {
	player  CharacterStatus
	monster Monsters
}

// 캐릭터 생성
func (p *Process) CreateCharacter(selection string) {
	newCharacter := p.player

	writer := uilive.New()
	writer.Start()

	newCharacter.BasicSetup()
	fmt.Println(Bold("케릭터를 생성합니다."))

	//종족 선택
	fmt.Fprintf(writer, "종족 선택\n")
	fmt.Fprintf(writer.Newline(), "1.인간 : 무기 - 검, 특징 - 공격력과 방어력이 높다 \n")
	fmt.Fprintf(writer.Newline(), "2.엘프 : 무기 - 활, 특징 - 공격속도와 회피율이 높다 \n")
	fmt.Fprintf(writer.Newline(), "3.오크 : 무기 - 둔기, 특징 - 공격속도가 느리지만 공격력이 매우 높다 \n")
	fmt.Fprintf(writer.Newline(), "케릭터의 종족을 선택해 주세요(1.인간 2.엘프 3.오크) :  ")
	fmt.Scanln(&selection)

	userSelection := checkSelection(selection, 3)

	newCharacter.SetSpecies(userSelection)

	//무기 선택
	switch userSelection {
	case 1:
		fmt.Fprintf(writer, "무기 선택\n")
		fmt.Fprintf(writer.Newline(), "1.Fist        : 기본적인 무기, 특별한 효과가 없다. \n")
		fmt.Fprintf(writer.Newline(), "2.Short sword : 인간전용 무기, 공격력 5%% 증가 \n")
		fmt.Fprintf(writer.Newline(), "3.Long sword  : 인간전용 무기, 공격력 10%% 증가 \n")
		fmt.Fprintf(writer.Newline(), "무기를 선택해 주세요(1.Fist 2.Short sword 3.Long sword) : ")
	case 2:
		fmt.Fprintf(writer, "무기 선택\n")
		fmt.Fprintf(writer.Newline(), "1.Fist        : 기본적인 무기, 특별한 효과가 없다. \n")
		fmt.Fprintf(writer.Newline(), "2.Short bow : 엘프전용 무기, 공격속도 5%% 증가 \n")
		fmt.Fprintf(writer.Newline(), "3.Iron bow  : 엘프전용 무기, 공격속도 10%% 증가 \n")
		fmt.Fprintf(writer.Newline(), "무기를 선택해 주세요(1.Fist 2.Short bow 3.Iron bow) : ")
	case 3:
		fmt.Fprintf(writer, "무기 선택\n")
		fmt.Fprintf(writer.Newline(), "1.Fist        : 기본적인 무기, 특별한 효과가 없다. \n")
		fmt.Fprintf(writer.Newline(), "2.Short Axe   : 오크전용 무기, 공격속도 5%% 감소, 공격력 10%% 증가\n")
		fmt.Fprintf(writer.Newline(), "3.Iron Hammer : 오크전용 무기, 공격속도 10%% 감소, 공격력 20%% 증가 \n")
		fmt.Fprintf(writer.Newline(), "무기를 선택해 주세요(1.Fist 2.Short Axe 3.Iron Hammer) : ")
	}
	fmt.Scanln(&selection)

	userSelection = checkSelection(selection, 3)
	newCharacter.SetWeapon(userSelection)

	//스킬 생성
	newCharacter.CreateSkills()

	fmt.Fprintf(writer, "%s \n", Bold("케릭터가 생성 되었습니다."))
	fmt.Fprintf(writer.Newline(), "종족 : %s \n", stringSpecies(newCharacter.species))
	fmt.Fprintf(writer.Newline(), "스테이터스 : ")
	fmt.Fprintf(writer.Newline(), "Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
		newCharacter.level,
		newCharacter.healthPoint, newCharacter.maxHealthPoint,
		newCharacter.magicPoint, newCharacter.maxMagicPoint,
		newCharacter.battleAttackPower,
		newCharacter.battleDefensivePower)
	fmt.Fprintf(writer.Newline(), " \n")

	writer.Stop()
}

// 입력 확인하기
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

func stringSpecies(species int) string {
	switch species {
	case 1:
		return "인간"
	case 2:
		return "엘프"
	case 3:
		return "오크"
	}
	return "선택안함"
}
