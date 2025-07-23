package main

import (
		"fmt"
		"math/rand"
		"math"
		"bufio"
		"os"
		"time"
		"strconv"
	
		"strings"
)
var N = 40 //マスの数
type Player struct {
	name string
	position int
	comSkill int //コミュニケーションスキル
	techSkill int //技術スキル
	money int //所持金
	naiteisaki []string //内定先の会社名
    gyoukaiLevel []int
}
func newPlayer(name string, position int, comSkill int, techSkill int , money int, naiteisaki []string) *Player {
	player := new(Player)
	player.name = name
	player.position = position
	player.comSkill = comSkill // コミュニケーションスキルを初期化
	player.techSkill = techSkill // 技術スキルを初期化
	player.money = money // 所持金を初期化
	player.naiteisaki = make([]string, 0) // 内定先の会社名を格納するスライスを初期化
    player.gyoukaiLevel = make([]int, 23)
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
func getNaitei(player *Player, company string, requiredComSkill int, requiredTechSkill int, successRate int)  {
	fmt.Println(company, "の内定の必要なコミュニケーションスキルは", requiredComSkill, "、技術スキルは", requiredTechSkill, "、成功率は", successRate, "%です。")
	//内定先のレートに応じて内定獲得の有無を変える
	//自分のスキルと会社の必要スキルを比較し、確率を変動
	
	successRate = successRate + int(math.Max(0,float64(player.comSkill - requiredComSkill))) +int(math.Max(0,float64(player.techSkill - requiredTechSkill)))
	if successRate > 100 {
		successRate = 100 // 成功率は100%を超えない
	}
	fmt.Println("最終的な成功率は", successRate, "%です。")
	//　エンターが入力されるまで待つ
	fmt.Println("エンターを押して内定結果を待ちましょう。")
	fmt.Println("1~100のうち、", successRate, "以下の数字が出れば内定獲得！")
	bufio.NewReader(os.Stdin).ReadBytes('\n') // エンターキーを待つ
	
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)+1
	fmt.Println("サイコロの目は", randomNumber)
	if randomNumber > successRate {
		fmt.Println("残念！", company, "の内定はもらえません。")
		return // 内定がもらえなかった場合は終了
	}
	fmt.Println("おめでとう！", company, "の内定を獲得しました！")
	player.naiteisaki = append(player.naiteisaki, company)
}
func getMoney(player *Player, amount int) {
	fmt.Println(player.name, "の所持金が", amount, "円増える。")
	// 所持金を増やす
	player.money += amount
	if player.money < 0 {
		player.money = 0 // 所持金がマイナスにならないようにする
	}
	fmt.Println(player.name, "の所持金は", player.money, "になりました。")
}
func defineNaiteisaki(player *Player) int {
	if len(player.naiteisaki) == 0 {
		fmt.Println("内定先の会社はありません。就活留年ですね。")
		return -1
	}
	fmt.Println("内定先の会社名は以下です:")
	for i := 0; i < len(player.naiteisaki); i++ {
		fmt.Println(i+1, ":", player.naiteisaki[i])
	}
	if len(player.naiteisaki) == 1 {
		fmt.Println("入社先は", player.naiteisaki[0], "です。")
		return 0
	}
	fmt.Println("入社する会社を決めましょう:")
	// 数値の入力を受け付ける
	var choice int
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("1~", len(player.naiteisaki), "の数字を入力してください: ")
		input,err :=reader.ReadString('\n')
		if err != nil {
			fmt.Println("入力エラー:", err)
			continue
		}
		input = strings.TrimSpace(input) // 改行を削除
		choice, err = strconv.Atoi(input) // 数値に変換
		if err != nil || choice < 1 || choice > len(player.naiteisaki) {
			fmt.Println("無効な入力です。もう一度入力してください。")
			continue
		}
		fmt.Println("あなたが選んだ会社は", player.naiteisaki[choice-1], "です。")
		return choice // 有効な入力があればループを抜ける

		break // 有効な入力があればループを抜ける
	}
	
	return -1 // 何も選ばなかった場合は-1を返す

}

