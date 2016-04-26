package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func firstIsUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}

func dropPunctuation(r rune) rune {
	if unicode.IsSpace(r) || unicode.IsLetter(r) || unicode.IsNumber(r) {
		return r
	}
	return -1
}

type byWackyAlpha []string

func (a byWackyAlpha) Len() int      { return len(a) }
func (a byWackyAlpha) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byWackyAlpha) Less(i, j int) bool {
	if firstIsUpper(a[i]) && firstIsUpper(a[j]) {
		return a[i] > a[j]
	}
	if !firstIsUpper(a[i]) && !firstIsUpper(a[j]) {
		return a[i] < a[j]
	}
	return firstIsUpper(a[j])
}

func main() {
	input := "A Person, who never Made a Mistake never tried anything New."
	alphaOnly := strings.Map(dropPunctuation, input) // HLMap
	words := strings.Split(alphaOnly, " ")
	sort.Sort(byWackyAlpha(words)) // HLSort
	fmt.Println(strings.Join(words, " "))
}
