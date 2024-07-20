package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "other":
		return 0
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Int := ParseCard(card1)
	card2Int := ParseCard(card2)
	dealerCardInt := ParseCard(dealerCard)
	playerSum := card1Int + card2Int
	switch {
	case playerSum == 21 && dealerCard != "ace" && dealerCard != "ten" && dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king":
		return "W"
	case playerSum >= 17 && playerSum <=20:
		return "S"	
	case playerSum >= 12 && playerSum <=16 && dealerCardInt >=7:
		return "H"
	case playerSum <=11:
		return "H"
	case card1 == "ace" && card2 == "ace":
		return "P"	
	default:
		return "S"
	}
}
