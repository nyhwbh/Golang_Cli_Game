package status

import (
	"fmt"
	"math/rand"
	"time"
)

//캐맄터 스테이터스
type CharacterStatus struct {
	level   int
	hp      int
	maxhp   int
	mp      int
	maxmp   int
	atk     int
	atks    int
	def     int
	avoid   int
	species int
}

//캐릭터 생성
func BasicSetup(c *CharacterStatus) {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	c.level = 1
	c.avoid = 30
	c.maxhp = random.Intn(50) + 50
	c.hp = c.maxhp
	c.maxmp = random.Intn(50) + 50
	c.mp = c.maxmp
	c.atk = random.Intn(10) + 20
	c.def = random.Intn(10) + 10
	c.atks = 1000
}

// 레벨 업
func Levelup(c *CharacterStatus) {
	c.level += 1
}

// 레벨값 가져오기
func Level(c *CharacterStatus) int {
	return c.level
}

//종족값 가져오기
func SetSpecies(c *CharacterStatus, cs int) {
	c.species = cs
}

// 휴식시 체력회복
func CharacterRecovery(c *CharacterStatus) {
	c.hp = c.maxhp
	c.mp = c.maxmp
}

// 캐릭터 체력 가져오기
func CallCharacterHp(c *CharacterStatus) int {
	return c.hp
}

// 캐릭터 최대 체력 가져오기
func CallCharacterMaxHp(c *CharacterStatus) int {
	return c.maxhp
}

// 캐릭터 공격속도 가져오기
func CallCharacterATK(c *CharacterStatus) int {
	return c.atk
}

// 캐릭터 공격력 가져오기
func CallCharacterATKS(c *CharacterStatus) int {
	return c.atks
}

// 캐릭터 방어력 가져오기
func CallCharacterDEF(c *CharacterStatus) int {
	return c.def
}

// 몬스터에게 공격당함
func CharacterAttacked(c *CharacterStatus, damage int) {
	c.hp = c.hp - damage
}

// 스킬 힐 사용
func SkillHeal(c *CharacterStatus) {
	if (c.hp + 30) > c.maxhp {
		c.hp = c.maxhp
		UsingMagicPoint(c)
	} else {
		c.hp += 30
		UsingMagicPoint(c)
	}
}

// 스킬 사용시 마나 감소
func UsingMagicPoint(c *CharacterStatus) bool {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	usedMP := random.Intn(10) + 10
	check_mp := c.mp - usedMP
	if check_mp <= 0 {
		fmt.Println("MP가 부족하여 스킬을 사용할 수 없습니다.")
		return false
	} else {
		c.mp = check_mp
		return true
	}

}

// 상태창 출력
func PrintCharacterStatus(c *CharacterStatus) {
	fmt.Printf("[ Charactor Status  Level : %d HP : %d/%d MP : %d/%d ATK : %d DEF : %d ]\n", c.level, c.hp, c.maxhp, c.mp, c.maxmp, c.atk, c.def)
}

// 전투중 상태창 출력
func PrintCharacterStatusInbattle(c *CharacterStatus, character_hp int, character_ATK int, character_DEF int) {
	fmt.Printf("[ Charactor Status  Level : %d HP : %d/%d MP : %d/%d ATK : %d DEF : %d ]\n", c.level, character_hp, c.maxhp, c.mp, c.maxmp, character_ATK, character_DEF)
}

//기능 확인 함수

// 전투시 체력 마력 30씩 사용한 상황을 가정
func UsingHMp(c *CharacterStatus) {
	c.hp -= 30
	c.mp -= 30
}

// 궁극스킬 추가및 사용 확인을 위한 레벨업
func HiperLevelup(c *CharacterStatus) {
	c.level = 99
}
