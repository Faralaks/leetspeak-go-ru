package leetspeak_go_ru

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strings"
)

var Alphabet map[rune][]rune

func init() {
	Alphabet = make(map[rune][]rune)
	var preAlphabet map[string]string
	alphFile, err := ioutil.ReadFile("./alphabet.json")
	if err != nil {
		println("Panic!  Can't open alphabet.json | " + err.Error())
		panic(err)
	}
	err = json.Unmarshal(alphFile, &preAlphabet)
	if err != nil {
		println("Panic!  Can't unmarshal alphabet.json | " + err.Error())
		panic(err.Error())
	}
	for key, val := range preAlphabet {
		l := []rune(key)[0]
		splited := []rune{}
		for _, letter := range strings.Split(val, " ") {
			splited = append(splited, []rune(letter)[0])
		}
		Alphabet[l] = splited
	}

}

func Translate(text string) string {
	text = strings.ToLower(text)
	runes := []rune(strings.ToLower(text))
	length := len(runes)
	leet := make([]rune, 0, length)
	for _, letter := range runes {
		if alt, ok := Alphabet[letter]; ok {

			leet = append(leet, alt[rand.Intn(len(alt))])
		} else {

			leet = append(leet, letter)
		}
	}
	return string(leet)
}