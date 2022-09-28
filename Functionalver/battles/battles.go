package battles

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nyhwbh/bigingame/character/skills"
	"github.com/nyhwbh/bigingame/character/status"
	"github.com/nyhwbh/bigingame/character/weapons"
	"github.com/nyhwbh/bigingame/monster"
)

// 전투 선택시 실행
func FightWithMonster(c *status.CharacterStatus, gameSelection *int, species int, weapon int) {
	newMonster := monster.Monsters{}
	//몬스터 생성 및 스테이터스 출력
	monster.CreateMonster(&newMonster)
	monster.PrintMonsterStatus(&newMonster)
	//몬스터와 전투 시작
	battleCharacterMonset(c, &newMonster, gameSelection, species, weapon)
}

// 회복선택시 안내메세지
func TakeRest(c *status.CharacterStatus) {
	fmt.Println("HP와 MP가 모두 회복되었습니다.")
	status.CharacterRecovery(c)
}

// 몬스터와 전투
func battleCharacterMonset(c *status.CharacterStatus, m *monster.Monsters, gameSelection *int, species int, weapon int) {
	check_hp := false

	// 캐릭터의 체력, 레벨, 공격력, 공격속도, 방어력, 추가회피율 가져오기
	character_hp := status.CallCharacterHp(c)
	level := status.Level(c)
	character_ATK := status.CallCharacterATK(c)
	character_ATKS := status.CallCharacterATKS(c)
	character_DEF := status.CallCharacterDEF(c)
	character_Avoid := 0

	// 몬스터의 체력, 공격력, 공격속도, 방어력 가져오기
	monster_hp := monster.CallMonsterHp(m)
	monster_ATK := monster.CallMonsterATK(m)
	monster_ATKS := 1200
	monster_DEF := monster.CallMonsterDEF(m)

	// 선택한 캐릭터 무기적용
	weapons.ApplyWeapons(species, weapon, &character_ATK, &character_ATKS)

	// 스킬 사용
	skills.ApplySkills(c, level, species, &character_hp, &character_ATK, &character_ATKS, &character_Avoid, &character_DEF, &monster_ATKS)

	// 자동 전투
	autoBattles(character_ATKS, monster_ATKS, c, m, character_ATK, character_DEF, character_Avoid, &character_hp, monster_ATK, monster_DEF, &monster_hp, gameSelection, &check_hp)

	// 전투종료시 함수종료
	for {
		if check_hp == true {
			break
		}
	}
}

// 캐릭터가 몬스터를 공격
func charcterAttack(c *status.CharacterStatus, m *monster.Monsters, character_ATK int, character_DEF int, character_hp *int, monster_ATK int, character_Avoid int, monster_DEF int, monster_hp *int) int {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	counterattack := random.Intn(100)
	avoidingrate := random.Intn(100)
	// 캐릭터 공격력
	totalDamage := character_ATK - monster_DEF
	if counterattack > 30 {
		//캐릭터 공격
		fmt.Println("캐릭터가 공격했습니다.")
		monster.MonsterAttacked(m, totalDamage)
		monster.PrintMonsterStatus(m)
		fmt.Println(" ")
		*monster_hp = monster.CallMonsterHp(m)
		return *monster_hp
	} else {
		// 몬스터 반격
		fmt.Println("캐릭터가 공격했습니다.")
		fmt.Println("몬스터 반격!")
		// 몬스터의 반격중 캐릭터가 푀피
		if avoidingrate > 30+character_Avoid {
			status.CharacterAttacked(c, (int(float64(monster_ATK)*0.7) - character_DEF))
			status.PrintCharacterStatusInbattle(c, *character_hp, character_ATK, character_DEF)
			fmt.Println(" ")
			*character_hp = status.CallCharacterHp(c)
			return *character_hp
		} else {
			//몬스터의 공격 회피
			fmt.Println("몬스터의 반격을 회피하였습니다.")
			fmt.Println(" ")
		}
		return *character_hp
	}
}

