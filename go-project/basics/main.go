// Executable Package
package main

import "fmt"

func main() {
	/* Notes for Variables:
	* Types can be inferred from variable
	* := is only used for initialization
	* var card string= "Ace of Spades"
	 */

	card := newCard()
	fmt.Println(card)

	// Create a Slice of Cards
	cards := []string{newCard()}

	// Append() does not modify the existing slice- it creates a new slice with the added item
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// Explicitly define return type
func newCard() string {
	return "Five of Diamonds"
}
