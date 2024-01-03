package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

)

type Game struct {
	ID int
	// Highest value pulled from the bag.
	Red int
	Blue int
	Green int
}

func loadGames(filename string) []Game {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't read from file %s", filename)
	}
	defer file.Close()

	games := []Game{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := parseGame(scanner.Text())
		games = append(games, game)
		fmt.Printf("Game %d: Red = %d, Blue = %d, Green = %d\n", game.ID, game.Red, game.Blue, game.Green)
	}

	return games
}

func parseGame(gamestr string) Game {
	game := Game{}

	gameData := strings.Split(gamestr, ": ")
	game.ID, _ = strconv.Atoi(strings.Split(gameData[0], " ")[1])

	gameIterations := strings.Split(gameData[1], "; ")
	for _, iteration := range gameIterations {
		colors := strings.Split(iteration, ", ")

		for _, cubes := range colors {

			cubeData := strings.Split(cubes, " ")
			count, _ := strconv.Atoi(cubeData[0])
			
			switch cubeData[1] {
				case "red":
					if (count > game.Red) {
						game.Red = count
					}
				case "blue":
					if (count > game.Blue) {
						game.Blue = count
					}
				case "green":
					if (count > game.Green) {
						game.Green = count
					}
			}
		}
	}
	return game
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing input file in arguments")
		os.Exit(1)
	}
	inputFile := os.Args[1]

	games := loadGames(inputFile)
	var gameIDCount int
	var gamePowerSetSum int
	for _, game := range games {
		if game.Red <= 12 && game.Blue <= 14 && game.Green <= 13 {
			gameIDCount += game.ID 
		}
		gamePowerSetSum += (game.Red * game.Blue * game.Green)
	}

	fmt.Printf("gameIDCount: %d\n", gameIDCount)
	fmt.Printf("gamePowerSetSum: %d\n", gamePowerSetSum)
}

