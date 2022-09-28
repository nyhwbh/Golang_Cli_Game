package weapons

import "fmt"

//종족별 무기 선택
func ChooseHunmanWeapons() {
	fmt.Printf("(1.no equip 2.Short sword 3.Long sword)  ")
}

func ChooseElfWeapons() {
	fmt.Printf("(1.no equip 2.Short bow 3.Long bow)  ")
}

func ChooseOakWeapons() {
	fmt.Printf("(1.no equip 2.Short Axe 3.Iron hammer)  ")
}

//선택한 무기 공격력, 공격속도에 적욕
func ApplyWeapons(species int, weapon int, character_ATK *int, character_ATKS *int) {
	switch species {
	case 1:
		switch weapon {
		case 1:
			noEquip()
		case 2:
			shortSword(character_ATK)
		case 3:
			longSword(character_ATK)
		}
	case 2:
		switch weapon {
		case 1:
			noEquip()
		case 2:
			shortbow(character_ATKS)
		case 3:
			longbow(character_ATKS)
		}
	case 3:
		switch weapon {
		case 1:
			noEquip()
		case 2:
			shortAxe(character_ATK, character_ATKS)
		case 3:
			ironHammer(character_ATK, character_ATKS)
		}
	}
}

// 장비 없음
func noEquip() {

}

// 각 종족별 무기 옵션 적용
func shortSword(character_ATK *int) int {
	*character_ATK = *character_ATK + int(float64(*character_ATK)*0.05)
	return *character_ATK
}

func longSword(character_ATK *int) int {
	*character_ATK = *character_ATK + int(float64(*character_ATK)*0.1)
	return *character_ATK
}

func shortbow(character_ATKS *int) int {
	*character_ATKS = *character_ATKS - int(float64(*character_ATKS)*0.05)
	return *character_ATKS
}

func longbow(character_ATKS *int) int {
	*character_ATKS = *character_ATKS - int(float64(*character_ATKS)*0.1)
	return *character_ATKS
}

func shortAxe(character_ATK *int, character_ATKS *int) (int, int) {
	*character_ATK = *character_ATK + int(float64(*character_ATK)*0.1)
	*character_ATKS = *character_ATKS + int(float64(*character_ATKS)*0.05)
	return *character_ATK, *character_ATKS
}

func ironHammer(character_ATK *int, character_ATKS *int) (int, int) {
	*character_ATK = *character_ATK + int(float64(*character_ATK)*0.2)
	*character_ATKS = *character_ATKS + int(float64(*character_ATKS)*0.1)
	return *character_ATK, *character_ATKS
}
