package poker13tw

func GetCardSuit(card int) string {
	suits := []string{"黑桃", "愛心", "方塊", "梅花"}
	return suits[(card-1)/13]
}

func GetCardRank(card int) string {
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	return ranks[(card-1)%13]
}

func ConvertToCardValue(suit string, rank string) int {
	rankValue := 0
	switch rank {
	case "2":
		rankValue = 2
	case "3":
		rankValue = 3
	case "4":
		rankValue = 4
	case "5":
		rankValue = 5
	case "6":
		rankValue = 6
	case "7":
		rankValue = 7
	case "8":
		rankValue = 8
	case "9":
		rankValue = 9
	case "T":
		rankValue = 10
	case "J":
		rankValue = 11
	case "Q":
		rankValue = 12
	case "K":
		rankValue = 13
	case "A":
		rankValue = 14
	default:
		panic("無效的點數")
	}

	switch suit {
	case "黑桃":
		return rankValue - 1
	case "愛心":
		return rankValue - 1 + 13
	case "方塊":
		return rankValue - 1 + 26
	case "梅花":
		return rankValue - 1 + 39
	default:
		panic("無效的花色")
	}
}

func ConvertRankToPoint(rank string) int {
	rk := []string{"", "", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	for i, r := range rk {
		if r == rank {
			return i
		}
	}
	panic("無效的點數")
}

func ConvertCardToPoint(card int) int {
	rank := GetCardRank(card)
	point := ConvertRankToPoint(rank)
	return point

}
