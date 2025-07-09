package main

import (
		"fmt"
		"math/rand"
		"bufio"
		"os"
		"time"
)
type Player struct {
	name string
	position int
	isNaitei int //内定の有無
	naiteisaki []string //内定先の会社名
}
func newPlayer(name string, position int, isNaitei int, naiteisaki []string) *Player {
	player := new(Player)
	player.name = name
	player.position = position
	player.isNaitei = isNaitei
	player.naiteisaki = make([]string, 0) // 内定先の会社名を格納するスライスを初期化
	return player
	
}
func grade(i int) int{
	if i<6{
		return  3
	}else {
		return  4
	}
}
func month(i int) int {
	if i<4 {
		return 2*i+5
	}else if i<10{
		return 2*i -7
	}else{
		return 2*i-19
	}
}
func event(player *Player, i int) string{
	return "イベントはまだ実装されていません。"	
}



func main(){
	fmt.Println("就活始めよう！")
	// var banmen [30]int
	var name string
	fmt.Print("名前を入力してください: ")
	fmt.Scan(&name)
	var player1 = newPlayer(name, 0,0, nil) // 内定の有無と内定先の会社名は初期化
	fmt.Println(player1.name, "の位置は", player1.position)
	for i := 0; i < 12 ; i++ {
		fmt.Println("現在", grade(i),"年生の",month(i),"月")//何年生の何カ月
		if player1.position < 30 {
			//サイコロを振る
			fmt.Println("サイコロを振ります。エンターを押してください")
			bufio.NewReader(os.Stdin).ReadBytes('\n') // エンターキーを待つ

			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(6)+1
			//positionの更新
			player1.position += randomNumber
			if player1.position > 30 {
				player1.position = 30
			}
			//現状の表示
			fmt.Println("サイコロの目は", randomNumber)
			fmt.Println(player1.name, "の位置は", player1.position,"\n")
			//ここで、マス目ごとのイベント
			fmt.Println(event(player1, player1.position))


			//残りのマスの表示
			fmt.Println("残りのマスは",30-player1.position,"\n")
			if player1.position == 30 {
				fmt.Println("就活終了！就職先を選ぼう！")
			}
		}else{
			fmt.Println("ボーナスステージ")
			// fmt.Println("サイコロを振ります。エンターを押してください")
			// bufio.NewReader(os.Stdin).ReadBytes('\n') // エンターキーを待つ

			// rand.Seed(time.Now().UnixNano())
			// randomNumber := rand.Intn(6)+1
		}
	}
	fmt.Println("社会人頑張れ！")
}