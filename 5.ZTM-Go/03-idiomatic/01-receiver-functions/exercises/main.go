//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth int
	energy            int
	maxEnergy         int
}

func (p *Player) addHealth(amount int) {
	p.health += amount

	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}

	fmt.Println(p.name, "Add", amount, "health ->", p.health)
}

func (p *Player) applyDamage(amount int) {
	if p.health-amount > p.health {
		p.health = 0
	} else {
		p.health -= amount
	}

	fmt.Println(p.name, "Damage", amount, "->", p.health)
}

func (p *Player) addEnergy(amount int) {

	p.energy += amount

	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}

	fmt.Println(p.name, "Add", amount, "energy ->", p.energy)
}

func (p *Player) consumeEnergy(amount int) {
	if p.energy-amount < p.energy {
		p.energy = 0
	} else {
		p.energy -= amount
	}

	fmt.Println(p.name, "Consume", amount, "energy ->", p.energy)
}

func main() {

	p := Player{
		name:      "batman",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	p.applyDamage(99)
	p.addHealth(10)
	p.consumeEnergy(20)
	p.addEnergy(99)

	p.consumeEnergy(9999)
}
