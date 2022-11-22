package main

import (
	"fmt"
	"jogo_da_forca/domain/game"
	"jogo_da_forca/domain/word"
)

func main() {
	myStorage := game.NewStorage()
	myGame := game.NewGame(myStorage)

	myGame.SetWord("BANANA")
	partialResult, err := myGame.GetPartialResult()
	if err != nil {
		fmt.Printf("word not found: %s", err)
	}
	fmt.Printf(word.ConvertWordToString(partialResult))
}
