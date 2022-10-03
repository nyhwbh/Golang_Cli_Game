package main

import (
	"cligame/app"
	"fmt"

	. "github.com/logrusorgru/aurora"
)

// cli-game OOPver with TDD
func main() {
	fmt.Println(Bold("게임을 실행합니다."))
	fmt.Println("")

	//게임 실행
	app.RunGame()
}
