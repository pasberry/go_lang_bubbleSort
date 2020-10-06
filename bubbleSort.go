//Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
//The program should print the integers out on one line, in sorted order, from least to greatest.
//Use your favorite search tool to find a description of how the bubble sort algorithm works.
//As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an
//argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.
//A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent
//elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should
//take two arguments, a slice of integers and an index value i which indicates a position in the slice.
//The Swap() function should return nothing, but it should swap the contents of the slice in position i with the
//contents in position i+1.

//Submit your Go program source code.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Please enter a sequence of integers separated by a space to be sorted.")

	inputScanner := setUpConsoleScanner()

	unsortedStrings := getSequence(inputScanner)

	validateList(&unsortedStrings)

	unsortedSequence := convertToListOfInts(&unsortedStrings)

	BubbleSort(&unsortedSequence)

	fmt.Println(unsortedSequence)
}

func BubbleSort(unsortedSequence *[]int) {

	unsorted := true

	for unsorted {

		for index  := 0; index < len(*unsortedSequence) - 1 ; index++ {

			Swap(index , unsortedSequence)
		}

		unsorted = isSequenceUnSorted(unsortedSequence)
	}
}


func Swap(index int, ints *[]int) {

	firstNum := (*ints)[index]
	secondNum := (*ints)[index + 1]

	if firstNum > secondNum {
		(*ints)[index] = secondNum
		(*ints)[index + 1] = firstNum
	}
}

func isSequenceUnSorted(unsortedSequence *[]int) (unsorted bool) {

	unsorted = false

	for index  := 1; index < len(*unsortedSequence); index++ {

		if (*unsortedSequence)[index - 1] > (*unsortedSequence)[index] {
			unsorted = true
			break
		}
	}

	return
}

func convertToListOfInts(list *[]string) (ints []int) {

	ints = make([]int, 0)

	for _, str := range *list {

		num , _ := strconv.ParseInt(str, 10, 64 )

		ints = append(ints , int(num))
	}
	return
}

func setUpConsoleScanner() *bufio.Scanner {

	inputScanner := bufio.NewScanner(os.Stdin)
	return inputScanner
}

func getSequence(inputScanner *bufio.Scanner) []string {
	inputScanner.Scan()
	return strings.Split(strings.TrimSpace(inputScanner.Text()), " ")
}

func validateList(list *[]string) {

	for _,str := range *list {
		_,err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			panic(err)
		}
	}
}