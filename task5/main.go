package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	arrayPtr := flag.String("array", "[1, 2, 3]", "a string")
	flag.Parse()

	pattern := `^\[\d+(, \d+)*\]$`

	matched, err := regexp.MatchString(pattern, *arrayPtr)
	if err != nil {
		log.Fatal(err.Error())
	}

	if !matched {
		log.Fatal("The input is not in the expected format")
	}

	array := extractArray(*arrayPtr)
	fmt.Println(array)

	fmt.Println("After deduplicating the array")
	array = deduplicateArray(array)
	fmt.Println(array)

	fmt.Println("After sorting the array")
	array = sortArray(array)
	fmt.Println(array)

}

func extractArray(input string) []int {
	pattern := `\d+`
	reg := regexp.MustCompile(pattern)
	matches := reg.FindAllString(input, -1)

	array := make([]int, len(matches))
	for i, match := range matches {
		el, err := strconv.Atoi(match)
		if err != nil {
			log.Fatal(err.Error())
		}
		array[i] = el
	}
	return array
}

func deduplicateArray(array []int) []int {
	for currIndex, el := range array {
		for i := currIndex + 1; i < len(array); i++ {
			if array[i] == el {
				array = remove(array, i)
			}
		}
	}
	return array
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func sortArray(array []int) []int {
	sortable := sort.IntSlice(array[:])
	sort.Sort(sortable)
	return array
}
