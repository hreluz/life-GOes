package main

import (
	"fmt"

	"github.com/hreluz/monster-slayer-game/actions"
	"github.com/hreluz/monster-slayer-game/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := ""

	for winner == "" {
		executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0
	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	var playerHealth int
	var monsterHealth int

	if userChoice == "ATTACK" {
		actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(true)
	}

	actions.AttackPlayer()

	playerHealth, monsterHealth = actions.GetHealthAmounts()

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame() {
}
