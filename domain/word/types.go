package word

type Word struct {
	word []Letter
}

func (w *Word) New(myWord []Letter) {
	w.word = myWord
}

func (w *Word) AppendLetter(l Letter) {
	w.word = append(w.word, l)
}

func (w *Word) Get() []Letter {
	return w.word
}

type Letter struct {
	char string
}

func (l *Letter) New(char string) {
	l.char = char
}

func (l *Letter) Get() string {
	return l.char
}

func ConvertWordToString(w Word) string {
	var objResult string
	for _, letter := range w.Get() {
		objResult = objResult + letter.Get()
	}
	return objResult
}
func ConvertStringToWord(s string) Word {
	var objResult Word
	for _, letter := range s {
		var l Letter
		l.New(string(letter))
		objResult.AppendLetter(l)
	}
	return Word{objResult.Get()}
}
