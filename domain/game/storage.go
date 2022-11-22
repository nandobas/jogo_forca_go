package game

import "jogo_da_forca/domain/word"

type Storage interface {
	SetWord(word word.Word) error
	GetWord() (word.Word, error)
}

type storage struct {
	word word.Word
}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) SetWord(myWord word.Word) error {
	s.word = myWord
	return nil
}

func (s *storage) GetWord() (word.Word, error) {
	return s.word, nil
}
