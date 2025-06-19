package main

import "fmt"

func tvStationName(channelNumber uint) string {
    switch channelNumber {
    case 1:
        return "NHK総合"
    case 2:
        return "NHK教育"
    case 4:
        return "日本テレビ"
    case 5:
        return "テレビ朝日"
    case 6:
        return "TBSテレビ"
    case 7:
        return "テレビ東京"
    case 8:
        return "フジテレビ"
    default:
        return "該当なし"
    }
}

func main() {
    var channelNumber uint
    fmt.Print("チャンネル番号入力 = ")
    fmt.Scan(&channelNumber)
    fmt.Printf("チャンネル番号 %d に対応するテレビ局名は、%s¥n", channelNumber, tvStationName(channelNumber))
}