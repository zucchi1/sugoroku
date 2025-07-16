package main

import (
		"fmt"
		"math/rand"
		"bufio"
		"os"
		"time"
)
var N = 40 //マスの数
type Player struct {
	name string
	position int
	comSkill int //コミュニケーションスキル
	techSkill int //技術スキル
	money int //所持金
	naiteisaki []string //内定先の会社名
}
func newPlayer(name string, position int, comSkill int, techSkill int , money int, naiteisaki []string) *Player {
	player := new(Player)
	player.name = name
	player.position = position
	player.comSkill = comSkill // コミュニケーションスキルを初期化
	player.techSkill = techSkill // 技術スキルを初期化
	player.money = money // 所持金を初期化
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
func skillup(player *Player,comSkill int, techSkill int) {
	// コミュニケーションスキルを上げる
	player.comSkill += comSkill
	// 技術スキルを上げる
	player.techSkill += techSkill
	fmt.Println(player.name, "のコミュニケーションスキルは", player.comSkill, "、技術スキルは", player.techSkill, "になりました。")
}
func getNaitei(player *Player, company string, rate int)  {
	//内定先のレートに応じて内定獲得の有無を変える
	
	// 内定先に会社名を追加
	player.naiteisaki = append(player.naiteisaki, company)
}

func defineEvent(player *Player, i int) string{
	switch i {
	case 1:
		fmt.Println("企業研究")
		skillup(player,1,1)
		return "1マス目: 企業研究を始めよう！"
	case 2:
		skillup(player,1,1)
		return "2マス目: エントリーシートを書こう！"

	case 3:
		skillup(player,1,1)
		return "3マス目: 面接の練習をしよう！"
	case 4:
		skillup(player,1,1)
		return "4マス目: グループディスカッションの練習"
	case 5:
		skillup(player,1,1)
		return "5マス目: OB訪問をしよう！"
	case 6:
		skillup(player,1,1)
		return "6マス目: インターンシップに参加しよう！"
	default:
		return "イベントはまだ実装されていません。"	
	}
}



func main(){
	fmt.Println("就活始めよう！")
	var name string
	fmt.Print("名前を入力してください: ")
	fmt.Scan(&name)
	var player1 = newPlayer(name, 0,0,0,10000, nil) 
	fmt.Println(player1.name, "の位置は", player1.position)
	for i := 0; i < 12 ; i++ {
		fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
		fmt.Println("現在", grade(i),"年生の",month(i),"月")//何年生の何カ月
		fmt.Println("コミュニケーションスキル:",player1.comSkill, "技術スキル:", player1.techSkill,"所持金:", player1.money)
		if player1.position < N {
			//６マス先のイベント内容を表示する処理
			//サイコロを振る
			fmt.Println("サイコロを振ります。エンターを押してください")
			bufio.NewReader(os.Stdin).ReadBytes('\n') // エンターキーを待つ

			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(6)+1
			//positionの更新
			player1.position += randomNumber
			if player1.position > N {
				player1.position = N
			}
			//現状の表示
			fmt.Println("サイコロの目は", randomNumber)
			fmt.Println(player1.name, "の位置は", player1.position,"\n")
			//ここで、マス目ごとのイベント
			fmt.Println(defineEvent(player1, player1.position))


			//残りのマスの表示
			fmt.Println("残りのマスは",N-player1.position,"\n")
			if player1.position == N {
				fmt.Println("就活終了！就職先を選ぼう！")
				//内定先の会社名を表示
			}
		}else{
			fmt.Println("ボーナスステージ")
			fmt.Println("サイコロを振ります。エンターを押してください")
			bufio.NewReader(os.Stdin).ReadBytes('\n') // エンターキーを待つ

			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(6)+1
			//現状の表示
			fmt.Println("サイコロの目は", randomNumber)

		}
	}
	fmt.Println("ゲーム終了")
	fmt.Println("最終結果")
	fmt.Println("内定先の会社名:", player1.naiteisaki)
	fmt.Println("コミュニケーションスキル:",player1.comSkill, "技術スキル:", player1.techSkill,"所持金:", player1.money)
}