package component

import (
	"fmt"
	"math/rand"
	"time"
)

//캐맄터 스테이터스
type CharacterStatus struct {
	level                int
	healthPoint          int
	maxHealthPoint       int
	magicPoint           int
	maxMagicPoint        int
	basicAttackPower     int
	weaponAttackPower    int
	battleAttackPower    int
	basicAttackSpeed     int
	weaponAttackSpeed    int
	battleAttackSpeed    int
	basicDefensivePower  int
	battleDefensivePower int
	basicAvoid           int
	battleAvoid          int
	species              int
	characterSkills      []string
}

// ----------- 기본 설정 함수 ----------- //
// 케릭터 생성
func (c *CharacterStatus) BasicSetup() {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	c.level = 1
	c.maxHealthPoint = random.Intn(50) + 50
	c.healthPoint = c.maxHealthPoint
	c.maxMagicPoint = random.Intn(50) + 50
	c.magicPoint = c.maxMagicPoint
	c.basicAttackPower = random.Intn(10) + 20
	c.weaponAttackPower = c.basicAttackPower
	c.battleAttackPower = c.weaponAttackPower
	c.basicDefensivePower = random.Intn(10) + 10
	c.battleDefensivePower = c.basicDefensivePower
	c.basicAvoid = 30
	c.battleAvoid = c.basicAvoid
	c.basicAttackSpeed = 1000
	c.weaponAttackSpeed = c.basicAttackSpeed
	c.battleAttackSpeed = c.weaponAttackSpeed
}

// 종족 설정
func (c *CharacterStatus) SetSpecies(selection int) {
	c.species = selection
}

// 무기 변경
func (c *CharacterStatus) SetWeapon(selection int) {
	switch c.species {
	case 1:
		switch selection {
		case 1:
			c.weaponAttackPower = c.basicAttackPower
			c.battleAttackPower = c.weaponAttackPower
		case 2:
			c.weaponAttackPower = int(float64(c.basicAttackPower) * 1.05)
			c.battleAttackPower = c.weaponAttackPower
		case 3:
			c.weaponAttackPower = int(float64(c.basicAttackPower) * 1.1)
			c.battleAttackPower = c.weaponAttackPower
		}
	case 2:
		switch selection {
		case 1:
			c.weaponAttackSpeed = c.basicAttackSpeed
			c.battleAttackSpeed = c.weaponAttackSpeed
		case 2:
			c.weaponAttackSpeed = int(float64(c.basicAttackSpeed) * 0.95)
			c.battleAttackSpeed = c.weaponAttackSpeed
		case 3:
			c.weaponAttackSpeed = int(float64(c.basicAttackSpeed) * 0.9)
			c.battleAttackSpeed = c.weaponAttackSpeed
		}
	case 3:
		switch selection {
		case 1:
			c.weaponAttackPower = c.basicAttackPower
			c.battleAttackPower = c.weaponAttackPower
			c.weaponAttackSpeed = c.basicAttackSpeed
		case 2:
			c.weaponAttackPower = int(float64(c.basicAttackPower) * 1.10)
			c.battleAttackPower = c.weaponAttackPower
			c.weaponAttackSpeed = int(float64(c.basicAttackSpeed) * 1.05)
			c.battleAttackSpeed = c.weaponAttackSpeed
		case 3:
			c.weaponAttackPower = int(float64(c.basicAttackPower) * 1.2)
			c.battleAttackPower = c.weaponAttackPower
			c.weaponAttackSpeed = int(float64(c.basicAttackSpeed) * 1.1)
			c.battleAttackSpeed = c.weaponAttackSpeed
		}
	}
}

// 스킬 생성
func (c *CharacterStatus) CreateSkills() {
	characterSkills := []string{"1.Heal", "2.Steam"}
	switch c.species {
	case 1:
		characterSkills = append(characterSkills, "3.Guard", "4.Invincible")
	case 2:
		characterSkills = append(characterSkills, "3.Elusion", "4.Rapid")
	case 3:
		characterSkills = append(characterSkills, "3.Anger", "4.Frenzy")
	}
	c.characterSkills = characterSkills
}

// ----------- 전투관련 함수 ----------- //
// 케릭터가 공격 받음
func (c *CharacterStatus) CharacterAttacked(damage int) {
	c.healthPoint -= damage
}

// 레벨업
func (c *CharacterStatus) CharacterLevelUp() {
	c.level++
}

// ----------- 스킬 사용 함수 ----------- //
// 스킬 사용
func (c *CharacterStatus) ApplySkill(selection int) {
	mpCheck := c.UsingMagicPoint()
	if mpCheck {
		switch selection {
		case 1:
			c.SkillHeal()
		case 2:
			c.SkillSteam()
		case 3:
			switch c.species {
			case 1:
				c.SkillGuard()
			case 2:
				c.SkillElusion()
			case 3:
				c.SkillAnger()
			}
		case 4:
			switch c.species {
			case 1:
				c.SkillInvincible()
			case 2:
				c.SkillRapid()
			case 3:
				c.SkillFrenzy()
			}
		}
	} else {
		fmt.Println("MP가 부족하여 스킬을 사용할 수 없습니다.")
	}
}

// MP 차감
func (c *CharacterStatus) UsingMagicPoint() bool {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	usedMP := random.Intn(10) + 10
	check_mp := c.magicPoint - usedMP
	if check_mp <= 0 {
		return false
	} else {
		c.magicPoint = check_mp
		return true
	}
}

// Skills
func (c *CharacterStatus) SkillHeal() {
	if (c.healthPoint + 30) >= c.maxHealthPoint {
		c.healthPoint = c.maxHealthPoint
	} else {
		c.healthPoint += 30
	}
}

func (c *CharacterStatus) SkillSteam() {
	c.battleAttackPower = int(float64(c.weaponAttackPower) * 1.2)
}

func (c *CharacterStatus) SkillGuard() {
	c.battleDefensivePower = int(float64(c.basicDefensivePower) * 1.3)
}

func (c *CharacterStatus) SkillInvincible() {
	//10초 동안 무적이 됨
}

func (c *CharacterStatus) SkillElusion() {
	c.battleAvoid = int(float64(c.basicAvoid) * 1.3)
}

func (c *CharacterStatus) SkillRapid() {
	// 1분동안
	c.battleAttackSpeed = int(float64(c.weaponAttackSpeed) / 5)
}

func (c *CharacterStatus) SkillAnger() {
	c.battleAttackPower = int(float64(c.weaponAttackPower) * 1.5)
	c.battleDefensivePower = int(float64(c.basicDefensivePower) * 0.9)
}

func (c *CharacterStatus) SkillFrenzy() {
	// 1분동안
	c.battleAttackPower = int(float64(c.weaponAttackPower) * 5)

}

// ----------- 출력 함수 ----------- //
//케릭터 상태 출력
func (c *CharacterStatus) PrintCharterStatus() {
	fmt.Printf("Player Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n", c.level, c.healthPoint, c.maxHealthPoint, c.magicPoint, c.maxMagicPoint, c.battleAttackPower, c.battleDefensivePower)
}

// 스킬 출력
func (c *CharacterStatus) PrintSkills() {
	if c.level < 99 {
		fmt.Printf("%s", c.characterSkills[0:3])
	} else {
		fmt.Printf("%s", c.characterSkills)
	}
}
