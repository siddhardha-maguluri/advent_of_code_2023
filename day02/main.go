package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"
	var originalGameData = map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	idsTotal := 0
	powerSetOfCubes := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineData := strings.Split(line, ":")
		gameNumber, _ := strconv.Atoi(strings.Split(lineData[0], " ")[1])
		gameData := strings.Split(lineData[1], ";")

		redCubesCount := 0
		blueCubesCount := 0
		greenCubesCount := 0

		for _, data := range gameData {
			data := strings.Trim(data, " ")
			for _, val := range strings.Split(data, ",") {
				val := strings.Trim(val, " ")
				currentCount, _ := strconv.Atoi(strings.Split(val, " ")[0])
				currentCube := strings.Split(val, " ")[1]

				if strings.EqualFold(currentCube, "red") {
					redCubesCount = int(math.Max(float64(redCubesCount), float64(currentCount)))
				} else if strings.EqualFold(currentCube, "blue") {
					blueCubesCount = int(math.Max(float64(blueCubesCount), float64(currentCount)))
				} else {
					greenCubesCount = int(math.Max(float64(greenCubesCount), float64(currentCount)))
				}
			}
		}

		powerSetOfCubes += redCubesCount * blueCubesCount * greenCubesCount

		if (redCubesCount <= originalGameData["red"]) && (blueCubesCount <= originalGameData["blue"]) && (greenCubesCount <= originalGameData["green"]) {
			idsTotal += gameNumber
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("game ids total sum: ", idsTotal)
	fmt.Println("sum of the power of these sets", powerSetOfCubes)

}