// 몬스터가 캐릭터를 공격
func monsterAttack(c *status.CharacterStatus, m *monster.Monsters, character_ATK int, character_DEF int, character_hp *int, monster_ATK int, character_Avoid int) int {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	avoidingrate := random.Intn(100)
	totalDamage := monster_ATK - character_DEF
	if avoidingrate > 30+character_Avoid {
		//몬스터 공격
		fmt.Println("몬스터가 공격했습니다.")
		status.CharacterAttacked(c, totalDamage)
		status.PrintCharacterStatusInbattle(c, *character_hp, character_ATK, character_DEF)
		fmt.Println(" ")
		return *character_hp
	} else {
		//몬스터의 공격 회피
		fmt.Println("몬스터가 공격했습니다.")
		fmt.Println("몬스터의 공격을 회피하였습니다.")
		fmt.Println(" ")
	}
	return *character_hp
}

// 케릭터 혹은 몬스터가 죽었는지 확인하는 함수
func checkHps(c *status.CharacterStatus, c_hp int, m_hp int, gameSelection *int) bool {
	if c_hp <= 0 || m_hp <= 0 {
		if c_hp <= 0 {
			fmt.Println("캐릭터가 사망하였습니다.")
			fmt.Println("Game over")
			*gameSelection = 9
			return true
		} else if m_hp <= 0 {
			fmt.Println("전투에서 승리하였습니다. 캐릭터 레벨이 상승합니다.")
			status.Levelup(c)
			return true
		}
	}
	return false
}

// 캐릭터와 몬스터의 공격속도에 따른 딜레이 부여
func autoBattles(character_ATKS int, monster_ATKS int, c *status.CharacterStatus, m *monster.Monsters, character_ATK int, character_DEF int, character_Avoid int, character_hp *int, monster_ATK int, monster_DEF int, monster_hp *int, gameSelection *int, check_hp *bool) {
	i, j := 1, 1
	attackSpeed := 0
	check_Hp := false
	for {
		time.Sleep(time.Millisecond * 10)
		attackSpeed += 10
		// 캐릭터와 몬스터가 동시에 공격하는 경우
		if attackSpeed == (character_ATKS*i) && attackSpeed == (monster_ATKS*j) {
			charcterAttack(c, m, character_ATK, character_DEF, character_hp, monster_ATK, character_Avoid, monster_DEF, monster_hp)
			i++
			// 해당 공격으로 캐릭터 혹은 몬스터가 죽었는지 확인
			check_Hp = checkHps(c, *character_hp, *monster_hp, gameSelection)
			if check_Hp == true {
				*check_hp = check_Hp
				break
			}
			monsterAttack(c, m, character_ATK, character_DEF, character_hp, monster_ATK, character_Avoid)
			j++
			//해당 공격으로 캐릭터가 죽었는지 확인
			check_Hp = checkHps(c, *character_hp, *monster_hp, gameSelection)
			if check_Hp == true {
				*check_hp = check_Hp
				break
			}
			// 캐릭터가 몬스터를 공격하는 경우
		} else if attackSpeed == (character_ATKS * i) {
			charcterAttack(c, m, character_ATK, character_DEF, character_hp, monster_ATK, character_Avoid, monster_DEF, monster_hp)
			i++
			check_Hp = checkHps(c, *character_hp, *monster_hp, gameSelection)
			if check_Hp == true {
				*check_hp = check_Hp
				break
			}
			// 몬스터가 케릭터를 공격하는 경우
		} else if attackSpeed == (monster_ATKS * j) {
			monsterAttack(c, m, character_ATK, character_DEF, character_hp, monster_ATK, character_Avoid)
			j++
			check_Hp = checkHps(c, *character_hp, *monster_hp, gameSelection)
			if check_Hp == true {
				*check_hp = check_Hp
				break
			}
		}
	}
}
