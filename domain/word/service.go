package word

type Service interface {
	SetWord(word string) error
	GetWord() (Word, error)
	CleanWord() (Word, error)
}
