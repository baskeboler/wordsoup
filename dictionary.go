package wordsoup

import (
	"bufio"
	"errors"
	"io"
	"math/rand"
	"net/http"
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

const DefaultDictURL = "https://rawgit.com/baskeboler/wordsoup/master/palabras.txt"

// NewDictionary initializes a new Dictionary
func NewDictionary() (Dictionary, error) {
	return NewDictionaryFromURL(DefaultDictURL)
}

// NewDictionaryFromTextFile initializes a new Dictionary from a text file
func NewDictionaryFromTextFile(path string) (Dictionary, error) {
	f, e := os.Open(path)
	var words []string
	if e != nil {
		return nil, ErrDictionaryLoadFailure
	}
	defer f.Close()
	words = readWords(f)
	rand.Seed(time.Now().UnixNano())
	return &dictionaryImpl{words: words}, nil
}

func readWords(r io.Reader) []string {
	var words []string

	reader := bufio.NewReader(r)
	for {
		if s, e := reader.ReadString('\n'); e == nil && len(s) > 5 {
			s = strings.TrimSpace(s)
			s = strings.ToUpper(s)
			words = append(words, s)
		} else if e != nil {
			break
		}
	}
	return words
}

func NewDictionaryFromURL(url string) (Dictionary, error) {
	var words []string
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body := res.Body
	defer body.Close()
	words = readWords(body)
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
