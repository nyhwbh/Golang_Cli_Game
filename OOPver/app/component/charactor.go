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
	battleAttackPower    int
	basicAttackSpeed     int
	battleAttackSpeed    int
	basicDefensivePower  int
	battleDefensivePower int
	basicAvoid           int
	battleAvoid          int
	species              int
	characterSkills      []string
}

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
	c.battleAttackPower = c.basicAttackPower
	c.basicDefensivePower = random.Intn(10) + 10
	c.battleDefensivePower = c.basicDefensivePower
	c.basicAvoid = 30
	c.battleAvoid = c.basicAvoid
	c.basicAttackSpeed = 1000
	c.battleAttackSpeed = c.basicAttackSpeed
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
			c.battleAttackPower = c.basicAttackPower
		case 2:
			c.battleAttackPower = int(float64(c.basicAttackPower) * 1.05)
		case 3:
			c.battleAttackPower = int(float64(c.basicAttackPower) * 1.1)
		}
	case 2:
		switch selection {
		case 1:
			c.battleAttackSpeed = c.basicAttackSpeed
		case 2:
			c.battleAttackSpeed = int(float64(c.basicAttackSpeed) * 1.05)
		case 3:
			c.battleAttackSpeed = int(float64(c.basicAttackSpeed) * 1.1)
		}
	case 3:
		switch selection {
		case 1:
			c.battleAttackPower = c.basicAttackPower
			c.battleAttackSpeed = c.basicAttackSpeed
		case 2:
			c.battleAttackPower = int(float64(c.basicAttackPower) * 1.10)
			c.battleAttackSpeed = int(float64(c.basicAttackSpeed) * 0.95)
		case 3:
			c.battleAttackPower = int(float64(c.basicAttackPower) * 1.2)
			c.battleAttackSpeed = int(float64(c.basicAttackSpeed) * 0.9)
		}
	}
}

// 스킬 생성
func (c *CharacterStatus) CreateSkills() {
	characterSkills := []string{"Heal", "steam"}
	switch c.species {
	case 1:
		characterSkills = append(characterSkills, "Guard", "invincible")
	case 2:
		characterSkills = append(characterSkills, "Elusion", "Rapid")
	case 3:
		characterSkills = append(characterSkills, "Anger", "Frenzy")
	}
	c.characterSkills = characterSkills
}

//케릭터 상태 출력
func (c *CharacterStatus) PrintCharterStatus() {
	fmt.Printf("Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d\n", c.level, c.healthPoint, c.maxHealthPoint, c.magicPoint, c.maxMagicPoint, c.battleAttackPower, c.battleDefensivePower)
}
