package wordsoup

import (
	"math/rand"
	"testing"
	"time"
)

func TestWS(t *testing.T) {
	rand.Seed(time.Now().Unix())

	s := &WordSoup{W: 40, H: 40}
	s.TryToAddWord("Victor")
	s.TryToAddWord("Olivia")
	for i := 0; i < 100; i++ {
		s.TryToAddWord("fer")
		s.TryToAddWord("Victor")
		s.TryToAddWord("Olivia")
	}
	t.Logf("\n%v", s.String())
}

func TestRandomWS(t *testing.T) {
	dict, err := NewDictionary()
	if err != nil {
		t.Error(err)
		return
	}
	ws, err := GenerateRandomWordSoup(55, 55, 300, dict)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ws)
	t.Log(ws.Words)
}
func TestAddWordsFailure(t *testing.T) {
	s := &WordSoup{W: 10, H: 10}
	word := "Supercalifragilisticexpialidocious"
	err := s.TryToAddWord(word)
	if err != ErrFailedToAddWord {
		t.Fail()
	}
}

func TestAddTooManyWordFailure(t *testing.T) {
	dict, err := NewDictionary()
	if err != nil {
		t.Error(err)
		return
	}
	_, err = GenerateRandomWordSoup(20, 20, 300, dict)
	if err != ErrGenerationFailure {
		t.Fail()
	}
	_, err = GenerateRandomWordSoup(0, 20, 2, dict)
	if err != ErrGenerationFailure {
		t.Fail()
	}
}
