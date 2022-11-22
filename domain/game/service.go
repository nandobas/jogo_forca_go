package game

import (
	"fmt"
	"jogo_da_forca/domain/word"
)

type Service interface {
	SetWord(word string) error
	TryLetter(char word.Letter) error
	TryWord(word word.Word) error
	GetTotalErrors() int
	GetPartialResult() (word.Word, error)
}

type game struct {
	storage Storage
}

func (g game) SetWord(w string) error {
	if len(w) > 20 {
		return fmt.Errorf("word must be max 20 character")
	}
	ObjWord := word.ConvertStringToWord(w)
	g.storage.SetWord(ObjWord)
	return nil
}

func (g game) TryLetter(char word.Letter) error {
	//TODO implement me
	panic("implement me")
}

func (g game) TryWord(word word.Word) error {
	//TODO implement me
	panic("implement me")
}

func (g game) GetTotalErrors() int {
	//TODO implement me
	panic("implement me")
}

func (g game) GetPartialResult() (word.Word, error) {
	objWord, err := g.storage.GetWord()
	if err != nil {
		return word.Word{}, fmt.Errorf("word not found")
	}
	return objWord, nil
}

func NewGame(s Storage) Service {
	return game{storage: s}
}
