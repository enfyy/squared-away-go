package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

const (
	GridTilesMovedPerEnemyMovementStep   = 2
	ChanceOfAIMoving                     = 90
	ChanceOfAIMovingInASensibleDirection = 70
)

var (
	EnemyGridLowerBounds = rl.NewVector3(-5, 0, -5)
	EnemyGridUpperBounds = rl.NewVector3(5, 0, 5)
)

type Enemy struct {
	Position     rl.Vector3
	Color        rl.Color
	OutlineColor rl.Color
	MoveCooldown FrameCooldown
}

func (enemy *Enemy) Update(playerPosition rl.Vector3, grid *Grid) {
	if rand.Intn(101) > ChanceOfAIMoving || enemy.MoveCooldown.Tick() {
		return
	}

	// find path to player
	xDiff := playerPosition.X - enemy.Position.X
	zDiff := playerPosition.Z - enemy.Position.Z
	xDir := xDiff / Abs(xDiff) // TODO: coincidentally cant be 0 but maybe check anyways
	zDir := zDiff / Abs(zDiff)
	moveX := Abs(xDiff) > .5
	moveZ := Abs(zDiff) > .5
	movementVec := rl.Vector3{}

	// move in a sensible direction
	if rand.Intn(101) <= ChanceOfAIMovingInASensibleDirection {

		// move in x dir
		if moveX && moveZ {
			// roll for direction
			if rand.Intn(2) == 1 {
				// correct x dir
				movementVec.X = xDir * GridTilesMovedPerEnemyMovementStep
			} else {
				movementVec.Z = zDir * GridTilesMovedPerEnemyMovementStep
			}

		} else if moveX {
			movementVec.X = xDir * GridTilesMovedPerEnemyMovementStep
		} else if moveZ {
			movementVec.Z = zDir * GridTilesMovedPerEnemyMovementStep
		}
	} else {
		// move in a opposite direction
		if rand.Intn(2) == 1 {
			// correct x dir
			movementVec.X = -xDir * GridTilesMovedPerEnemyMovementStep
		} else {
			movementVec.Z = -zDir * GridTilesMovedPerEnemyMovementStep
		}
	}

	resultPos := rl.Vector3Add(enemy.Position, movementVec)
	resultPos.X = rl.Clamp(resultPos.X, EnemyGridLowerBounds.X, EnemyGridUpperBounds.X)
	resultPos.Z = rl.Clamp(resultPos.Z, EnemyGridLowerBounds.Z, EnemyGridUpperBounds.Z)

	//check if new pos is free on the grid and move there if it is
	if resultPos != enemy.Position && !grid.CheckAny(GetGridCoordinatesFromEnemyPosition(resultPos)) {
		enemy.MoveCooldown.Active = true
		grid.MoveEnemy(enemy.Position, resultPos)
		enemy.Position = resultPos
	}
}
