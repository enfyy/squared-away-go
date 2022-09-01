package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var (
	playerStartPosition = rl.Vector3{X: -5.5, Y: 0.5, Z: -5.5}
	enemy1StartPosition = rl.Vector3{X: 5, Y: 1, Z: -5}
	enemy2StartPosition = rl.Vector3{X: 5, Y: 1, Z: 5}
	enemy3StartPosition = rl.Vector3{X: -5, Y: 1, Z: 5}
)

func main() {

	camera := rl.Camera{
		Position: rl.Vector3{Y: 11.0, Z: 12.5},
		Up:       rl.Vector3{Y: 1.0}, Fovy: 45.0,
		Projection: rl.CameraPerspective,
		Target:     rl.Vector3{Z: 0.75},
	}

	grid := NewGrid(enemy1StartPosition, enemy2StartPosition, enemy3StartPosition)

	player := Player{
		Name:         "Player",
		Position:     playerStartPosition,
		MoveCooldown: FrameCooldown{Max: 10},
	}

	enemies := [3]Enemy{
		{
			Position:     enemy1StartPosition,
			Color:        rl.Blue,
			OutlineColor: rl.Lime,
			MoveCooldown: FrameCooldown{Max: 30},
		},
		{
			Position:     enemy2StartPosition,
			Color:        rl.Yellow,
			OutlineColor: rl.Orange,
			MoveCooldown: FrameCooldown{Max: 30},
		},
		{
			Position:     enemy3StartPosition,
			Color:        rl.Green,
			OutlineColor: rl.SkyBlue,
			MoveCooldown: FrameCooldown{Max: 30},
		},
	}

	game := Game{
		GameState: NotStarted,
		Camera:    camera,
		Timer:     Timer{DurationInSeconds: 30},
		Player:    player,
		Enemies:   enemies,
		IsWon:     false,
		Countdown: Timer{DurationInSeconds: 4},
		Grid:      &grid,
	}

	// Set up raylib.
	rl.InitWindow(screenWidth, screenHeight, "FUN WITH GOLANG & RAYLIB")
	rl.SetCameraMode(game.Camera, rl.CameraFree)
	rl.SetTargetFPS(60)

	// seeding the random
	rand.Seed(time.Now().UnixNano())

	game.Loop()
}
