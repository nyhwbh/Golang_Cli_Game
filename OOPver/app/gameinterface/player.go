package gameinterface

import (
	"math/rand"
	"time"
)

type Player interface {
	CreatePlayer()
	PlayerSpeciesSetUp()
	// 레벨업
	PlayerLevelUp()
	StatusLevelUpHuman()
	StatusLevelUpElf()
	StatusLevelUpOak()
	PlayerRecovery()
	// 전투 - 공격, 스킬 사용
	MonsterAttacks(int)
	UsingMagicPoint()
	Heal()
	Steam()
	Guard()
	Invincible()
	Elusion()
	Rapid()
	Anger()
	Frenzy()
	// 무기선택
	Fist()
	ShortSword()
	LongSword()
	ShortBow()
	LongBow()
	ShortAxe()
	IronHammer()
}

/***********************************************
************** 		status		****************
***********************************************/
type PlayerStatus struct {
	level       int
	hp          int
	maxhp       int
	mp          int
	maxmp       int
	basicAtk    int
	battleAtk   int
	basicAtks   int
	battleAtks  int
	basicDef    int
	battleDef   int
	basicAvoid  int
	battleAvoid int
	species     int
	weapons     []string
	skills      []string
}

// 캐릭터 생성
func (p *PlayerStatus) CreatePlayer() {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	p.level = 1
	p.basicAvoid = 30
	p.battleAvoid = p.basicAvoid
	p.maxhp = random.Intn(50) + 50
	p.hp = p.maxhp
	p.maxmp = random.Intn(50) + 50
	p.mp = p.maxmp
	p.basicAtk = random.Intn(10) + 10
	p.battleAtk = p.basicAtk
	p.basicDef = random.Intn(10) + 5
	p.battleDef = p.basicDef
	p.basicAtks = 1000
	p.battleAtks = p.basicAtks
}

// 캐릭터 종족 선택
func (p *PlayerStatus) PlayerSpeciesSetUp(species int) {
	p.species = species
	switch species {
	case 1:
		p.weapons = []string{"Fist", "Short Sword", "Long Sword"}
		p.skills = []string{"Heal", "Steam", "Guard", "Invincible"}
	case 2:
		p.weapons = []string{"Fist", "Short Bow", "Long Bow"}
		p.skills = []string{"Heal", "Steam", "Elusion", "Rapid"}
	case 3:
		p.weapons = []string{"Fist", "Short Axe", "Iron Hammer"}
		p.skills = []string{"Heal", "Steam", "Anger", "Frenzy"}
	}
}

// 캐릭터 레벨업
func (p *PlayerStatus) PlayerLevelUp(species int) {
	p.level += 1

	// 최대체력 증가
	if p.maxhp == 150 {
	} else {
		if p.maxhp+3 > 150 {
			p.maxhp = 150
		} else {
			p.maxhp += 3
		}
	}

	// 최대 마력 증가
	if p.maxmp == 150 {
	} else {
		if p.maxmp+3 > 150 {
			p.maxmp = 150
		} else {
			p.maxmp += 3
		}
	}
	// 종족별 능력치 상승
	if p.level%5 == 0 {
		switch species {
		case 1: // Human
			p.StatusLevelUpHuman()
			return
		case 2: // Elf
			p.StatusLevelUpElf()
			return
		case 3: // Oak
			p.StatusLevelUpOak()
			return
		}
	}
}

// 종족별 레벨업 능력치
// human
func (p *PlayerStatus) StatusLevelUpHuman() {
	p.basicAtk += 1
	p.basicDef += 1
}

// elf
func (p *PlayerStatus) StatusLevelUpElf() {
	p.basicAtk += 1
	p.basicAtks += 100
}

// oak
func (p *PlayerStatus) StatusLevelUpOak() {
	p.basicAtk += 2
}

// 캐릭터 휴식 -> 체력 및 마나 회복
func (p *PlayerStatus) PlayerRecovery() {
	p.hp = p.maxhp
	p.mp = p.maxmp
}

// 피해를 입음
func (p *PlayerStatus) MonsterAttacks(m *MonsterStatus) {
	p.hp -= (m.basicAtk - p.battleDef)
}

/***********************************************
************** 		skills		****************
***********************************************/
// 마나소모
func (p *PlayerStatus) UsingMagicPoint() {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	p.mp -= (random.Intn(15) + 10)
}

// 공통스킬
func (p *PlayerStatus) Heal() {
	p.hp += 30
	p.UsingMagicPoint()
}

func (p *PlayerStatus) Steam() {
	p.battleAtk = int(float64(p.battleAtk) * 0.2)
	p.UsingMagicPoint()
}

// 휴먼스킬
//HumanBasicSkill
func (p *PlayerStatus) Guard() {
	p.basicDef = int(float64(p.basicDef) * 0.3)
	p.UsingMagicPoint()
}

//HumanHyperSkill
func (p *PlayerStatus) Invincible() {
	//10초간 무적
	p.UsingMagicPoint()
}

// 엘프스킬
//ElfbasicSkill
func (p *PlayerStatus) Elusion() {
	p.battleAvoid += 30
	p.UsingMagicPoint()
}

//ElfHyperSkill
func (p *PlayerStatus) Rapid() {
	//1분동안 공격속도 500%상승
	p.UsingMagicPoint()
}

// 오크스킬
//OakBasicSkill
func (p *PlayerStatus) Anger() {
	p.battleAtk += int(float64(p.battleAtk) * 0.5)
	p.battleDef -= int(float64(p.battleDef) * 0.1)
	p.UsingMagicPoint()
}

//OakHyperSkill
func (p *PlayerStatus) Frenzy() {
	//1분동안 공격력 500%상승
	p.UsingMagicPoint()
}

/***********************************************
************** 		weapon		****************
***********************************************/
// 공통무기 - 주먹
func (p *PlayerStatus) Fist() {
	p.battleAtk = p.basicAtk
}

// Human
func (p *PlayerStatus) ShortSword() {
	p.battleAtk = p.basicAtk + int(float64(p.basicAtk)*0.05)
}
func (p *PlayerStatus) LongSword() {
	p.battleAtk = p.basicAtk + int(float64(p.basicAtk)*0.1)
}

// Elf
func (p *PlayerStatus) ShortBow() {
	p.battleAtks = p.basicAtks + int(float64(p.basicAtks)*0.05)
}
func (p *PlayerStatus) LongBow() {
	p.battleAtks = p.basicAtks + int(float64(p.basicAtks)*0.1)
}

//Oak
func (p *PlayerStatus) ShortAxe() {
	p.battleAtk = p.basicAtk + int(float64(p.basicAtk)*0.1)
	p.battleAtks = p.basicAtks - int(float64(p.basicAtks)*0.05)
}
func (p *PlayerStatus) IronHammer() {
	p.battleAtk = p.basicAtk + int(float64(p.basicAtk)*0.2)
	p.battleAtks = p.basicAtks - int(float64(p.basicAtks)*0.1)
}
