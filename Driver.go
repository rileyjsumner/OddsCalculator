package main

import (
	"bufio"
	"fmt"
	"os"
)

type Deck struct {
	Cards [52]Card
}
type Card struct {
	Suit, Val string
	inHand, onBoard bool
}
type Hand struct {
	Cards [2]Card
	isUser bool
}
type Board struct {
	Cards [5]Card
}
type Table struct {
	Players [10]Hand
	Deck Deck
	Board Board
}

func isEmptyHand(hand Hand) bool {
	return hand.Cards[0].Val == "" &&
			hand.Cards[0].Suit == "" &&
			hand.Cards[1].Val == "" &&
			hand.Cards[1].Suit == "" &&
			!hand.isUser
}
func findOpenPlayer(hands [10]Hand) int {
	var index = -1
	fmt.Println("find open player")
	for i:= 0; i < 9; i++ {
		fmt.Println(hands[i])
		if isEmptyHand(hands[i]) {
			index = i
			break
		}
	}
	return index
}
func createDeck(suits [4]string, vals [13]string) [52]Card {
	var cards [52]Card
	for s := 0; s < len(suits); s++ {
		for v:= 0; v < len(vals); v++ {
			cards[(13*s)+v] = Card{Suit: suits[s], Val: vals[v]}
		}
	}
	return cards
}
func printMenu() {
	fmt.Println("SELECT AN OPTION")
	fmt.Println("[A] Add Your Hand")
	fmt.Println("[B] Add Opponents Hand")
	fmt.Println("[C] Deal Flop")
	fmt.Println("[X] Quit")
}
func printFlopMenu() {
	fmt.Println("Pick an Option")
	fmt.Println("[A] Calculate Winning Odds")
	fmt.Println("[B] Calculate Making a Hand Odds")
	fmt.Println("[C] Show Turn")
}
func printTurnMenu() {
	fmt.Println("Pick an Option")
	fmt.Println("[A] Calculate Winning Odds")
	fmt.Println("[B] Calculate Making a Hand Odds")
	fmt.Println("[C] Show River")
}
func printRiverMenu() {
	fmt.Println("Pick an Option")
	fmt.Println("[A] Calculate Winning Odds")
	fmt.Println("[B] Calculate Making a Hand Odds")
	fmt.Println("[C] Clear Board")
}
func printHandsMenu() {
	fmt.Println("What are the odds of hitting a...")
	fmt.Println("[A] Pair")
	fmt.Println("[B] Two Pair")
	fmt.Println("[C] Three of a Kind")
	fmt.Println("[D] Straight")
	fmt.Println("[E] Flush")
	fmt.Println("[F] Straight Flush")
	fmt.Println("[G] Royal Flush")
	fmt.Println("[X] Exit")
}
func readCard() [2]string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Suit")
	scanner.Scan()
	s1 := scanner.Text()

	fmt.Println("Value")
	scanner.Scan()
	v1 := scanner.Text()

	return [2]string{s1,v1}
}

func isValidCard(suit string, val string) bool {

	isValid := false

	switch suit {
	case "H",
		 "D",
		 "S",
		 "C":
		 	isValid = true
	}
	if isValid {
		switch val {
		case "2",
			 "3",
			 "4",
		 	 "5",
			 "6",
			 "7",
			 "8",
			 "9",
			 "10",
			 "J",
			 "Q",
			 "K",
			 "A":
				isValid = true
		default:
			isValid = false
		}
	}
	return isValid
}
func toCard(card [2]string, inHand bool, onBoard bool) Card {
	newCard := Card{
		Suit:card[0],
		Val:card[1],
		inHand:inHand,
		onBoard:onBoard,
	}

	return newCard
}
func cardsToString(cards [5]Card) [5]string {
	var strArr [5]string
	for i := 0; i < len(cards); i++ {
		strArr[i] = cards[i].Val
	}
	return strArr
}
func compareTo(a Card, b Card) bool {
	return a.Suit == b.Suit && a.Val == b.Val
}
func hasDuplicate(arr [5]string) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}
func getUserHand(table Table) int {
	for i := 0; i < len(table.Players); i++ {
		if table.Players[i].isUser {
			return i
		}
	}
	return -1
}

func addHand(table Table, isUser bool) Table {

	fmt.Println("Card One:")
	card1 := readCard()
	fmt.Println("Card Two:")
	card2 := readCard()

	if isValidCard(card1[0],card1[1]) && isValidCard(card2[0],card2[1]) {
		table.Players[findOpenPlayer(table.Players)] = Hand{
			Cards: [2]Card{ toCard(card1, true, false), toCard(card2, true, false) },
			isUser: isUser,
		}
	} else {
		fmt.Println("Invalid Card")
		table = addHand(table, isUser)
	}

	return table

}
func dealFlop(table Table) Table {
	fmt.Println("Card One")
	flop1 := readCard()
	fmt.Println("Card Two")
	flop2 := readCard()
	fmt.Println("Card Three")
	flop3 := readCard()

	table.Board.Cards[0] = toCard(flop1, false, true)
	table.Board.Cards[1] = toCard(flop2, false, true)
	table.Board.Cards[2] = toCard(flop3, false, true)

	isFlop := true

	for isFlop {
		printFlopMenu()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "A":
			fmt.Println("Calculate Winning Odds")
			table = calculateWinOdds(table)
		case "B":
			fmt.Println("Calculate Making Hand Odds")
			table = calculateHandOdds(table)
		case "C":
			fmt.Println("Show Turn")
			table = dealTurn(table)
			isFlop = false
		}

	}
	return table
}
func dealTurn(table Table) Table {

	fmt.Println("Card 4")
	turn := readCard()
	table.Board.Cards[3] = toCard(turn, false, true)
	isTurn := true

	for isTurn {
		printTurnMenu()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "A":
			fmt.Println("Calculate Winning Odds")
			table = calculateWinOdds(table)
		case "B":
			fmt.Println("Calculate Making Hand Odds")
			table = calculateHandOdds(table)
		case "C":
			fmt.Println("Show River")
			table = dealRiver(table)
			isTurn = false
		}

	}
	return table
}
func dealRiver(table Table) Table {

	fmt.Println("Card 5")
	river := readCard()
	table.Board.Cards[4] = toCard(river, false, true)
	isRiver := true

	for isRiver {
		printRiverMenu()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "A":
			fmt.Println("Calculate Winning Odds")
			table = calculateWinOdds(table)
		case "B":
			fmt.Println("Calculate Making Hand Odds")
			table = calculateHandOdds(table)
		case "C":
			fmt.Println("Clear Board")
			clearBoard(table.Board)
			isRiver = false
		}

	}
	return table
}
func clearBoard(board Board) {
	board.Cards = [5]Card{}
}

