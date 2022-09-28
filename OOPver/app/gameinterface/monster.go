package gameinterface

import (
	"math/rand"
	"time"
)

type Monster interface {
	CreateMonster()
	PlayerAttack()
}
type MonsterStatus struct {
	hp         int
	mp         int
	basicAtk   int
	basicAtks  int
	basicDef   int
	basicAvoid int
}

func (m *MonsterStatus) CreateMonster() {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	m.basicAvoid = 30
	m.hp = random.Intn(50) + 50
	m.mp = random.Intn(50) + 50
	m.basicAtk = random.Intn(10) + 10
	m.basicDef = random.Intn(10) + 5
	m.basicAtks = 800
}

func (m *MonsterStatus) PlayerAttack(p *PlayerStatus) {
	m.hp -= (p.battleAtk - m.basicDef)
}

func (m *MonsterStatus) CounterAttack(p *PlayerStatus) {
	p.hp -= (int(float64(m.basicAtk)*0.7) - p.battleDef)
}
