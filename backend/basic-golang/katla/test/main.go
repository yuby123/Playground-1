package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const wordLength = 5
const maxGuess = 6

//NOTE: err handling is not yet taught, we don't handle errors in this example
//don't worry about the content of this method for now. We haven't learn some concepts
func getDictionaryWords() []string {
	kbbiURL := "https://gist.githubusercontent.com/fikriauliya/c7024f9629ba7d515f01a625c66a4f2f/raw/141f629d452145ce7e02215a98cde04d9f1bbb20/kbbi.txt"

	data, err := http.Get(kbbiURL)
	if err != nil {
		panic(err)
	}

	defer data.Body.Close()

	body, err := io.ReadAll(data.Body)
	if err != nil {
		panic(err)
	}

	lines := string(body)
	words := make([]string, 0)

	// check regex buat word
	for _, line := range strings.Split(lines, "\n") {

	}
	return words
}

// this is not optimal search algorithm, but it's a good example
func isInDictionary(word string, dictionary []string) bool {
	for _, entry := range dictionary {
		if entry == word {
			return true
		}
	}
	return false
}

type hint int

const (
	notFound        hint = 0
	correctPosition hint = 1
	correctLetter   hint = 2
)

func calculateHints(guess, answer string) (hints []hint) {
	guessChars := []rune(guess)   //["t","e","s","t","a"]
	answerChars := []rune(answer) //["d", "o", "m", "b", "a"]

	hints = make([]hint, wordLength) //[[],[],[],[],[]]

	for i := 0; i < wordLength; i++ {
		if guessChars[i] == answerChars[i] {
			hints[i] = correctPosition
		} else {
			// karakter ada atau di jawabannya
			for j := 0; j < wordLength; j++ {
				if i != j { //0 != 1
					//when the answer is:
					//STROK, and we guess:
					//SOSOK
					//the answer should be:
					//GXXGG
					//not:
					//GYYGG
					//Reason: the second 'O' has been marked as correct position ('Y')
					//if we mark 'Y' for the first 'O', people would guess there should be yet another 'O'
					//while in fact there is only one 'O' in 'STROK'

					// ["t"] == ["o"] dan ["e"] != ["o"]

				}
			}
		}
	}
	return
}

func main() {
	// ambil kamus kata
	dictionary := getDictionaryWords()

	// dapetin random berdasarkan waktu
	rand.Seed(time.Now().UnixNano())

	// ambil jawaban random dari kamus
	answer := dictionary[rand.Intn(len(dictionary))]
	// fmt.Printf("Answer: %s\n", answer)

	isWin := false
	for trial := 0; trial < maxGuess; trial++ {
		var guess string
		fmt.Printf("Guess %d: \n", trial+1)
		// ngecheck kondisi
		for {

		}

		// kasih character
		hints := calculateHints(guess, answer) //answer = "domba" guess="test"
		for i := 0; i < wordLength; i++ {      //0 1 2 3 4

		}
		fmt.Println()
		fmt.Println()

		if guess == answer {
			fmt.Println("You win!")
			isWin = true
			break
		}
	}

	//print kalo lose
	if !isWin {
		fmt.Printf("The correct answer is: %s\n", answer)
	}
}
