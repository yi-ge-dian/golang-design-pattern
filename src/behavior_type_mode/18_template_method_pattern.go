package behavior_type_mode

import "fmt"

// Game 定义一个抽象的游戏接口
type Game interface {
	Initialize()
	StartPlay()
	EndPlay()
}

// GameBase 实现一个游戏抽象基类,用于实现模板方法
type GameBase struct {
	Game
}

func (g *GameBase) Play() {
	g.Initialize()
	g.StartPlay()
	g.EndPlay()
}

// BasketballGame 实现具体的游戏类:篮球
type BasketballGame struct {
	GameBase
}

func (b *BasketballGame) Initialize() {
	fmt.Println("Basketball game initialized.")
}
func (b *BasketballGame) StartPlay() {
	fmt.Println("Basketball game started.")
}
func (b *BasketballGame) EndPlay() {
	fmt.Println("Basketball game ended.")
}

// FootballGame 实现具体的游戏类:足球
type FootballGame struct {
	GameBase
}

func (f *FootballGame) Initialize() {
	fmt.Println("Football game initialized.")
}
func (f *FootballGame) StartPlay() {
	fmt.Println("Football game started.")
}
func (f *FootballGame) EndPlay() {
	fmt.Println("Football game ended.")
}

func CallTemplateMethodPattern() {
	basketballGame := &BasketballGame{}
	basketballGame.Game = basketballGame
	basketballGame.GameBase.Play()
	fmt.Println()
	footballGame := &FootballGame{}
	footballGame.Game = footballGame
	footballGame.GameBase.Play()
}
