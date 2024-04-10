package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	asciiFigure := figure.NewFigure("Monster Slayer Game", "", true)
	asciiFigure.Print()
	fmt.Println("Starting a new game ...")
	fmt.Println("Good luck")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack agains monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	asciiFigure := figure.NewColorFigure("Game Over!", "", "red", true)
	fmt.Println("-------------------------")
	asciiFigure.Print()
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {

	// go build
	expPath, err := os.Executable()

	if err != nil {
		fmt.Println("Writing log file failed. Exiting!")
		return
	}

	expPath = filepath.Dir(expPath)

	file, err := os.Create(expPath + "/gamelog.txt")

	// go run
	// file, err := os.Create("gamelog.txt")

	if err != nil {
		fmt.Println("Saving a log file failed. Exiting.")
		return
	}

	for index, round := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                round.Action,
			"Player Attack Damage":  fmt.Sprint(round.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(round.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(round.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(round.PlayerHealth),
			"Monster Health":        fmt.Sprint(round.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString((logLine))

		if err != nil {
			fmt.Println("Writing into log file failed. Exiting.")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log")

}
