// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package rcv

import "testing"

func NewPlayer() Player {
	return Player{
		name:      "batman",
		health:    90,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 51,
	}
}

func TestMaxHealthAndEnergy(t *testing.T) {

	p := NewPlayer()
	p.addHealth(100)

	if p.health > 100 {
		t.Fatalf("Health is beyond allowed, want %v, got %v", p.maxHealth, p.health)
	}

	p.addEnergy(50)

	if p.energy > 100 {
		t.Fatalf("Energy is beyond allowed, want %v, got %v", p.maxEnergy, p.energy)
	}
}

func TestCantBeBelowZero(t *testing.T) {
	p := NewPlayer()

	p.applyDamage(300)

	if p.health > 100 {
		t.Fatalf("Health can't below zero, got %v", p.health)
	}

	p.consumeEnergy(300)

	if p.energy > 100 {
		t.Fatalf("Energy can't below zero, got %v", p.energy)
	}

}
