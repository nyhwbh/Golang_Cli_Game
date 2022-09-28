package skills

import (
	"fmt"

	"github.com/nyhwbh/bigingame/character/status"
)

//종족별 스킬 사용
func ApplySkills(c *status.CharacterStatus, level int, speies int, character_hp *int, character_ATK *int, character_ATKS *int, character_Avoid *int, character_DEF *int, monster_ATKS *int) {
	skill, skillSelection := 0, 0
	if level < 99 {
		chooseSkill(skill, &skillSelection, level)
		switch skillSelection {
		case 1:
			skillHeal(c, character_hp)
			status.PrintCharacterStatusInbattle(c, *character_hp, *character_ATK, *character_DEF)
		case 2:
			skillSteam(c, character_ATK)
		case 3:
			switch speies {
			case 1:
				skillGuard(c, character_DEF)
			case 2:
				skillElusion(c, character_Avoid)
			case 3:
				skillAnger(c, character_ATK, character_DEF)
			}
		}
	} else {
		chooseSkill(skill, &skillSelection, level)
		switch skillSelection {
		case 1:
			skillHeal(c, character_hp)
			status.PrintCharacterStatusInbattle(c, *character_hp, *character_ATK, *character_DEF)
		case 2:
			skillSteam(c, character_ATK)
		case 3:
			switch speies {
			case 1:
				skillGuard(c, character_DEF)
			case 2:
				skillElusion(c, character_Avoid)
			case 3:
				skillAnger(c, character_ATK, character_DEF)
			}
		case 4:
			switch speies {
			case 1:
				skillInvincible(c, monster_ATKS)
			case 2:
				skillRapid(c, character_ATKS)
			case 3:
				skillFrenzy(c, character_ATK)
			}
		}
	}
}

//스킬 저장
var bagic_skills []string

//종족별 스킬저장
func CreateSkills() {
	bagic_skills = append(bagic_skills, "Heal", "Steam")
}

func AddHumanSkill() {
	bagic_skills = append(bagic_skills, "Gaurd")
}

func AddHumanHiperSkill() {
	bagic_skills = append(bagic_skills, "Invincible")
}

func AddElfSkills() {
	bagic_skills = append(bagic_skills, "Elusion")
}

func AddElfHiperSkills() {
	bagic_skills = append(bagic_skills, "Rapid")
}

func AddOakSkills() {
	bagic_skills = append(bagic_skills, "Anger")
}

func AddOakHiperSkills() {
	bagic_skills = append(bagic_skills, "Frenzy")
}

//궁극기 추가
func AddHiperSKill(species int) {
	switch species {
	case 1:
		AddHumanHiperSkill()
	case 2:
		AddElfHiperSkills()
	case 3:
		AddOakHiperSkills()
	}
}

//스킬사용
func chooseSkill(skill int, skillSelection *int, level int) {
	for {
		if level > 99 {
			if skill > 0 && skill < 4 {
				*skillSelection = skill
				break
			} else if skill == 4 {
				fmt.Println("아직은 사용할 수 없는 스킬입니다.")
				fmt.Printf("사용하실 스킬의 숫자를 입력해 주세요 (1.%s 2.%s 3.%s)  ", bagic_skills[0], bagic_skills[1], bagic_skills[2])
				fmt.Scanln(&skill)
			} else {
				fmt.Printf("사용하실 스킬의 숫자를 입력해 주세요 (1.%s 2.%s 3.%s)  ", bagic_skills[0], bagic_skills[1], bagic_skills[2])
				fmt.Scanln(&skill)
			}
		} else {
			if skill > 0 && skill < 5 {
				*skillSelection = skill
				break
			} else {
				fmt.Printf("사용하실 스킬의 숫자를 입력해 주세요 (1.%s 2.%s 3.%s 4.%s)  ", bagic_skills[0], bagic_skills[1], bagic_skills[2], bagic_skills[3])
				fmt.Scanln(&skill)
			}
		}
	}
}

//스킬 적용
func skillHeal(c *status.CharacterStatus, character_hp *int) int {
	check_mp := status.UsingMagicPoint(c)
	c_hp := status.CallCharacterHp(c)
	c_maxhp := status.CallCharacterMaxHp(c)
	if check_mp == true {
		if (c_hp + 30) > c_maxhp {
			*character_hp = c_maxhp
			return *character_hp
		} else {
			return *character_hp + 30
		}
	} else {
		return *character_hp
	}
}

func skillSteam(c *status.CharacterStatus, character_ATK *int) int {
	check_mp := status.UsingMagicPoint(c)
	if check_mp == true {
		*character_ATK = *character_ATK + int(float64(*character_ATK)*0.2)
		return *character_ATK
	} else {
		return *character_ATK
	}
}

func skillGuard(c *status.CharacterStatus, character_DEF *int) int {
	check_mp := status.UsingMagicPoint(c)
	if check_mp == true {
		*character_DEF = *character_DEF + int(float64(*character_DEF)*0.3)
		return *character_DEF
	} else {
		return *character_DEF
	}
}

func skillElusion(c *status.CharacterStatus, charactor_Avoid *int) int {
	check_mp := status.UsingMagicPoint(c)
	if check_mp == true {
		return *charactor_Avoid + 30
	} else {
		return *charactor_Avoid
	}
}

func skillAnger(c *status.CharacterStatus, charactor_ATK *int, charactor_DEF *int) (int, int) {
	check_mp := status.UsingMagicPoint(c)
	if check_mp == true {
		*charactor_ATK = *charactor_ATK + int(float64(*charactor_ATK)*0.5)
		*charactor_DEF = *charactor_DEF - int(float64(*charactor_DEF)*0.1)
		return *charactor_ATK, *charactor_DEF
	} else {
		return *charactor_ATK, *charactor_DEF
	}
}

func skillInvincible(c *status.CharacterStatus, monster_ATKS *int) int {
	check_mp := status.UsingMagicPoint(c)
	if check_mp == true {
		*monster_ATKS = 10000
		return *monster_ATKS
	} else {
		return *monster_ATKS
	}
}

func skillRapid(c *status.CharacterStatus, charactor_ATKS *int) int {
	check_mp := status.UsingMagicPoint(c)
	if check_mp == true {
		return *charactor_ATKS / 5
	} else {
		return *charactor_ATKS
	}
}

func skillFrenzy(c *status.CharacterStatus, charactor_ATK *int) int {
	check_mp := status.UsingMagicPoint(c)
	if check_mp == true {
		return *charactor_ATK * 5
	} else {
		return *charactor_ATK
	}
}