func defineEvent(player *Player, i int) {
    var company [23]string
    company[0] = "サイバーエージェント"
    company[1] = "NTT西日本"
    company[2] = "マイクロンメモリジャパン"
    company[3] = "パナソニック"
    company[4] = "トヨタ自動車九州"
    company[5] = "富士通"
    company[6] = "SONY"
    company[7] = "マツダ"
    company[8] = "Amazon"
    company[9] = "広島大学"
    company[10] = "両備システムズ"
    company[11] = "アクセンチュア"
    company[12] = "キーエンス"
    company[13] = "ソフトバンク"
    company[14] = "日立製作所"
    company[15] = "三菱総合研究所"
    company[16] = "トヨタ自動車"
    company[17] = "NTTデータ"
    company[18] = "中国電力"
    company[19] = "野村総合研究所"
    company[20] = "Apple"
    company[21] = "伊藤忠商事"
    company[22] = "Google"
    switch i {
    case 1:
        fmt.Println(i,"マス目：スキルアップイベント「自己分析をしよう」")
        fmt.Println("コミュニケーションスキルが10上がった！")
        skillup(player,10,0)
 
    case 2:
        fmt.Println(i,"マス目：スキルアップイベント「自身の研究を進めよう！」")
        fmt.Println("技術スキルが10上がった！")
        skillup(player,0,10)
 
    case 3:
        fmt.Println(i,"マス目：お金稼ぎイベント「リゾートバイト！」")
        fmt.Println("15000円稼いだ！")
        getMoney(player,15000)
 
    case 4:
        fmt.Println(i,"マス目：スキルアップイベント「企業研究をしよう！」")
        fmt.Println("コミュニケーションスキルが10上がった！")
        fmt.Println("技術スキルが5上がった！")
        skillup(player,10,5)
 
    case 5:
        fmt.Println(i,"マス目：スキルアップイベント「自己分析をしよう！」")
        fmt.Println("コミュニケーションスキルが10上がった！")
        skillup(player,10,0)
 
    case 6:
        fmt.Println(i,"マス目：スキルアップイベント「自身の研究を進めよう！」")
        fmt.Println("研究の調子が良くかなり進んだ！")
        fmt.Println("技術スキルが20上がった！")
        skillup(player,0,20)
 
    case 7:
        fmt.Println(i,"マス目：スキルアップ＋お金稼ぎイベント「報酬ありインターンに参加しよう！」")
        fmt.Println("技術スキルが15上がった！")
        skillup(player,15,0)
        fmt.Println("10000円稼いだ！")
        getMoney(player,10000)
 
    case 8:
        fmt.Println(i,"マス目：スキルアップイベント「企業研究をしよう！」")
        fmt.Println("コミュニケーションスキルが10上がった！")
        fmt.Println("技術スキルが5上がった！")
        skillup(player,10,5)
 
    case 9:
        fmt.Println(i,"マス目：スキルアップイベント「自身の研究を進めよう！」")
        fmt.Println("進捗報告で教授に褒められた！")
        fmt.Println("技術スキルが10上がった！")
        fmt.Println("コミュニケーションスキルが10上がった！")
        skillup(player,10,10)
 
    case 10:
        fmt.Println(i,"マス目：スキルアップイベント「学会発表をしよう！」")
        fmt.Println("技術スキルが10上がった！")
        fmt.Println("コミュニケーションスキルが20上がった！")
        skillup(player,10,20)
 
    case 11:
        fmt.Println(i,"マス目：スキルアップイベント「論文を書こう！」")
        fmt.Println("技術スキルが20上がった！")
        fmt.Println("コミュニケーションスキルが10上がった！")
        skillup(player,20,10)
 
    case 12:
        fmt.Println(i,"マス目：お金稼ぎイベント「接客アルバイトをしよう！」")
        fmt.Println("10000円稼いだ！")
        getMoney(player,10000)
        fmt.Println("コミュニケーションスキルが10上がった！")
        skillup(player,5,0)
 
    case 13:
        fmt.Println(i,"マス目：お金稼ぎイベント「プログラミングアルバイトをしよう！」")
        fmt.Println("10000円稼いだ！")
        getMoney(player,10000)
        fmt.Println("技術スキルが10上がった！")
        skillup(player,0,10)
 
    case 14:
        fmt.Println(i,"マス目：スキルアップイベント「業界研究セミナーに参加しよう！」")
        fmt.Println("どの業界のセミナーに参加するか選んでください。")
        fmt.Println("○○の合格確率が20上がった！")
 
    case 15:
        fmt.Println(i,"マス目：スキルアップイベント「ある会社のインターンに参加しよう！」")
        fmt.Println("どの会社のインターンに参加するか選んでください。")
        fmt.Println("○○の合格確率が50上がった！")
 
    case 16:
		fmt.Println(i,"マス目：スキルアップイベント「インターンに参加しよう（交通費支給なし）！」")
        var input string
        if player.money >= 15000 {
            fmt.Print("インターンに参加しますか？(yes/no)：")
            fmt.Scanln(&input)

            if input == "yes"{
                fmt.Println("かなり内容の濃いインターンだった！")
                fmt.Println("技術スキルが25上がった！")
                skillup(player,0,25)
                fmt.Println("交通費の支給がなかった...")
                fmt.Println("15000円失った...")
                getMoney(player,-15000)
            } else {
                fmt.Println("今回はインターンに参加しなかった。")
            }
        }
        

    case 17:
        fmt.Println(i,"マス目：お金稼ぎイベント「単発バイトをしよう！」")
        fmt.Println("10000円稼いだ！")
        getMoney(player,10000)
 
    case 18:
        fmt.Println(i,"マス目：内定イベント「サイバーエージェント」")
        fmt.Println("合格確率:30%")
        fmt.Println("スキル目安 コミュニケーションスキル:10 技術スキル:10")
        getNaitei(player, "サイバーエージェント", 10, 10, 30)
 
    case 19:
        fmt.Println(i,"マス目：内定イベント「NTT西日本」")
        fmt.Println("合格確率:50%")
        fmt.Println("スキル目安 コミュニケーションスキル:20 技術スキル:5")
        getNaitei(player, "NTT西日本", 20, 5, 50)
    case 20:
        fmt.Println(i,"マス目:内定イベント「マイクロンメモリジャパン」")
        fmt.Println("合格確率:60%")
        fmt.Println("スキル目安 コミュニケーションスキル:10 技術スキル:10")
        getNaitei(player,"マイクロンメモリジャパン",10,10,60)
    case 21:
        fmt.Println(i,"マス目:内定イベント「パナソニック」")
        fmt.Println("合格確率:50%")
        fmt.Println("スキル目安 コミュニケーションスキル:15 技術スキル:15")
        getNaitei(player,"パナソニック",15,15,50)
    case 22:
        fmt.Println(i,"マス目:内定イベント「トヨタ自動車九州」")
        fmt.Println("合格確率:50%")
        fmt.Println("スキル目安 コミュニケーションスキル:15 技術スキル:15")
        getNaitei(player,"トヨタ自動車九州",15,15,50)
    case 23:
        fmt.Println(i,"マス目:内定イベント「富士通」")
        fmt.Println("合格確率:30%")
        fmt.Println("スキル目安 コミュニケーションスキル:17 技術スキル:19")
        getNaitei(player,"富士通",17,19,30)
    case 24:
        fmt.Println(i,"マス目:内定イベント「SONY」")
        fmt.Println("合格確率:20%")
        fmt.Println("スキル目安 コミュニケーションスキル:18 技術スキル:20")
        getNaitei(player,"SONY",18,20,20)
    case 25:
        fmt.Println(i,"マス目:内定イベント「マツダ」")
        fmt.Println("合格確率:60%")
        fmt.Println("スキル目安 コミュニケーションスキル:12 技術スキル:10")
        getNaitei(player,"マツダ",12,10,60)
    case 26:
        fmt.Println(i,"マス目:内定イベント「Amazon」")
        fmt.Println("合格確率:10%")
        fmt.Println("スキル目安 コミュニケーションスキル:18 技術スキル:20")
        getNaitei(player,"Amazon",18,20,10)
    case 27:
        fmt.Println(i,"マス目:内定イベント「広島大学」")
        fmt.Println("合格確率:50%")
        fmt.Println("スキル目安 コミュニケーションスキル:15 技術スキル:14")
        getNaitei(player,"広島大学",15,14,50)
    case 28:
        fmt.Println(i,"マス目:内定イベント「両備システムズ」")
        fmt.Println("合格確率:80%")
        fmt.Println("スキル目安 コミュニケーションスキル:11 技術スキル:8")
        getNaitei(player,"両備システムズ",11,8,80)
    case 29:
        fmt.Println(i,"マス目:内定イベント「アクセンチュア」")
        fmt.Println("合格確率:50%")
        fmt.Println("スキル目安 コミュニケーションスキル:13 技術スキル:16")
        getNaitei(player,"アクセンチュア",13,16,50)
    case 30:
        fmt.Println(i,"マス目:内定イベント「キーエンス」")
        fmt.Println("合格確率:20%")
        fmt.Println("スキル目安 コミュニケーションスキル:23 技術スキル:15")
        getNaitei(player,"キーエンス",23,15,20)
    case 31:
        fmt.Println(i,"マス目:内定イベント「ソフトバンク」")
        fmt.Println("合格確率:60%")
        fmt.Println("スキル目安 コミュニケーションスキル:19 技術スキル:18")
        getNaitei(player,"ソフトバンク",19,18,60)
    case 32:
        fmt.Println(i,"マス目:内定イベント「日立製作所」")
        fmt.Println("合格確率:50%")
        fmt.Println("スキル目安 コミュニケーションスキル:15 技術スキル:16")
        getNaitei(player,"日立製作所",15,16,50)
    case 33:
        fmt.Println(i,"マス目:内定イベント「三菱総合研究所」")
        fmt.Println("合格確率:30%")
        fmt.Println("スキル目安 コミュニケーションスキル:18 技術スキル:18")
        getNaitei(player,"三菱総合研究所",18,18,30)
    case 34:
        fmt.Println(i,"マス目:内定イベント「トヨタ自動車」")
        fmt.Println("合格確率:60%")
        fmt.Println("スキル目安 コミュニケーションスキル:12 技術スキル:14")
        getNaitei(player,"トヨタ自動車",12,14,60)
    case 35:
        fmt.Println(i,"マス目:内定イベント「NTTデータ」")
        fmt.Println("合格確率:20%")
        fmt.Println("スキル目安 コミュニケーションスキル:18 技術スキル:19")
        getNaitei(player,"NTTデータ",18,19,20)
    case 36:
        fmt.Println(i,"マス目:内定イベント「中国電力」")
        fmt.Println("合格確率:60%")
        fmt.Println("スキル目安 コミュニケーションスキル:17 技術スキル:12")
        getNaitei(player,"中国電力",17,12,60)
    case 37:
        fmt.Println(i,"マス目:内定イベント「野村総合研究所」")
        fmt.Println("合格確率:15%")
        fmt.Println("スキル目安 コミュニケーションスキル:19 技術スキル:15")
        getNaitei(player,"野村総合研究所",19,15,15)
    case 38:
        fmt.Println(i,"マス目:内定イベント「Apple」")
        fmt.Println("合格確率:5%")
        fmt.Println("スキル目安 コミュニケーションスキル:17 技術スキル:22")
        getNaitei(player,"Apple",17,22,5)
    case 39:
        fmt.Println(i,"マス目:内定イベント「伊藤忠商事」")
        fmt.Println("合格確率:5%")
        fmt.Println("スキル目安 コミュニケーションスキル:22 技術スキル:17")
        getNaitei(player,"伊藤忠商事",22,17,5)
    case 40:
        fmt.Println(i,"マス目:内定イベント「Google」")
        fmt.Println("合格確率:5%")
        fmt.Println("スキル目安 コミュニケーションスキル:18 技術スキル:21")
        getNaitei(player,"Google",18,21,5)
    default:
        fmt.Println("ゴール！")
    }
}


