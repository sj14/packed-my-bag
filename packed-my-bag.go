package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		bag      []string
		input    string
		answered      = true
		player   uint = 1
		players  uint
	)

	clear()
	fmt.Printf("\rNumber Of Players?:\n")

	_, err := fmt.Scan(&players)
	if err != nil {
		fmt.Println(err)
	}

	clear()
	for answered {

		fmt.Printf("\rPlayer %d \n", player)
		fmt.Printf("\rAdd New Object:\n")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
		}

		// Sleep 3 seconds to let the players remember the word
		for i := 3; i > 0; i-- {
			fmt.Printf("\r %v", i)
			time.Sleep(1 * time.Second)
		}
		clear()
		bag = append(bag, input)

		player++
		if player > players {
			player = 1
		}

		for key := range bag {

			fmt.Printf("\rPlayer %d \n", player)
			fmt.Printf("\rRemember Object %v:\n", key+1)

			_, err = fmt.Scan(&input)
			if err != nil {
				fmt.Println(err)
			}

			clear()
			if input != bag[key] {
				answered = false
				fmt.Printf("Wrong, your said %s but %s was correct.\n", input, bag[key])
				fmt.Printf("Player %d lost the game!\n\n", player)
				fmt.Printf("Contents of the bag:\n")

				for key := range bag {
					fmt.Printf("%d: %s\n", key+1, bag[key])
				}
				break
			}
		}
	}
}

// Clears the console
func clear() {
	fmt.Printf("\033[2J")
}
