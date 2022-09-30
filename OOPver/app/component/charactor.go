package component

import (
	"fmt"
	"math/rand"
	"time"
)

//캐맄터 스테이터스
type CharacterStatus struct {
	level              int
	healthPoint        int
	maxHealthPoint     int
	magicPoint         int
	maxMagicPoint      int
	basicAttackPoint   int
	battleAttackPoint  int
	basicAttackSpeed   int
	battleAttackSpeed  int
	basicDefencePoint  int
	battleDefencePoint int
	basicAvoid         int
	battleAvoid        int
	species            int
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
	c.basicAttackPoint = random.Intn(10) + 20
	c.battleAttackPoint = c.basicAttackPoint
	c.basicDefencePoint = random.Intn(10) + 10
	c.battleDefencePoint = c.basicDefencePoint
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
			c.battleAttackPoint = c.basicAttackPoint
		case 2:
			c.battleAttackPoint = int(float64(c.basicAttackPoint) * 1.05)
		case 3:
			c.battleAttackPoint = int(float64(c.basicAttackPoint) * 1.1)
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
			c.battleAttackPoint = c.basicAttackPoint
			c.battleAttackSpeed = c.basicAttackSpeed
		case 2:
			c.battleAttackPoint = int(float64(c.battleAttackPoint) * 1.10)
			c.battleAttackSpeed = int(float64(c.basicAttackSpeed) * 0.95)
		case 3:
			c.battleAttackPoint = int(float64(c.battleAttackPoint) * 1.2)
			c.battleAttackSpeed = int(float64(c.basicAttackSpeed) * 0.9)
		}
	}
}

func (c *CharacterStatus) PrintCharterStatus() {
	fmt.Printf("Lv: %d HP: %d / %d MP: %d / %d ATK: %d DEF: %d ATKS: %d", c.level, c.healthPoint, c.maxHealthPoint, c.magicPoint, c.maxMagicPoint, c.battleAttackPoint, c.battleDefencePoint, c.battleAttackSpeed)
}
