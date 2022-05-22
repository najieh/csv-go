package main

import (
	"fmt"
	"os" // for os interaction
	"bufio" // for scanner
	"strconv"
	"math/rand"
	"math"
	"time"
)

func main() {
	num, err := total()
	checkError(err)
	
	words, err := listOfWords()
	checkError(err)

	random := randomWords(words, num)

	accuracy, time := startGame(random)

	//fmt.Println("Accuracy is", accuracy, " and duration is", time)

	var wpm float64 = float64(accuracy) / 100
	wpm *= float64(num)

	var multiply float64 = 0
	if time < 60 {
		multiply = float64(60) / float64(time)
	} else {
		multiply = float64(time) / float64(60)
	}

	wpm *= multiply

	finalWpm := int(math.Round(wpm))

	writeToCsv(strconv.Itoa(finalWpm), strconv.Itoa(accuracy))
}

func listOfWords () ([] string, error) {
	f, err := os.Open("../words.txt")
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
	fmt.Print("Enter the number of words (greater than 10): ")
	fmt.Scanln(&response)
	num, err := strconv.ParseInt(response, 10, 64)
	if err != nil || num <= 1 {
		err = fmt.Errorf("Invalid number or failed read")
		return 0, err
	}
	
	return num, nil
}

func randomWords (words []string, num int64) ([]string) {
	random := make([]string, num)
	var i int64
	rand.Seed(time.Now().UnixNano())
	for i = 0; i < num; i++ {
		random[i] = words[rand.Intn(len(words))]
	}
	return random
}

func checkError (err error) () {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startGame (random [] string) (int, float64) {
	arr := make([]string, len(random))

	for i := 0; i < 3; i++ {
		fmt.Printf("%d...\n", i)
		time.Sleep(time.Second)
	}
	fmt.Println("go!\n")

	start := time.Now()
	for i, word := range random {
		fmt.Print(word)
		if i + 1 < len(random) {
			fmt.Print(" ", random[i+1])
			if i + 2 < len(random) {
				fmt.Print(" ", random[i+2])
				if i + 3 < len(random) {
					fmt.Print(" ", random[i+3])
				}
			}
		}
		fmt.Print("\n")
		fmt.Scanln(&arr[i])
		fmt.Print("\n")
	}

	duration := time.Since(start)

	return processResults(arr, random), duration.Seconds()
}

func processResults(arr [] string, random [] string) (int) {
	var sum float64 = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == random[i] {
			sum += 1
		}
	}

	sum /= float64(len(arr))
	sum *= 100

	return int(math.Round(sum))
}