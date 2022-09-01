package main

import rl "github.com/gen2brain/raylib-go/raylib"

const GridTilesMovedPerPlayerMovementStep = 1

var (
	PlayerGridLowerBounds = rl.NewVector3(-5.5, 0, -5.5)
	PlayerGridUpperBounds = rl.NewVector3(5.5, 0, 5.5)
)

type Player struct {
	Name         string
	Position     rl.Vector3
	MoveCooldown FrameCooldown
}

func (player *Player) Update(grid *Grid) {
	if player.MoveCooldown.Tick() {
		return
	}

	movementVec := rl.Vector3{}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		movementVec.X += GridTilesMovedPerPlayerMovementStep
		player.MoveCooldown.Active = true
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		movementVec.X -= GridTilesMovedPerPlayerMovementStep
		player.MoveCooldown.Active = true
	}

	// don't allow diagonal movement. -> TODO: movement in X axis currently always "wins"
	if movementVec.X == 0 {
		if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
			movementVec.Z += GridTilesMovedPerPlayerMovementStep
			player.MoveCooldown.Active = true
		}
		if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
			movementVec.Z -= GridTilesMovedPerPlayerMovementStep
			player.MoveCooldown.Active = true
		}
	}

	resultPos := rl.Vector3Add(player.Position, movementVec)
	resultPos.X = rl.Clamp(resultPos.X, PlayerGridLowerBounds.X, PlayerGridUpperBounds.X)
	resultPos.Z = rl.Clamp(resultPos.Z, PlayerGridLowerBounds.Z, PlayerGridUpperBounds.Z)

	//check if new pos is free on the grid and move there if it is
	if resultPos != player.Position && !grid.Check(GetGridCoordinatesFromPlayerPosition(resultPos)) {
		player.MoveCooldown.Active = true
		player.Position = resultPos
	}
}
