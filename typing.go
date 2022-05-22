package main

import (
	"fmt"
	"os" // for os interaction
	"bufio" // for scanner
	"strconv"
	"math/rand"
)

func main() {
	num, err := total()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(num)

	//words := getWords()
}

func listOfWords () ([] string, error) {
	f, err := os.Open("words.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
		//fmt.Println(scanner.Text())
	}
	return words, nil
}

func total () (int64, error) {
	var response string
	fmt.Print("Enter the number of words: ")
	fmt.Scanln(&response)
	num, err := strconv.ParseInt(response, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func randomWords (words []string, num int64) ([]string) {
	random := make([]string, num)
	var i int64
	for i = 0; i < num; i++ {
		random[i] = words[rand.Intn(len(words))]
	}
	return random
}