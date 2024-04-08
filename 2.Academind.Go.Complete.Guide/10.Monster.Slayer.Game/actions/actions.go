package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = 100
var currentPlayerHealth = 100

func AttackMonster(isSpecialAttack bool) {
	minAttack := 5
	maxAttack := 10

	if isSpecialAttack {
		minAttack = 10
		maxAttack = 20
	}

	dmgValue := generateRandBetween(minAttack, maxAttack)
	currentMonsterHealth -= dmgValue
}

func HealPlayer() {
	minHealValue := 10

	maxHealValue := 20

	healValue := generateRandBetween(minHealValue, maxHealValue)

	healthDiff := 100 - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = 100
	}
}

func AttackPlayer() {
	minAttack := 9
	maxAttack := 13

	dmgValue := generateRandBetween(minAttack, maxAttack)
	currentPlayerHealth -= dmgValue
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
