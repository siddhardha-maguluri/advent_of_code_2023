package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func getCalibrationValueFromString(line string) int {

	firstDigitFound, secondDigitFound := false, false
	var firstDigit, secondDigit, numberAsString string
	var calibrationValue int

	for _, char := range line {
		if unicode.IsDigit(char) && !firstDigitFound {
			firstDigitFound = true
			firstDigit = string(char)
		} else if unicode.IsDigit(char) {
			secondDigit = string(char)
			secondDigitFound = true
		}
	}

	if secondDigitFound {
		numberAsString = firstDigit + secondDigit
	} else {
		numberAsString = firstDigit + firstDigit
	}

	calibrationValue, _ = strconv.Atoi(numberAsString)
	return calibrationValue
}

func main() {

	filePath := "input.txt"
	totalSum := 0

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		totalSum += getCalibrationValueFromString(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of all of the calibration values: ", totalSum)

}