func banmen(i int,event [47]string){
   fmt.Println("今のマスは",event[i])
    fmt.Println("1を出せば",event[i+1])
    fmt.Println("2を出せば",event[i+2])
    fmt.Println("3を出せば",event[i+3])
    fmt.Println("4を出せば",event[i+4])
    fmt.Println("5を出せば",event[i+5])
    fmt.Println("6を出せば",event[i+6])
}


func main(){
	fmt.Println("就活を始めよう")
    var event [47]string
    event[0] = "スタート地点"
	event[1] = "スキルアップイベント「自己分析をしよう」"
    event[2] = "スキルアップイベント「自身の研究を進めよう！」"
    event[3] = "お金稼ぎイベント「リゾートバイト！」"
    event[4] = "スキルアップイベント「企業研究をしよう！」"
    event[5] = "スキルアップイベント「自己分析をしよう！」"
    event[6] = "スキルアップイベント「自身の研究を進めよう！」"
    event[7] = "スキルアップ＋お金稼ぎイベント「報酬ありインターンに参加しよう！」"
    event[8] = "スキルアップイベント「企業研究をしよう！」"
    event[9] = "スキルアップイベント「自身の研究を進めよう！」"
    event[10] = "スキルアップイベント「学会発表をしよう！」"
    event[11] = "スキルアップイベント「論文を書こう！」"
    event[12] = "お金稼ぎイベント「接客アルバイトをしよう！」"
    event[13] = "お金稼ぎイベント「プログラミングアルバイトをしよう！」"
    event[14] = "スキルアップイベント「業界研究セミナーに参加しよう！」"
    event[15] = "スキルアップイベント「ある会社のインターンに参加しよう！」"
    event[16] = "スキルアップイベント「インターンに参加しよう！」"
    event[17] = "お金稼ぎイベント「単発バイトをしよう！」"
    event[18] = "内定イベント「サイバーエージェント」"
    event[19] = "内定イベント「NTT西日本」"
    event[20] = "内定イベント「マイクロンメモリジャパン」"
    event[21] = "内定イベント「パナソニック」"
    event[22] = "内定イベント「トヨタ自動車九州」"
    event[23] = "内定イベント「富士通」"
    event[24] = "内定イベント「SONY」"
    event[25] = "内定イベント「マツダ」"
    event[26] = "内定イベント「Amazon」"
    event[27] = "内定イベント「広島大学」"
    event[28] = "内定イベント「両備システムズ」"
    event[29] = "内定イベント「アクセンチュア」"
    event[30] = "内定イベント「キーエンス」"
    event[31] = "内定イベント「ソフトバンク」"
    event[32] = "内定イベント「日立製作所」"
    event[33] = "内定イベント「三菱総合研究所」"
    event[34] = "内定イベント「トヨタ自動車」"
    event[35] = "内定イベント「NTTデータ」"
    event[36] = "内定イベント「中国電力」"
    event[37] = "内定イベント「野村総合研究所」"
    event[38] = "内定イベント「Apple」"
    event[39] = "内定イベント「伊藤忠商事」"
    event[40] = "内定イベント「Google」"
    event[41] = "ボーナスステージ"
    event[42] = "ボーナスステージ"
    event[43] = "ボーナスステージ"
    event[44] = "ボーナスステージ"
    event[45] = "ボーナスステージ"
    event[46] = "ボーナスステージ"
	var name string
	var nyuushasaki int
    var input string
    var input_num int
	fmt.Print("名前を入力してください: ")
	fmt.Scan(&name)
	var player1 = newPlayer(name, 0,0,0,10000, nil) 
	fmt.Println("毎ターンアイテムを10000円で買うかどうか決めることができます。")
	fmt.Println("アイテムを買うと、1から6のサイコロの目を好きに決めることができます。")
	fmt.Println(player1.name, "の位置は", player1.position)
	for i := 0; i < 12 ; i++ {
		fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
		fmt.Println("現在", grade(i),"年生の",month(i),"月")//何年生の何カ月
		fmt.Println("コミュニケーションスキル:",player1.comSkill, "技術スキル:", player1.techSkill,"所持金:", player1.money)
		if player1.position < N {
			//６マス先のイベント内容を表示する処理
            //次のマス目を表示する関数
		    banmen(player1.position,event)
			//アイテムを使うかどうか
			fmt.Println("アイテムを買いますか？(yes/no):")
			fmt.Scanln(&input)
            var randomNumber int 

			if input == "yes"{
				if player1.money < 10000 {
					fmt.Println("お金が足りません。")
					//サイコロを振る
					fmt.Println("サイコロを振ります。エンターを押してください")
					bufio.NewReader(os.Stdin).ReadBytes('\n') // エンターキーを待つ

					rand.Seed(time.Now().UnixNano())
					randomNumber = rand.Intn(6)+1
				} else {
					fmt.Print("1から6で好きなマス目を入力してください（半角数字）：")
					fmt.Scanln(&input_num)
					randomNumber = input_num
                    player1.money -= 10000
				}
			} else {
				//サイコロを振る
				fmt.Println("サイコロを振ります。エンターを押してください")
				bufio.NewReader(os.Stdin).ReadBytes('\n') // エンターキーを待つ

				rand.Seed(time.Now().UnixNano())
				randomNumber = rand.Intn(6)+1
			}
			//positionの更新
			player1.position += randomNumber
			// if player1.position > N+1 {
			// 	player1.position = N+1
			// }
			//現状の表示
			fmt.Println("サイコロの目は", randomNumber)
			fmt.Println(player1.name, "の位置は", player1.position,"\n")
			//ここで、マス目ごとのイベント
			defineEvent(player1, player1.position)


			//残りのマスの表示
			fmt.Println("残りのマスは",N-player1.position,"\n")
			if player1.position == N {
				fmt.Println("就活終了！就職先を選ぼう！")
				//内定先の会社名を表示
				nyuushasaki=defineNaiteisaki(player1)
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
	fmt.Println("入社先:",player1.naiteisaki[nyuushasaki])
	fmt.Println("コミュニケーションスキル:",player1.comSkill, "技術スキル:", player1.techSkill,"所持金:", player1.money)
}