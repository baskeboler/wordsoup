package wordsoup

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Orientation int

var (
	ErrFailedToAddWord   = errors.New("Could not add word to soup")
	ErrGenerationFailure = errors.New("Failed to generate wordsoup")
)

const (
	Horizontal Orientation = iota
	Vertical
	AddRetries int = 10
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Position struct {
	X, Y int
}

type Word struct {
	Letters     string
	Orientation Orientation
	Pos         Position
}

type WordSoup struct {
	W, H  int
	Words []Word
}

type Cell struct {
	Position
	Letter rune
}

// func RandomPos(maxX, maxY int) Position {
// 	rand.New
// }
func (p Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

func (w Word) Cells() []Cell {
	var res []Cell
	for i, l := range w.Letters {
		p := w.Pos
		if w.Orientation == Vertical {
			p.Y += i
		} else {
			p.X += i
		}
		res = append(res, Cell{Position: p, Letter: l})
	}
	return res
}

func (w Word) Conflicts(other Word) bool {

	cells := w.Cells()
	for _, othersCell := range other.Cells() {
		for _, myCell := range cells {
			if myCell.Position.Equals(othersCell.Position) && myCell.Letter != othersCell.Letter {
				return true
			}
		}
	}
	return false
}

func (s *WordSoup) Fits(w Word) bool {
	if w.Orientation == Vertical {
		if w.Pos.Y+len(w.Letters) > s.H {
			return false
		}
	} else {
		if w.Pos.X+len(w.Letters) > s.W {
			return false
		}
	}
	return true
}
func (s *WordSoup) Conflicts(w Word) bool {
	for _, myWord := range s.Words {
		if myWord.Conflicts(w) {
			return true
		}
	}
	return false
}
func (s *WordSoup) TryToAddWord(word string) error {
	if len(word) > s.H && len(word) > s.W {
		return ErrFailedToAddWord
	}
	for try := 0; try < 10000; try++ {

		p := Position{rand.Intn(s.W), rand.Intn(s.H)}
		orient := Vertical
		if rand.Int()%2 == 0 {
			orient = Horizontal
		}
		w := Word{Letters: strings.ToUpper(word), Pos: p, Orientation: orient}
		if s.Fits(w) && !s.Conflicts(w) {
			s.Words = append(s.Words, w)
			return nil
		}
	}
	return ErrFailedToAddWord
}

func (s *WordSoup) render() map[int]rune {
	res := make(map[int]rune)

	for _, w := range s.Words {
		for _, c := range w.Cells() {
			res[c.Y*s.W+c.X] = c.Letter
		}
	}
	for i := 0; i < s.W*s.H; i++ {
		if _, occupied := res[i]; !occupied {
			// var r rune
			res[i] = letters[rand.Intn(len(letters))]

		}
	}
	return res
}

func (s *WordSoup) String() string {
	var res string
	m := s.render()
	for i := 0; i < s.W*s.H; i++ {
		if (i)%s.W == 0 {
			res = fmt.Sprintf("%s\n", res)
		}
		res = fmt.Sprintf("%s%c ", res, m[i])
	}
	return res
}

func GenerateRandomWordSoup(height, width, nWords int, dict Dictionary) (*WordSoup, error) {
	if height < 1 || width < 1 {
		return nil, ErrGenerationFailure
	}
	ws := &WordSoup{W: width, H: height}

MainLoop:
	for i := 0; i < nWords; i++ {
		for try := 0; try < AddRetries; try++ {
			if e := ws.TryToAddWord(dict.GetRandomWord()); e == nil {
				continue MainLoop
			}
		}
		return nil, ErrGenerationFailure
	}
	return ws, nil
}
