package process

import (
	"fmt"

	"github.com/nyhwbh/bigingame/battles"
	"github.com/nyhwbh/bigingame/character/status"
	"github.com/nyhwbh/bigingame/character/weapons"
)

// 캐릭터 생성시 종족 선택
func SetCharacterSpecies(species int, selectSpecies *int) {
	for {
		if species >= 1 && species <= 3 {
			fmt.Println(" ")
			*selectSpecies = species
			break
		} else {
			fmt.Printf("원하는 종족에 해당하는 숫자를 입력해 주세요 (1.Human 2.Elf 3.Oak)  ")
			fmt.Scanln(&species)
		}
	}
}

// 캬릭터 생성시 무기 선택, 전투전 무기 선택
func SetCharacterWeapons(weapon int, selectWeapons *int, speciesSelection int) {
	for {
		if weapon >= 1 && weapon <= 3 {
			fmt.Println("")
			*selectWeapons = weapon
			break
		} else {
			fmt.Printf("원하는 무기에 해당하는 숫자를 입력해 주세요 ")
			switch speciesSelection {
			case 1:
				weapons.ChooseHunmanWeapons()
			case 2:
				weapons.ChooseElfWeapons()
			case 3:
				weapons.ChooseOakWeapons()
			}
			fmt.Scanln(&weapon)
		}
	}
}

// 게임 진행
func GameProcess(c *status.CharacterStatus, selection int, gameSelection *int, characterSpecies int, weapon int, selectWeapons *int) {
	status.PrintCharacterStatus(c)
	level := status.Level(c)
	// level 99 달성시 궁극스킬 사용가능 메세지 출력
	if level == 99 {
		fmt.Println("이제 궁극스킬을 사용할 수 있습니다.")
	}
	for {
		if selection == 1 {
			// 몬스터와 전투 시작
			battles.FightWithMonster(c, gameSelection, characterSpecies, *selectWeapons)
			break
		} else if selection == 2 {
			// 휴식 - 체력과 마나 회족
			battles.TakeRest(c)
			break
		} else if selection == 3 {
			// 무기 변경
			SetCharacterWeapons(weapon, selectWeapons, characterSpecies)
			break
		} else if selection == 9 {
			// 게임 종료
			*gameSelection = selection
			break
		} else {
			fmt.Printf("무엇을 하시겠습니까? \n (1.Fight with Monster 2.Take rest 3.Change Weapon  9.Quit game)  ")
			fmt.Scanln(&selection)
			fmt.Println(" ")
		}
	}

}
