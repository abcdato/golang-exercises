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
	health, maxHealth uint
	energy, maxEnergy uint
}

func (p *Player) addHealth(num uint) {
	p.health += num
	if p.health >= p.maxHealth {
		p.health = p.maxHealth
	}
}

func (p *Player) damagePlayer(amount uint) {
	if p.health-amount > p.health {
		p.health = 0
	} else {
		p.health -= amount
	}
}

func (p *Player) modifyEnergy(num uint) {
	p.energy += num
	if p.energy >= p.maxEnergy {
		p.energy = p.maxEnergy
	}
}

func (p *Player) consumeEnergy(amount uint) {
	if p.energy-amount > p.energy {
		p.energy = 0
	} else {
		p.energy -= amount
	}
}

func main() {
	p := Player{
		name:      "John",
		health:    50,
		maxHealth: 100,
		energy:    80,
		maxEnergy: 100,
	}
	fmt.Println(p)

	p.addHealth(10)
	fmt.Printf("Player health: %v\nPLayer Max Health: %v\n", p.health, p.maxHealth)

	p.damagePlayer(20)
	fmt.Printf("Player health: %v\nPLayer Max Health: %v\n", p.health, p.maxHealth)

	p.modifyEnergy(5)
	fmt.Printf("Player energy: %v\nPLayer Max Energy: %v\n", p.energy, p.maxEnergy)

	p.consumeEnergy(30)
	fmt.Printf("Player energy: %v\nPLayer Max Energy: %v\n", p.energy, p.maxEnergy)
}
