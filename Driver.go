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
}

func findOpenPlayer(hands [10]Hand) int {
	var index = -1
	for i:= 0; i < 9; i++ {
		if &hands[i] == nil {
			index = i-1
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
func toCard(card [2]string) Card {
	newCard := Card{
		Suit:card[0],
		Val:card[1],
	}

	return newCard
}
func compareTo(a Card, b Card) bool {
	return a.Suit == b.Suit && a.Val == b.Val
}

func addHand(deck Deck, table Table, isUser bool) {

	fmt.Println("Card One:")
	card1 := readCard()
	fmt.Println("Card Two:")
	card2 := readCard()

	if isValidCard(card1[0],card1[1]) && isValidCard(card2[0],card2[1]) {
		table.Players[findOpenPlayer(table.Players)] = Hand{
			Cards: [2]Card{ toCard(card1), toCard(card2) },
			isUser: isUser,
		}
	} else {
		addHand(deck, table, isUser)
	}

}
func dealFlop(board Board, table Table) {
	fmt.Println("Card One")
	flop1 := readCard()
	fmt.Println("Card Two")
	flop2 := readCard()
	fmt.Println("Card Three")
	flop3 := readCard()

	board.Cards[0] = toCard(flop1)
	board.Cards[1] = toCard(flop2)
	board.Cards[2] = toCard(flop3)

	isFlop := true

	for isFlop {
		printFlopMenu()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "A":
			fmt.Println("Calculate Winning Odds")
			calculateWinOdds(board, table)
		case "B":
			fmt.Println("Calculate Making Hand Odds")
			calculateHandOdds(board, table)
		case "C":
			fmt.Println("Show Turn")
			dealTurn(board, table)
			isFlop = false
		}

	}

}
func dealTurn(board Board, table Table) {

	fmt.Println("Card 4")
	turn := readCard()
	board.Cards[3] = toCard(turn)
	isTurn := true

	for isTurn {
		printTurnMenu()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "A":
			fmt.Println("Calculate Winning Odds")
			calculateWinOdds(board, table)
		case "B":
			fmt.Println("Calculate Making Hand Odds")
			calculateHandOdds(board, table)
		case "C":
			fmt.Println("Show River")
			dealRiver(board, table)
			isTurn = false
		}

	}
}
func dealRiver(board Board, table Table) {

	fmt.Println("Card 5")
	river := readCard()
	board.Cards[3] = toCard(river)
	isRiver := true

	for isRiver {
		printRiverMenu()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "A":
			fmt.Println("Calculate Winning Odds")
			calculateWinOdds(board, table)
		case "B":
			fmt.Println("Calculate Making Hand Odds")
			calculateHandOdds(board, table)
		case "C":
			fmt.Println("Clear Board")
			clearBoard(board)
			isRiver = false
		}

	}
}
func clearBoard(board Board) {
	board.Cards = [5]Card{}
}

func calculateWinOdds(board Board, table Table) {

}

func calculateHandOdds(board Board, table Table) {

}

func main() {

	suits := [4]string{"H","D","S","C"}
	vals := [13]string{"2","3","4","5","6","7","8","9","10","J","Q","K"}

	var deck = Deck{Cards: createDeck(suits, vals)}
	var table = Table{}
	var board = Board{}

	deck.Cards[0].inHand = true
	fmt.Println(deck.Cards[0])

	var isRunning = true

	for isRunning {

		printMenu()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "A":
			fmt.Println("\nAdd Your Hand")
			addHand(deck, table, true)
		case "B":
			fmt.Println("\nAdd Opponents Hand")
			addHand(deck, table, false)
		case "C":
			fmt.Println("\nDeal Flop")
			dealFlop(board, table)
		case "X":
			fmt.Println("\nQuit")
			isRunning = false
		default:
			fmt.Println("\nNone")
		}
	}
}
