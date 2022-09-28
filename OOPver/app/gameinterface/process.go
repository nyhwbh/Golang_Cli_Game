package gameinterface

import (
	"fmt"
)

type Selections struct {
	gameStart int
	choice    int
	skill     int
	PlayerStatus
	MonsterStatus
}

// 게임 시작
func (s *Selections) GameStart() {
	for {
		if s.gameStart == 1 {
			fmt.Println("게임이 시작됩니다.")
			fmt.Println("플레이 하실 케릭터를 생성합니다.")
			// 케릭터 생성
			p := PlayerStatus{}
			p.CreatePlayer()
			fmt.Print("플레이 하실 종족을 선택해 주세요(1. Human 2. Elf 3.Oak) : ")
			fmt.Scanln(&s.choice)
			// 종족 설정
			p.PlayerSpeciesSetUp(s.choice)
			fmt.Printf("무기를 선택해 주세요(1. %s 2.%s 3. %s)", p.weapons[0], p.weapons[1], p.weapons[2])
			fmt.Scanln(s.choice)
			s.ChangeWeapon()

		} else if s.gameStart == 2 {
			fmt.Println("게임이 종료됩니다.")
			break
		} else {
			fmt.Println("선택지에 있는 번호를 입력해 주세요")
			fmt.Println("1. 게임시작 2. 종료")
			fmt.Sprintln(&s.gameStart)
		}
	}
}

// 스킬 사용
func (s *Selections) UsingSkill() {
	switch s.PlayerStatus.species {
	case 1:
		switch s.skill {
		case 1:
			s.PlayerStatus.Heal()
			return
		case 2:
			s.PlayerStatus.Steam()
			return
		case 3:
			s.PlayerStatus.Guard()
			return
		case 4:
			if s.PlayerStatus.level >= 99 {
				s.PlayerStatus.Invincible()
				return
			} else {
				fmt.Println("레벨이 부족하여 사용할 수 없습니다.")
				return
			}
		}
	case 2:
		switch s.skill {
		case 1:
			s.PlayerStatus.Heal()
			return
		case 2:
			s.PlayerStatus.Steam()
			return
		case 3:
			s.PlayerStatus.Elusion()
			return
		case 4:
			if s.PlayerStatus.level >= 99 {
				s.PlayerStatus.Rapid()
				return
			} else {
				fmt.Println("레벨이 부족하여 사용할 수 없습니다.")
				return
			}
		}
	case 3:
		switch s.skill {
		case 1:
			s.PlayerStatus.Heal()
			return
		case 2:
			s.PlayerStatus.Steam()
			return
		case 3:
			s.PlayerStatus.Anger()
			return
		case 4:
			if s.PlayerStatus.level >= 99 {
				s.PlayerStatus.Frenzy()
				return
			} else {
				fmt.Println("레벨이 부족하여 사용할 수 없습니다.")
				return
			}
		}
	}
}

// 무기변경 및 선택
func (s *Selections) ChangeWeapon() {
	switch s.PlayerStatus.species {
	case 1:
		switch s.choice {
		case 1:
			s.PlayerStatus.Fist()
			return
		case 2:
			s.PlayerStatus.ShortSword()
			return
		case 3:
			s.PlayerStatus.LongSword()
			return
		}
	case 2:
		switch s.choice {
		case 1:
			s.PlayerStatus.Fist()
			return
		case 2:
			s.PlayerStatus.ShortBow()
			return
		case 3:
			s.PlayerStatus.LongBow()
			return
		}
	case 3:
		switch s.choice {
		case 1:
			s.PlayerStatus.Fist()
			return
		case 2:
			s.PlayerStatus.ShortAxe()
			return
		case 3:
			s.PlayerStatus.IronHammer()
			return
		}
	}
}
