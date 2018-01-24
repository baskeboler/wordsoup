package wordsoup_test

import (
	"testing"

	"github.com/baskeboler/wordsoup"
)

func TestDictionary(t *testing.T) {
	d, err := wordsoup.NewDictionary()
	if err != nil {
		t.Error(err)
	}
	w := d.GetRandomWord()
	t.Log(w)
	if len(w) < 5 {
		t.Fail()
	}
}

func TestGetWords(t *testing.T) {
	d, err := wordsoup.NewDictionary()
	if err != nil {
		t.Error(err)
	}
	w := d.GetRandomWords(100)
	t.Log(w)
	if len(w) != 100 {
		t.Fail()
	}
}
