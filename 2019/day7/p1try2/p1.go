package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fahlmant/intcode"
)

func main() {

	var instructionStrings []string
	max := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructionStrings = strings.Split(line, ",")
	}

	instructions := make([]int, len(instructionStrings))

	for i, v := range instructionStrings {
		instructions[i], _ = strconv.Atoi(v)
	}
	possiblePhaseSettings := buildCombinationsList(0, 5)
	for _, setting := range possiblePhaseSettings {

		bInput := []int{setting[1]}
		cInput := []int{setting[2]}
		dInput := []int{setting[3]}
		eInput := []int{setting[4]}

		computerA := intcode.Computer{PC: 0, Offset: 0, Input: []int{setting[0], 0}, Output: []int{}, Instructions: instructions}
		computerA.RunProgram()
		bInput = append(bInput, computerA.Output...)
		computerB := intcode.Computer{PC: 0, Offset: 0, Input: bInput, Output: []int{}, Instructions: instructions}
		computerB.RunProgram()
		cInput = append(cInput, computerB.Output...)
		computerC := intcode.Computer{PC: 0, Offset: 0, Input: cInput, Output: []int{}, Instructions: instructions}
		computerC.RunProgram()
		dInput = append(dInput, computerC.Output...)
		computerD := intcode.Computer{PC: 0, Offset: 0, Input: dInput, Output: []int{}, Instructions: instructions}
		computerD.RunProgram()
		eInput = append(eInput, computerD.Output...)
		computerE := intcode.Computer{PC: 0, Offset: 0, Input: eInput, Output: []int{}, Instructions: instructions}
		computerE.RunProgram()

		if computerE.Output[0] > max {
			max = computerE.Output[0]
		}
	}

	fmt.Println(max)

}

func buildCombinationsList(low, high int) [][]int {

	rand.Seed(time.Now().Unix())
	var combinationList [][]int

	for len(combinationList) < (factorial(high) - factorial(low)) {
		var combination string
		for len(combination) < 5 {

			if len(combination) < 1 {
				combination = strconv.Itoa(rand.Intn(high-low) + low)
			}

			nextNum := rand.Intn(high-low) + low
			if !strings.Contains(combination, strconv.Itoa(nextNum)) {
				combination = strings.Join([]string{combination, strconv.Itoa(nextNum)}, "")
			}
		}

		for _, item := range combinationList {
			if equal(item, stringToIntArray(combination)) {

			}
		}
		if !sliceInSlice(combinationList, stringToIntArray(combination)) {
			combinationList = append(combinationList, stringToIntArray(combination))
		}

	}

	return combinationList
}

func stringToIntArray(value string) []int {

	var final []int
	for _, v := range value {
		num, _ := strconv.Atoi(string(v))
		final = append(final, num)
	}
	return final
}

func sliceInSlice(slice [][]int, value []int) bool {
	for _, v := range slice {
		if equal(value, v) {
			return true
		}
	}
	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}
