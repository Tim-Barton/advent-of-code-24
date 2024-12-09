package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func forwards(input string) int {
	testWord := "XMAS"
	return strings.Count(input, testWord)
}

func backwards(input string) int {
	testWord := "SAMX"
	return strings.Count(input, testWord)
}

func upwards(input []string, wait *sync.WaitGroup, output chan string) {
	//using x & y cardinality likes it's a maths plot cos that works better in my head for keeping order intact
	for y := 3; y < len(input); y++ {
		for x := range input[y] {
			word := string(input[y][x]) + string(input[y-1][x]) + string(input[y-2][x]) + string(input[y-3][x])
			output <- word
		}
	}
	wait.Done()
}

func downwards(input []string, wait *sync.WaitGroup, output chan string) {
	//using x & y cardinality likes it's a maths plot cos that works better in my head for keeping order intact
	for y := 0; y < len(input)-3; y++ {
		for x := range input[y] {
			word := string(input[y][x]) + string(input[y+1][x]) + string(input[y+2][x]) + string(input[y+3][x])
			output <- word
		}
	}
	wait.Done()
}

func upleft(input []string, wait *sync.WaitGroup, output chan string) {
	//using x & y cardinality likes it's a maths plot cos that works better in my head for keeping order intact
	for y := 3; y < len(input); y++ {
		for x := 3; x < len(input[y]); x++ {
			word := string(input[y][x]) + string(input[y-1][x-1]) + string(input[y-2][x-2]) + string(input[y-3][x-3])
			output <- word
		}
	}
	wait.Done()
}

func upright(input []string, wait *sync.WaitGroup, output chan string) {
	//using x & y cardinality likes it's a maths plot cos that works better in my head for keeping order intact
	for y := 3; y < len(input); y++ {
		for x := 0; x < len(input[y])-3; x++ {
			word := string(input[y][x]) + string(input[y-1][x+1]) + string(input[y-2][x+2]) + string(input[y-3][x+3])
			output <- word
		}
	}
	wait.Done()
}

func downleft(input []string, wait *sync.WaitGroup, output chan string) {
	//using x & y cardinality likes it's a maths plot cos that works better in my head for keeping order intact
	for y := 0; y < len(input)-3; y++ {
		for x := 3; x < len(input[y]); x++ {
			word := string(input[y][x]) + string(input[y+1][x-1]) + string(input[y+2][x-2]) + string(input[y+3][x-3])
			output <- word
		}
	}
	wait.Done()
}

func downright(input []string, wait *sync.WaitGroup, output chan string) {
	//using x & y cardinality likes it's a maths plot cos that works better in my head for keeping order intact
	for y := 0; y < len(input)-3; y++ {
		for x := 0; x < len(input[y])-3; x++ {
			word := string(input[y][x]) + string(input[y+1][x+1]) + string(input[y+2][x+2]) + string(input[y+3][x+3])
			output <- word
		}
	}
	wait.Done()
}

func checkInput(input []string) int {
	simplecount := 0
	wait := sync.WaitGroup{}

	wait.Add(1)
	go func() {
		for _, line := range input {
			simplecount += forwards(line)
			simplecount += backwards(line)
		}
		wait.Done()
	}()

	assessChannel := make(chan string)

	wait.Add(6)
	go upwards(input, &wait, assessChannel)
	go downwards(input, &wait, assessChannel)
	go upleft(input, &wait, assessChannel)
	go upright(input, &wait, assessChannel)
	go downleft(input, &wait, assessChannel)
	go downright(input, &wait, assessChannel)

	complexCount := 0
	assess := sync.WaitGroup{}
	assess.Add(1)
	go func() {
		for word := range assessChannel {
			//fmt.Println(word)
			if word == "XMAS" {
				complexCount += 1
			}
		}
		assess.Done()
	}()

	wait.Wait()
	close(assessChannel)
	assess.Wait()
	fmt.Printf("Results. Simple: %d. Complex: %d\n", simplecount, complexCount)
	return simplecount + complexCount
}

func xMasCheck(input []string, wait *sync.WaitGroup, output chan []string) {
	for y := 0; y < len(input)-2; y++ {
		for x := 0; x < len(input[y])-2; x++ {
			word := string(input[y][x]) +
				string(input[y+1][x+1]) +
				string(input[y+2][x+2])
			word2 := string(input[y][x+2]) +
				string(input[y+1][x+1]) +
				string(input[y+2][x])
			output <- []string{word, word2}
		}
	}
	wait.Done()
}

func checkInput2(input []string) int {
	wait := sync.WaitGroup{}

	assessChannel := make(chan []string)

	wait.Add(1)
	go xMasCheck(input, &wait, assessChannel)

	complexCount := 0
	assess := sync.WaitGroup{}
	assess.Add(1)
	go func() {
		for words := range assessChannel {
			if (words[0] == "MAS" || words[0] == "SAM") && (words[1] == "MAS" || words[1] == "SAM") {
				complexCount += 1
			}
		}
		assess.Done()
	}()

	wait.Wait()
	close(assessChannel)
	assess.Wait()
	fmt.Printf("Results. X-Mas: %d\n", complexCount)
	return complexCount
}

func main() {
	input, _ := os.Open("input")
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	data := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	count := checkInput(data)
	count2 := checkInput2(data)

	fmt.Println(count)
	fmt.Println(count2)

}
