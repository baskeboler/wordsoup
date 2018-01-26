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
func TestLoadFromURL(t *testing.T) {
	d, err := wordsoup.NewDictionaryFromURL("https://rawgit.com/baskeboler/wordsoup/master/palabras.txt")
	if err != nil {
		t.Error(err)
	}
	w := d.GetRandomWord()
	if len(w) <= 0 {
		t.Fail()
	}
}
func TestLoadFromURLFailure(t *testing.T) {
	_, err := wordsoup.NewDictionaryFromURL("https://rawgit.com/baskeboler/wordsoup/master/palabrasNON-EXISTENT.txt")
	if err != wordsoup.ErrDictionaryLoadFailure {
		t.Fail()
	}
}
func TestDefaultLoad(t *testing.T) {
	d, err := wordsoup.NewDictionary()
	if err != nil {
		t.Error(err)
	}
	w := d.GetRandomWord()
	if len(w) <= 0 {
		t.Fail()
	}
}
