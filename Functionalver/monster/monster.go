package monster

import (
	"fmt"
	"math/rand"
	"time"
)

// 몬스터 스테이터스
type Monsters struct {
	healthpoint    int
	maxhp          int
	strikingpower  int
	attackspeed    float64
	depensivepower int
	counterattack  int
}

//몬스터 생성
func CreateMonster(m *Monsters) {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	m.maxhp = random.Intn(50) + 50
	m.healthpoint = m.maxhp
	m.strikingpower = random.Intn(10) + 20
	m.attackspeed = 1000
	m.depensivepower = random.Intn(10) + 10
	m.counterattack = 30
}

// 몬스터 상태창 출력
func PrintMonsterStatus(m *Monsters) {
	fmt.Printf("[ Monster Status  HP : %d/%d ATK : %d DEF : %d ]\n", m.healthpoint, m.maxhp, m.strikingpower, m.depensivepower)
}

// 몬스터 체력
func CallMonsterHp(m *Monsters) int {
	return m.healthpoint
}

// 몬스터 공격력
func CallMonsterATK(m *Monsters) int {
	return m.strikingpower
}

// 몬스터 방어력
func CallMonsterDEF(m *Monsters) int {
	return m.depensivepower
}

// 몬스터가 케릭터에게 공격 받음
func MonsterAttacked(m *Monsters, damage int) {
	m.healthpoint = m.healthpoint - damage
}
