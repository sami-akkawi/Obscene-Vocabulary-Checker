package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string

	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error no file with the name '%s' was found. %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words = make(map[string]bool)

	for scanner.Scan() {
		words[strings.ToLower(scanner.Text())] = true
	}

	sentence := readSentence()

	for sentence != "exit" {
		redactSentence(sentence, words)

		sentence = readSentence()
	}

	fmt.Println("Bye!")
}

func redactSentence(sentence string, badWords map[string]bool) {
	words := strings.Fields(sentence)
	for index, value := range words {
		_, ok := badWords[strings.ToLower(value)]
		if ok {
			words[index] = getHash(strlen(value))
		}
	}

	fmt.Println(strings.Join(words, " "))
}

func readSentence() string {
	var sentence string
	fmt.Scan(&sentence)
	return sentence
}

func strlen(str string) int {
	return len([]rune(str))
}

func getHash(length int) string {
	hash := ""
	for i := 0; i < length; i++ {
		hash = hash + "*"
	}

	return hash
}
