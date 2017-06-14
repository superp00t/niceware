package niceware

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/superp00t/niceware/words"
)

func BytesToPassphrase(input []byte) ([]string, error) {
	if len(input)%2 == 1 {
		return nil, fmt.Errorf("Only even-sized byte arrays are supported.")
	}

	var sentence []string

	for index, char := range input {
		if index+1 == len(input) {
			break
		}

		next := input[index+1]

		if index%2 == 0 {
			wordIndex := int(char)*256 + int(next)
			if words.WordList[wordIndex] == "" {
				return nil, fmt.Errorf("Could not convert byte 0x%x to word", wordIndex)
			} else {
				sentence = append(sentence, words.WordList[wordIndex])
			}
		}
	}

	return sentence, nil
}

func PassphraseToBytes(input []string) ([]byte, error) {
	decoded := make([]byte, len(input)*2)

	for index, word := range input {
		word = strings.ToLower(word)

		wordIndex := -1
		for i, wd := range words.WordList {
			if wd == word {
				wordIndex = i
				break
			}
		}

		if wordIndex == -1 {
			return nil, fmt.Errorf("Invalid word: %s", word)
		}

		decoded[2*index] = byte(wordIndex / 256)
		decoded[2*index+1] = byte(wordIndex % 256)
	}

	return decoded, nil
}

func randomBytes(size int) []byte {
	rando := make([]byte, size)
	rand.Read(rando)
	return rando
}

func RandomPassphrase(size int) ([]string, error) {
	return BytesToPassphrase(randomBytes(size))
}

func BytesToString(input []byte) (string, error) {
	strslice, err := BytesToPassphrase(input)
	if err != nil {
		return "", err
	}

	return strings.Join(strslice, " "), nil
}

func StringToBytes(input string) ([]byte, error) {
	words := strings.Split(input, " ")
	return PassphraseToBytes(words)
}

func RandomString() (string, error) {
	return BytesToString(randomBytes(12))
}
