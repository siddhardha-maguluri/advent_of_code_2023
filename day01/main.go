package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func changeWordToDigit(line string, wordsAndIntegersMap map[string]string) string {
	processedString := line

	for key, value := range wordsAndIntegersMap {
		processedString = strings.Replace(processedString, key, value, -1)
	}

	return processedString
}

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
	wordsAndIntegersMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = changeWordToDigit(line, wordsAndIntegersMap)
		totalSum += getCalibrationValueFromString(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of all of the calibration values: ", totalSum)

}
