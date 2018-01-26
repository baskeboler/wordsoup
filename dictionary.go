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

var charReplacements = map[string]string{
	"Á": "A",
	"É": "E",
	"Í": "I",
	"Ó": "O",
	"Ú": "U",
}

// DefaultDictURL points to URL of palabras.txt in github repository
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
			for invalid, replacement := range charReplacements {
				s = strings.Replace(s, invalid, replacement, len(s))
			}
			words = append(words, s)
		} else if e != nil {
			break
		}
	}
	return words
}

// NewDictionaryFromURL loads dictionary from file specified by URL
func NewDictionaryFromURL(url string) (Dictionary, error) {
	var words []string
	res, err := http.Get(url)
	if err != nil {
		return nil, ErrDictionaryLoadFailure
	}
	// if file not found
	if res.StatusCode == 404 {
		return nil, ErrDictionaryLoadFailure
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
