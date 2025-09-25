package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
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
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
	// If you have a pair of aces you must always split them.
	case card1 == "ace" && card2 == "ace":
		return "P"
	// If you have a Blackjack...
	case ParseCard(card1)+ParseCard(card2) == 21:
		// ...and the dealer has a ace or ten-valued card, you must stand.
		if dealerCard == "ace" || ParseCard(dealerCard) == 10 {
			return "S"
		}
		// ...and the dealer does not have a ace or value of ten card, you automatically win.
		return "W"
	// If you have a hand value within the range 17-20, you must always stand.
	case ParseCard(card1)+ParseCard(card2) >= 17:
		return "S"
	// If your cards sum up to 11 or lower you should always hit.
	case ParseCard(card1)+ParseCard(card2) <= 11:
		return "H"
	// If your cards sum up to 12-16...
	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16:
		// ...and the dealer has a 7 or higher, you must always hit.
		if ParseCard(dealerCard) >= 7 {
			return "H"
		}
		// ...and the dealer has a 6 or lower, you must always stand.
		return "S"
	}
	panic("No decision available in Alex's holy strategy.")
}
