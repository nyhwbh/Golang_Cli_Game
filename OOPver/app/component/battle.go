package component

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gosuri/uilive"
	. "github.com/logrusorgru/aurora"
)

// 몬스터와 전투
func (p *Process) battleWithMonster() {
	//몬스터 생성
	p.monster.CreateMoster()

	fmt.Printf("\n")
	fmt.Println(Bold("몬스터와 전투를 시작합니다."))
	fmt.Printf("\n")
	// 전투
	p.fightWithMoster()
	// 스킬사용
	// p.useingSkills()

}

// 케릭터와 몬스터 공격속도에 따른 게임 진행
func (p *Process) fightWithMoster() {
	characterAttack, monsterAttack := 1, 1
	fightTime := 0
	writer := uilive.New()
	writer.Start()
	for {
		//10 Microsecond 마다 전두시간 증가 -> 공격 속도에 따른 딜레이 주기
		time.Sleep(time.Millisecond * 10)
		fightTime += 10
		playerDamage := p.player.battleAttackPower - p.monster.depensivePower
		monsterDamage := p.monster.attackPower - p.player.battleDefensivePower

		timeSource := rand.NewSource(time.Now().UnixNano())
		random := rand.New(timeSource)
		counterAttack := random.Intn(100)
		avoidAttack := random.Intn(100)

		// 케릭터와 몬스터가 동시에 공격
		if fightTime == (p.player.battleAttackSpeed*characterAttack) &&
			fightTime == (p.monster.attackSpeed*monsterAttack) {
			//케릭터 공격
			characterAttack++
			if counterAttack > p.monster.counterAttack {
				p.monster.MonsterAttacked(playerDamage)
				fmt.Fprintf(writer, "%s \n", Bold(Blue("몬스터를 공격하였습니다. ")))
				fmt.Fprintf(writer.Newline(), "%d의 데미지를 입혔습니다. \n", Bold(Blue(playerDamage)))
				fmt.Fprintf(writer.Newline(), " \n")
				fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
					p.player.level,
					p.player.healthPoint, p.player.maxHealthPoint,
					p.player.magicPoint, p.player.maxMagicPoint,
					p.player.battleAttackPower,
					p.player.battleDefensivePower)
				fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
					p.monster.healthPoint, p.monster.maxHealthPoint,
					p.monster.attackPower,
					p.monster.depensivePower)
				fmt.Fprintf(writer.Newline(), " \n")
				if p.checkHealthPoint() {
					writer.Stop()
					break
				}
			} else {
				if avoidAttack > p.player.battleAvoid {
					// 몬스터의 반격
					monsterDamage = int(float64(monsterDamage) * 0.7)
					p.player.CharacterAttacked(monsterDamage)
					fmt.Fprintf(writer, "%s \n", Bold(Blue("몬스터를 공격하였습니다. ")))
					fmt.Fprintf(writer.Newline(), "%s \n", Bold(Red("몬스터의 반격!! ")))
					fmt.Fprintf(writer.Newline(), "%d의 데미지를 입었습니다. \n", Bold(Red(monsterDamage)))
					fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
						p.player.level,
						p.player.healthPoint, p.player.maxHealthPoint,
						p.player.magicPoint, p.player.maxMagicPoint,
						p.player.battleAttackPower,
						p.player.battleDefensivePower)
					fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
						p.monster.healthPoint, p.monster.maxHealthPoint,
						p.monster.attackPower,
						p.monster.depensivePower)
					fmt.Fprintf(writer.Newline(), "\n \n")
					if p.checkHealthPoint() {
						writer.Stop()
						break
					}
				} else {
					// 몬스터의 반격 회피
					fmt.Fprintf(writer, "%s \n", Bold(Blue("몬스터를 공격하였습니다. ")))
					fmt.Fprintf(writer.Newline(), "%s \n", Bold(Red("몬스터의 반격!! ")))
					fmt.Fprintf(writer.Newline(), "%s \n", Bold(Blue("몬스터의 공격을 회피하였습니다.")))
					fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
						p.player.level,
						p.player.healthPoint, p.player.maxHealthPoint,
						p.player.magicPoint, p.player.maxMagicPoint,
						p.player.battleAttackPower,
						p.player.battleDefensivePower)
					fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
						p.monster.healthPoint, p.monster.maxHealthPoint,
						p.monster.attackPower,
						p.monster.depensivePower)
					fmt.Fprintf(writer.Newline(), "\n \n")
				}
			}
			// 몬스터 공격
			monsterAttack++
			if avoidAttack > p.player.battleAvoid {
				p.player.CharacterAttacked(monsterDamage)
				fmt.Fprintf(writer, "%s \n", Bold(Red("몬스터가 공격하였습니다. ")))
				fmt.Fprintf(writer.Newline(), "%d의 데미지를 입었습니다. \n", Bold(Red(monsterDamage)))
				fmt.Fprintf(writer.Newline(), " \n")
				fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
					p.player.level,
					p.player.healthPoint, p.player.maxHealthPoint,
					p.player.magicPoint, p.player.maxMagicPoint,
					p.player.battleAttackPower,
					p.player.battleDefensivePower)
				fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
					p.monster.healthPoint, p.monster.maxHealthPoint,
					p.monster.attackPower,
					p.monster.depensivePower)
				fmt.Fprintf(writer.Newline(), "\n \n")
				if p.checkHealthPoint() {
					writer.Stop()
					break
				}
			} else {
				// 몬스터의 공격 회피
				fmt.Fprintf(writer, "%s \n", Bold(Red("몬스터가 공격하였습니다. ")))
				fmt.Fprintf(writer.Newline(), "%s \n", Bold(Blue("몬스터의 공격을 회피하였습니다.")))
				fmt.Fprintf(writer.Newline(), " \n")
				fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
					p.player.level,
					p.player.healthPoint, p.player.maxHealthPoint,
					p.player.magicPoint, p.player.maxMagicPoint,
					p.player.battleAttackPower,
					p.player.battleDefensivePower)
				fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
					p.monster.healthPoint, p.monster.maxHealthPoint,
					p.monster.attackPower,
					p.monster.depensivePower)
				fmt.Fprintf(writer.Newline(), "\n \n")
			}
			// 케릭터 공격
		} else if fightTime == (p.player.battleAttackSpeed * characterAttack) {
			characterAttack++
			if counterAttack > p.monster.counterAttack {
				p.monster.MonsterAttacked(playerDamage)
				fmt.Fprintf(writer, "%s \n", Bold(Blue("몬스터를 공격하였습니다. ")))
				fmt.Fprintf(writer.Newline(), "%d의 데미지를 입혔습니다. \n", Bold(Blue(playerDamage)))
				fmt.Fprintf(writer.Newline(), " \n")
				fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
					p.player.level,
					p.player.healthPoint, p.player.maxHealthPoint,
					p.player.magicPoint, p.player.maxMagicPoint,
					p.player.battleAttackPower,
					p.player.battleDefensivePower)
				fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
					p.monster.healthPoint, p.monster.maxHealthPoint,
					p.monster.attackPower,
					p.monster.depensivePower)
				fmt.Fprintf(writer.Newline(), "\n \n")
				if p.checkHealthPoint() {
					writer.Stop()
					break
				}
			} else {
				if avoidAttack > p.player.battleAvoid {
					// 몬스터의 반격
					monsterDamage = int(float64(monsterDamage) * 0.7)
					p.player.CharacterAttacked(monsterDamage)
					fmt.Fprintf(writer, "%s \n", Bold(Blue("몬스터를 공격하였습니다. ")))
					fmt.Fprintf(writer.Newline(), "%s \n", Bold(Red("몬스터의 반격!! ")))
					fmt.Fprintf(writer.Newline(), "%d의 데미지를 입었습니다. \n", Bold(Red(monsterDamage)))
					fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
						p.player.level,
						p.player.healthPoint, p.player.maxHealthPoint,
						p.player.magicPoint, p.player.maxMagicPoint,
						p.player.battleAttackPower,
						p.player.battleDefensivePower)
					fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
						p.monster.healthPoint, p.monster.maxHealthPoint,
						p.monster.attackPower,
						p.monster.depensivePower)
					fmt.Fprintf(writer.Newline(), "\n \n")
					if p.checkHealthPoint() {
						writer.Stop()
						break
					}
				} else {
					// 몬스터의 반격 회피
					fmt.Fprintf(writer, "%s \n", Bold(Blue("몬스터를 공격하였습니다. ")))
					fmt.Fprintf(writer.Newline(), "%s \n", Bold(Red("몬스터의 반격!! ")))
					fmt.Fprintf(writer.Newline(), "%s \n", Bold(Blue("몬스터의 공격을 회피하였습니다.")))
					fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
						p.player.level,
						p.player.healthPoint, p.player.maxHealthPoint,
						p.player.magicPoint, p.player.maxMagicPoint,
						p.player.battleAttackPower,
						p.player.battleDefensivePower)
					fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
						p.monster.healthPoint, p.monster.maxHealthPoint,
						p.monster.attackPower,
						p.monster.depensivePower)
					fmt.Fprintf(writer.Newline(), "\n \n")
				}
			}
			// 몬스터 공격
		} else if fightTime == (p.monster.attackSpeed * monsterAttack) {
			monsterAttack++
			if avoidAttack > p.player.battleAvoid {
				p.player.CharacterAttacked(monsterDamage)
				fmt.Fprintf(writer, "%s \n", Bold(Red("몬스터가 공격하였습니다. ")))
				fmt.Fprintf(writer.Newline(), "%d의 데미지를 입었습니다. \n", Bold(Red(monsterDamage)))
				fmt.Fprintf(writer.Newline(), " \n")
				fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
					p.player.level,
					p.player.healthPoint, p.player.maxHealthPoint,
					p.player.magicPoint, p.player.maxMagicPoint,
					p.player.battleAttackPower,
					p.player.battleDefensivePower)
				fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
					p.monster.healthPoint, p.monster.maxHealthPoint,
					p.monster.attackPower,
					p.monster.depensivePower)
				fmt.Fprintf(writer.Newline(), "\n \n")
				if p.checkHealthPoint() {
					writer.Stop()
					break
				}
			} else {
				// 몬스터의 공격 회피
				fmt.Fprintf(writer, "%s \n", Bold(Red("몬스터가 공격하였습니다. ")))
				fmt.Fprintf(writer.Newline(), "%s \n", Bold(Blue("몬스터의 공격을 회피하였습니다.")))
				fmt.Fprintf(writer.Newline(), " \n")
				fmt.Fprintf(writer.Newline(), "Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n",
					p.player.level,
					p.player.healthPoint, p.player.maxHealthPoint,
					p.player.magicPoint, p.player.maxMagicPoint,
					p.player.battleAttackPower,
					p.player.battleDefensivePower)
				fmt.Fprintf(writer.Newline(), "Monster HP: %d / %d ATK: %d DEF: %d\n",
					p.monster.healthPoint, p.monster.maxHealthPoint,
					p.monster.attackPower,
					p.monster.depensivePower)
				fmt.Fprintf(writer.Newline(), "\n \n")
			}
		}
	}
	// 스킬사용 적용값 되돌리기
	p.player.battleAttackPower = p.player.weaponAttackPower
	p.player.battleAttackSpeed = p.player.weaponAttackSpeed
	p.player.battleDefensivePower = p.player.basicDefensivePower
	p.player.battleAvoid = p.player.basicAvoid
}

//스킬 사용
func (p *Process) useingSkills() {
	selection := "go"
	fmt.Printf("스킬을 사용하시려면 Enter를 누르세요")
	fmt.Scanln()

	fmt.Printf("사용하실 스킬 번호를 입력해 주세요 ")
	p.player.PrintSkills()
	fmt.Scanln(&selection)
	if p.player.level < 99 {
		userSelection := CheckSelection(selection, 3)
		p.player.ApplySkill(userSelection)
		fmt.Println("스킬이 적용되었습니다.")
	} else {
		userSelection := CheckSelection(selection, 4)
		p.player.ApplySkill(userSelection)
		fmt.Println("스킬이 적용되었습니다.")
		fmt.Printf("\n \n")

	}
}

// 케릭터 혹은 몬스터가 죽었는지 확인하는 함수
func (p *Process) checkHealthPoint() bool {
	if p.player.healthPoint <= 0 || p.monster.healthPoint <= 0 {
		if p.player.healthPoint <= 0 {
			return true
		} else if p.monster.healthPoint <= 0 {
			return true
		}
	}
	return false
}
