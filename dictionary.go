package wordsoup

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Dictionary main interface
type Dictionary interface {

	// GetRandomWord returns a random word
	GetRandomWord() string

	// GetRandomWords returns a slice of n random words
	GetRandomWords(n int) []string
}

type dictionaryImpl struct {
	words []string
}

var (
	// ErrDictionaryLoadFailure is returned when dictionary file load fails
	ErrDictionaryLoadFailure = errors.New("Failed to load dictionary")
)

// NewDictionary initializes a new Dictionary
func NewDictionary() (Dictionary, error) {
	f, e := os.Open("palabras.txt")
	var words []string
	if e != nil {
		return nil, ErrDictionaryLoadFailure
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		if s, e := reader.ReadString('\n'); e == nil && len(s) > 5 {
			s = strings.TrimSpace(s)
			s = strings.ToUpper(s)
			words = append(words, s)
		} else if e != nil {
			break
		}
	}
	rand.Seed(time.Now().UnixNano())
	return &dictionaryImpl{words: words}, nil
}

func (d *dictionaryImpl) GetRandomWord() string {

	return d.words[rand.Intn(len(d.words))]
}

func (d *dictionaryImpl) GetRandomWords(n int) []string {
	var res []string
	for i := 0; i < n; i++ {
		res = append(res, d.GetRandomWord())
	}
	return res
}
