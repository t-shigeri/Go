package main

import "fmt"

func main() {
	var channelNumber uint32
	var tvStationName string

	fmt.Print("チャンネル番号を入力してください: ")
	fmt.Scan(&channelNumber)

	switch channelNumber {
	case 1:
		tvStationName = "NHK総合"
	case 2:
		tvStationName = "NHK教育"
	case 4:
		tvStationName = "日本テレビ"
	case 5:
		tvStationName = "TBS"
	case 6:
		tvStationName = "フジテレビ"
	case 7:
		tvStationName = "朝日テレビ"
	case 12:
		tvStationName = "東京12チャンネル"
	default:
		tvStationName = "ありません"
	}

	fmt.Printf("チャンネル番号 %d のテレビ局名は %s です\n", channelNumber, tvStationName)
}
