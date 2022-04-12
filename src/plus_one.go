package src

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// Problem taken from https://dojopuzzles.com/problems/vai-um/

func plusOne(numberOne string, numberTwo string) {
	runeOne := []rune(numberOne)
	runeTwo := []rune(numberTwo)
	sizeOne := len(runeOne)
	sizeTwo := len(runeTwo)
	lowest := sizeTwo
	countOne := sizeOne - 1
	countTwo := sizeTwo - 1
	plusOneCount := 0

	if sizeOne < sizeTwo {
		lowest = sizeOne
	}

	for i := 0; i < lowest; i++ {
		intOne, _ := strconv.Atoi(string(runeOne[countOne]))
		intTwo, _ := strconv.Atoi(string(runeTwo[countTwo]))

		if intOne+intTwo > 9 {
			plusOneCount++
		}

		countOne--
		countTwo--
	}

	fmt.Printf("Number of plus ones: %d\n", plusOneCount)
}

var plusOneCmd = &cobra.Command{
	Use:   "PlusOne NumberOne NumberTwo",
	Short: "Solve plus one problem",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("invalid number of arguments, please provide two integers")
		}

		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("invalid number one")
		}

		_, err = strconv.Atoi(args[1])
		if err != nil {
			return errors.New("invalid number two")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		plusOne(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(plusOneCmd)
}
