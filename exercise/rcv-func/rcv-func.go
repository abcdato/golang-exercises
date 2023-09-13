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
	energy, maxEnergy int
}

func (p *Player) modifyHealth(num int) {
	p.health += num
	if p.health >= p.maxHealth {
		p.health = p.maxHealth
	}
}

func (p *Player) modifyEnergy(num int) {
	p.energy += num
	if p.energy >= p.maxEnergy {
		p.energy = p.maxEnergy
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

	p.modifyHealth(10)
	fmt.Printf("Player health: %v\nPLayer Max Health: %v\n", p.health, p.maxHealth)

	p.modifyHealth(30)
	fmt.Printf("Player health: %v\nPLayer Max Health: %v\n", p.health, p.maxHealth)

	p.modifyEnergy(-30)
	fmt.Printf("Player energy: %v\nPLayer Max Energy: %v\n", p.energy, p.maxEnergy)
}
