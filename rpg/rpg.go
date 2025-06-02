package main

import (
	"fmt"
	"time"
)

const (
	Level   = 1
	Exp     = 0
	HP      = 100
	MP      = 100
	Attack  = 10
	Defense = 5
)

func monster_turn(monsterName string, monsterAttack int, playerName string, playerHP *int) {
	fmt.Println(monsterName + "のターン")
	time.Sleep(1 * time.Second)
	fmt.Println(monsterName + "はこうげきした！")
	time.Sleep(1 * time.Second)
	fmt.Printf("%dのダメージ！\n", monsterAttack)
	*playerHP -= monsterAttack
	fmt.Printf("%sの体力: %d\n", playerName, *playerHP)
	time.Sleep(1 * time.Second)
}

func main() {
	var playerName string
	fmt.Println("プレイヤー名を入力してください")
	fmt.Scan(&playerName)

	playerLevel := Level
	playerExp := Exp
	playerHP := HP
	playerMP := MP
	playerAttack := Attack
	playerDefense := Defense

	fmt.Printf("なまえ　　:\t%s\nたいりょく:\t%5d\nこうげき　:\t%5d\n魔法 :\t%5d\nぼうぎょ　:\t%5d\nレベル　　:\t%5d\nけいけんち:\t%5d\n",
		playerName, playerHP, playerAttack, playerMP, playerDefense, playerLevel, playerExp)

	monsterName := "スライム"
	monsterHP := 50
	monsterAttack := 5

	fmt.Println(monsterName, "があらわれた！")
	fmt.Println(monsterHP, ":", monsterName)
	time.Sleep(1 * time.Second)

	for i := 0; monsterHP > 0; i++ {
		fmt.Println("ターン", i+1)
		fmt.Printf("%sのHP: %d\n", playerName, playerHP)
		fmt.Printf("%sのHP: %d\n", monsterName, monsterHP)
		time.Sleep(1 * time.Second)

		fmt.Println("コマンドを選択してください")
		fmt.Println("1:こうげき")
		fmt.Println("2:にげる")
		fmt.Println("3:どうぐ")
		fmt.Println("4:じゅもん")

		var command int
		fmt.Scan(&command)

		switch command {
		case 1:
			fmt.Println("こうげきを選択しました")
			time.Sleep(1 * time.Second)
			damage := playerAttack
			fmt.Println(monsterName, "に", damage, "のダメージ！")
			monsterHP -= damage

			if monsterHP <= 0 {
				fmt.Println(monsterName, "がいなくなった！")
				fmt.Println(playerName, "は勝った！")
				fmt.Println("おめでとう！")
				break
			}
		case 2:
			fmt.Println("にげだした！")
		case 3:
			fmt.Println("どうぐを使おうとしたが、まだ何も持っていない！")
		case 4:
			fmt.Println("じゅもんを唱えようとしたが、まだ覚えていない！")
		default:
			fmt.Println("不正なコマンドです")
		}

		if monsterHP <= 0 {
			break
		}

		monster_turn(monsterName, monsterAttack, playerName, &playerHP)

		if playerHP <= 0 {
			fmt.Println(playerName, "は負けた！")
			return
		}
	}
}
