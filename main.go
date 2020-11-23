package leetspeak_go_ru

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strings"
)

var Alphabet map[rune][]string

func init() {
	Alphabet = make(map[rune][]string)
	var preAlphabet map[string]string
	alphFile, err := ioutil.ReadFile("./alphabet.json")
	if err != nil {
		println("Can't find alphabet.json! Switch to internal alphabet!")
		alphFile = InternalAlphabet
	}
	err = json.Unmarshal(alphFile, &preAlphabet)
	if err != nil {
		println("Panic!  Can't unmarshal alphabet.json | " + err.Error())
		panic(err.Error())
	}
	for key, val := range preAlphabet {
		l := []rune(key)[0]
		splited := []string{}
		for _, letter := range strings.Split(val, " ") {
			splited = append(splited, letter)
		}
		Alphabet[l] = splited
	}

}

func Translate(text string) string {
	text = strings.ToLower(text)
	runes := []rune(strings.ToLower(text))
	length := len(runes)
	leet := make([]string, 0, length)
	for _, letter := range runes {
		if alt, ok := Alphabet[letter]; ok {
			leet = append(leet, alt[rand.Intn(len(alt))])
		} else {
			leet = append(leet, string(letter))
		}
	}
	return strings.Join(leet, "")
}
