package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttack := PLAYER_ATTACK_MIN_DMG
	maxAttack := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttack = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttack = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}

	dmgValue := generateRandBetween(minAttack, maxAttack)
	currentMonsterHealth -= dmgValue
	return dmgValue
}

func HealPlayer() int {
	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)

	healthDiff := PLAYER_HEALTH - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
	}

	return healValue
}

func AttackPlayer() int {
	dmgValue := generateRandBetween(MONSTER_ATTACK_MIN_DMG, MONSTER_ATTACK_MAX_DMG)
	currentPlayerHealth -= dmgValue
	return dmgValue
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
