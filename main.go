package main

import (
	"fmt"
	"github.com/MnO2/go-pdqsort"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	runSorting("smallrandom")
	runSorting("random")
	runSorting("bigrandom")
	runSorting("asc")
	runSorting("desc")
	runSorting("equal")
}

func runSorting(sortingType string) {
	length := 1
	for length < 1000000000 {
		fmt.Println("Sorting " + strconv.Itoa(length) + " numbers, " + sortingType)
		numbers, numbers2, numbers3, numbers4, numbers5 := createSlicesOfLength(length, sortingType)
		if length <= 100000 {
			start := time.Now()
			numbers = insertionSort(numbers)
			fmt.Println("Insertion sort: " + time.Since(start).String())
			start = time.Now()
			numbers2 = selectionSort(numbers2)
			fmt.Println("Selection sort: " + time.Since(start).String())
		}
		start := time.Now()
		numbers3 = standardGoQuickSort(numbers3)
		quickSortDuration := time.Since(start)
		fmt.Println("Quick sort:     " + quickSortDuration.String())

		start = time.Now()
		numbers2 = pdqSort(numbers5)
		fmt.Println("PDQ sort:       " + time.Since(start).String())

		start = time.Now()
		numbers4 = pjSort(numbers4)
		pjSortDuration := time.Since(start)
		fmt.Println("PJ sort:        " + pjSortDuration.String())
		compareQuickSortAndPjSortNumbers(numbers3, numbers4)
		fmt.Println("========================================")
		length *= 10
	}
}

func pdqSort(numbers []int) []int {
	pdqsort.Ints(numbers)
	return numbers
}

func compareQuickSortAndPjSortNumbers(numbers3 []int, numbers4 []int) {
	for position, number := range numbers3 {
		if number != numbers4[position] {
			fmt.Println("PJ sort numbers not matching quicksort")
			os.Exit(-1)
		}
	}
}

func createSlicesOfLength(length int, sortingType string) ([]int, []int, []int, []int, []int) {
	rand.Seed(time.Now().UTC().UnixNano())
	var numbers, numbers2, numbers3, numbers4, numbers5 []int
	for j := 0; j < length; j++ {
		var number int
		if sortingType == "random" {
			number = rand.Intn(length) - length/2
		} else if sortingType == "bigrandom" {
			number = rand.Intn(length*10) - length*10/2
		} else if sortingType == "smallrandom" {
			number = (rand.Intn(length) - length/2) / 10
		} else if sortingType == "asc" {
			number = j
		} else if sortingType == "dec" {
			number = length - j
		} else {
			number = length / 2
		}
		numbers = append(numbers, number)
		numbers2 = append(numbers2, number)
		numbers3 = append(numbers3, number)
		numbers4 = append(numbers4, number)
		numbers5 = append(numbers5, number)
	}
	return numbers, numbers2, numbers3, numbers4, numbers5
}

func standardGoQuickSort(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

func pjSort(numbers []int) []int {
	numbersMap := make(map[int]int, len(numbers))
	min, max, number := 0, 0, 0
	for i := 0; i < len(numbers); i++ {
		number = numbers[i]
		if number > max {
			max = number
		} else if number < min {
			min = number
		}
		numbersMap[number]++
	}
	numbersFullRange := make([]int, max-min+1)
	if min < 0 {
		min = -min
	}
	for key, numberInMap := range numbersMap {
		numbersFullRange[key+min] = numberInMap
	}
	position := 0
	for i := 0; i < len(numbersFullRange); i++ {
		for j := 0; j < numbersFullRange[i]; j++ {
			numbers[position] = i - min
			position++
		}
	}
	return numbers
}

func insertionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i; j > 0 && numbers[j-1] > numbers[j]; j-- {
			numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
		}
	}
	return numbers
}

func selectionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		replace := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[replace] {
				replace = j
			}
		}
		numbers[replace], numbers[i] = numbers[i], numbers[replace]
	}
	return numbers
}
