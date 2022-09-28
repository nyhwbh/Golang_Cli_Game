package main

import (
	"fmt"

	"github.com/nyhwbh/learn_go/cli_game/Functionalver/character/skills"
	"github.com/nyhwbh/learn_go/cli_game/Functionalver/character/status"
	"github.com/nyhwbh/learn_go/cli_game/Functionalver/process"
)

func main() {
	selection, gameSelection := 0, 0
	species, speciesSelection := 0, 0
	weapon, weaponSelection := 0, 0
	fmt.Printf("게임을 시작하시겠습니까? (1.Yes/2.No)  ")
	fmt.Scanln(&selection)
	fmt.Println(" ")
	for {
		if selection == 1 {

			//캐릭터 생성
			newCharacter := status.CharacterStatus{}
			status.BasicSetup(&newCharacter)

			//종족 선택
			process.SetCharacterSpecies(species, &speciesSelection)
			status.SetSpecies(&newCharacter, speciesSelection)

			// 무기 선택
			process.SetCharacterWeapons(weapon, &weaponSelection, speciesSelection)

			//스킬 생성
			skills.CreateSkills()
			switch speciesSelection {
			case 1:
				skills.AddHumanSkill()
				skills.AddHumanHiperSkill()
			case 2:
				skills.AddElfSkills()
				skills.AddElfHiperSkills()
			case 3:
				skills.AddOakSkills()
				skills.AddOakHiperSkills()
			}

			fmt.Println("캐릭터가 생성되었습니다.")
			status.PrintCharacterStatus(&newCharacter)
			fmt.Println(" ")

			//게임 진행
			for {
				selection = 0
				process.GameProcess(&newCharacter, selection, &gameSelection, speciesSelection, weapon, &weaponSelection)
				if gameSelection == 9 {
					break
				}
			}
			fmt.Println("게임을 종료합니다.")
			break
		} else if selection == 2 {
			fmt.Println("게임을 종료합니다.")
			break
		} else {
			fmt.Printf("선택지에 해당하는 숫자를 입력해 주세요 (1.Yes/2.No)  ")
			fmt.Scanln(&selection)
		}
	}
}
