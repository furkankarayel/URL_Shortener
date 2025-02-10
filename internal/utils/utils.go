package utils

import (
	"fmt"
	"math/rand"
	"path"
	"strings"
	"time"
)

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

var consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "x", "y", "z"}
var vowels = []string{"a", "e", "i", "o", "u"}
var nouns = []string{"master", "developer", "designer", "programmer", "coder", "engineer", "scientist", "artist", "writer", "musician", "lawyer", "doctor", "teacher", "scientist", "artist", "writer", "musician", "lawyer", "doctor", "teacher"}

func GenerateHybridString() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a pronounceable syllable
	var syllable strings.Builder
	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			syllable.WriteString(consonants[rand.Intn(len(consonants))])
		} else {
			syllable.WriteString(vowels[rand.Intn(len(vowels))])
		}
	}

	noun := nouns[rand.Intn(len(nouns))]

	number := rand.Intn(100)

	return strings.Title(syllable.String()) + strings.Title(noun) + fmt.Sprintf("%02d", number)
}
