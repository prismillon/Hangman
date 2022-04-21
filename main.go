package main

import (
	hangman "hangman/.fonction"
	"os"
	"strings"
)

func main() {
	hangman.MemeLaTailleDuCmdOnYAPense()
	hangman.RandInit()
	args := os.Args[1:]
	if len(args) == 0 {
		hangman.SlowPrintln("Merci de fournir un nom de fichier.")
		os.Exit(0)
	}
	// Create our GameState instance
	var challenge = &hangman.GameState{}

	// Check for arguments
	load := hangman.FindArgValue([]string{"-sw", "--startWith"}, args, true)
	ascii := hangman.FindArgValue([]string{"-lf", "--letterFile"}, args, true)

	// If no save is loaded, recreate an instance of GameState from args[0]
	// Else, extract saved data from file
	if load == "" {
		challenge = hangman.NewGame(hangman.GetRandomWord(args[0]))
	} else {
		_, err := os.Stat(load)
		if err != nil {
			hangman.SlowPrintln("La sauvegarde \"" + load + "\" n'existe pas.")
			os.Exit(0)
		}
		challenge.Save(false, load)
	}

	// If ascii is set, add the ascii flag and the ascii alphabet to the current
	// instance of GameState
	if ascii != "" {
		alphabet := hangman.GetFile(ascii, "\n\n")
		challenge.AsciiChar = true
		for _, el := range alphabet {
			letter := strings.Split(el, "\r\n")
			if len(letter) == 1 {
				letter = strings.Split(el, "\n")
			}
			challenge.AsciiLetters = append(challenge.AsciiLetters, letter)
		}
	}

	challenge.Jose = hangman.GetFile("hangman.txt", "\n\n")

	// Main loop
	repeat := 0
	for {
		challenge.DisplayCurrentState()
		guess := hangman.ValidInput(hangman.Message(repeat), challenge)
		if len(guess) >= 4 && guess[:4] == "stop" {
			challenge.Save(true, guess[5:])
		} else if len(guess) > 1 {
			if guess == hangman.RemoveAccents(string(challenge.Total)) {
				break
			} else {
				challenge.ErrorCount += 2
			}
		} else if hangman.IsIn(guess, challenge.UsedLetters) {
			repeat++
			if repeat > 4 {
				challenge.ErrorCount = 9
				break
			}
			continue
		} else if !challenge.AddLetter(guess) {
			challenge.ErrorCount++
		}
		challenge.UsedLetters = append(challenge.UsedLetters, guess)
		repeat = 0

		// break if lose / win
		if challenge.IsFinish() {
			break
		}
	}
	hangman.ClearConsole()
	os.Remove(challenge.SavedFile)
	challenge.End()
}
