package wordsoup_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/baskeboler/wordsoup"
)

func TestWS(t *testing.T) {
	rand.Seed(time.Now().Unix())

	s := &wordsoup.WordSoup{W: 40, H: 40}
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
	dict, err := wordsoup.NewDictionary()
	if err != nil {
		t.Error(err)
		return
	}
	ws, err := wordsoup.GenerateRandomWordSoup(55, 55, 300, dict)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ws)
	t.Log(ws.Words)
}
