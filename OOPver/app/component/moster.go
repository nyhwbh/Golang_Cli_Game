package component

import (
	"fmt"
	"math/rand"
	"time"
)

type Monsters struct {
	healthPoint    int
	maxHealthPoint int
	attackPower    int
	attackSpeed    float64
	depensivePower int
	counterAttack  int
}

// ----------- 기본 설정 함수 ----------- //
// 몬스터 생성
func (m *Monsters) CreateMoster() {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	m.maxHealthPoint = random.Intn(50) + 50
	m.healthPoint = m.maxHealthPoint
	m.attackPower = random.Intn(10) + 20
	m.attackSpeed = 1000
	m.depensivePower = random.Intn(10) + 10
	m.counterAttack = 30
}

// ----------- 전투 관련 함수 ----------- //
// 몬스터가 공격 받음
func (m *Monsters) MonsterAttacked(damage int) {
	m.healthPoint -= damage
}

// ----------- 출력 함수 ----------- //
// 몬스터 상태창 출력
func (m *Monsters) PrintMonsterStatus() {
	fmt.Printf("[ Monster Status  HP : %d/%d ATK : %d DEF : %d ]\n", m.healthPoint, m.maxHealthPoint, m.attackPower, m.depensivePower)
}
