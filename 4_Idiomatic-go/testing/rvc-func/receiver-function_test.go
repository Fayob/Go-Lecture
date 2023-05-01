package main

import "testing"

func TestHealth(t *testing.T)  {
	player := Player{
		name: "James",
		health: 200,
		maxHealth: 500,
		energy: 150,
		maxEnergy: 400,
	}
	player.addHealth(10000)
	if player.health > player.maxHealth {
		t.Fatalf("Health went over limit: %v, want: %v", player.health, player.maxHealth)
	}
	player.applyDamage(player.health + 1)
	if player.health < 0 {
		t.Fatalf("Health: %v. Minimum: 0", player.health)
	}
	if player.health > player.maxHealth {
		t.Fatalf("Health: %v. Minimum: %v", player.health, player.maxHealth)
	}
}

func TestEnergy(t *testing.T)  {
	player := Player{
		name: "James",
		health: 200,
		maxHealth: 500,
		energy: 150,
		maxEnergy: 400,
	}
	player.addEnergy(10000)
	if player.health > player.maxEnergy {
		t.Fatalf("Energy went over limit: %v, want: %v", player.energy, player.maxEnergy)
	}
	player.consumeEnergy(player.energy + 1)
	if player.energy < 0 {
		t.Fatalf("Energy: %v. Minimum: 0", player.energy)
	}
	if player.energy > player.maxEnergy {
		t.Fatalf("Energy: %v. Minimum: %v", player.energy, player.maxEnergy)
	}
}