func calculateWinOdds(table Table) Table {
	return table
}
func calculateHandOdds(table Table) Table {
	for i := 0; i < len(table.Players); i++ {
		if table.Players[i].isUser {
			printHandsMenu()

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			choice := scanner.Text()
			isCalculating := true

			for isCalculating {

				var odds = float64(0)
				switch choice {
				case "A": // pair
					odds = float64(getPairOdds(table))
				case "B": // two pair
					odds = float64(getTwoPairOdds(table))
				case "C": // three of a kind
					odds = float64(getTripsOdds(table))
				case "D": // straight
					odds = float64(getStraightOdds(table))
				case "E": // flush
					odds = float64(getFlushOdds(table))
				case "F": // straight flush
					odds = float64(getStraightFlushOdds(table))
				case "G": // royal
					odds = float64(getRoyalOdds(table))
				case "X": // exit
					isCalculating = false
				}
				fmt.Println("Odds of hitting: ", odds)
				table = calculateHandOdds(table)
			}
		}
	}
	return table
}
func getRemainingCards(table Table) [52]Card {
	remainingCards := [52]Card{}
	var counter = 0
	for i := 0; i < len(table.Deck.Cards); i++ {
		if !table.Deck.Cards[i].onBoard && !table.Deck.Cards[i].inHand {
			remainingCards[counter] = table.Deck.Cards[i]
		}
	}
	return remainingCards
}

func getPairOdds(table Table) float64 {
	cards := getRemainingCards(table)
	hand := table.Players[getUserHand(table)]
	board := table.Board

	if hand.Cards[0].Val == hand.Cards[1].Val {
		return 1
	} else {
		var outs = 0
		var cardsToSee = 5
		if &board.Cards[0] == nil {
			for i := 0; i < len(cards); i++ {
				if hand.Cards[0].Val == cards[i].Val || hand.Cards[1].Val == cards[i].Val {
					outs++
				}
			}
		} else if &board.Cards[3] == nil {
			cardsToSee = 2
			if !hasDuplicate(cardsToString(board.Cards)) {
				for i := 0; i < len(cards); i++ {
					if hand.Cards[0].Val == cards[i].Val || hand.Cards[1].Val == cards[i].Val {
						outs++
					}
				}
				for i := 0; i < len(board.Cards); i++ {
					if hand.Cards[0].Val == board.Cards[i].Val || hand.Cards[1].Val == board.Cards[i].Val {
						return 1
					}
				}
			} else {
				return 1
			}
		} else if &board.Cards[4] == nil {
			cardsToSee = 1
			if !hasDuplicate(cardsToString(board.Cards)) {
				for i := 0; i < len(cards); i++ {
					if hand.Cards[0].Val == cards[i].Val || hand.Cards[1].Val == cards[i].Val {
						outs++
					}
				}
				for i := 0; i < len(board.Cards); i++ {
					if hand.Cards[0].Val == board.Cards[i].Val || hand.Cards[1].Val == board.Cards[i].Val {
						return 1
					}
				}
			} else {
				return 1
			}
		} else {
			cardsToSee = 0
			if !hasDuplicate(cardsToString(board.Cards)) {
				for i := 0; i < len(board.Cards); i++ {
					if hand.Cards[0].Val == board.Cards[i].Val || hand.Cards[1].Val == board.Cards[i].Val {
						return 1
					}
				}
			} else {
				return 1
			}
		}
		return float64(cardsToSee / outs)
	}

}
func getTwoPairOdds(table Table) float64 {
	return 0
}
func getTripsOdds(table Table) float64 {
	return 0
}
func getStraightOdds(table Table) float64 {
	return 0
}
func getFlushOdds(table Table) float64 {
	return 0
}
func getStraightFlushOdds(table Table) float64 {
	return 0
}
func getRoyalOdds(table Table) float64 {
	return 0
}

func main() {

	suits := [4]string{"H","D","S","C"}
	vals := [13]string{"2","3","4","5","6","7","8","9","10","J","Q","K"}

	var deck = Deck{Cards: createDeck(suits, vals)}
	var board = Board{}
	var table = Table{
		Board: board,
		Deck: deck,
	}

	deck.Cards[0].inHand = true

	var isRunning = true

	for isRunning {

		printMenu()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "A":
			fmt.Println("\nAdd Your Hand")
			table = addHand(table, true)
		case "B":
			fmt.Println("\nAdd Opponents Hand")
			table = addHand(table, false)
		case "C":
			fmt.Println("\nDeal Flop")
			table = dealFlop(table)
		case "X":
			fmt.Println("\nQuit")
			isRunning = false
		default:
			fmt.Println("\nNone")
		}
	}
}
