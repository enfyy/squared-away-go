package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState int

const (
	NotStarted GameState = iota
	OnGoing
	Finished
)

type Game struct {
	GameState GameState
	Camera    rl.Camera3D
	Timer     Timer
	Player    Player
	Enemies   [3]Enemy
	IsWon     bool
	Countdown Timer
	Grid      *Grid
}

// Loop - The main game loop.
func (game *Game) Loop() {
	for !rl.WindowShouldClose() {

		// Update
		//--------------------------------------------------------------------------------------------------------------
		rl.UpdateCamera(&game.Camera)
		game.Update()

		// Draw
		//--------------------------------------------------------------------------------------------------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(game.Camera) //---3D Section---------------------------------------------------------------------

		// Draw Player
		rl.DrawCube(game.Player.Position, 1.0, 1.0, 1.0, rl.Red)
		rl.DrawCubeWires(game.Player.Position, 1.0, 1.0, 1.0, rl.Maroon)

		//Draw Enemies
		for _, enemy := range game.Enemies {
			rl.DrawCube(enemy.Position, 2.0, 2.0, 2.0, enemy.Color)
			rl.DrawCubeWires(enemy.Position, 2.0, 2.0, 2.0, enemy.OutlineColor)
		}

		rl.DrawGrid(GridSlices, GridSpacing)
		rl.EndMode3D() //-----------------------------------------------------------------------------------------------

		game.DrawTextElements()
		rl.EndDrawing()
		//--------------------------------------------------------------------------------------------------------------
	}

	rl.CloseWindow()
}

func (game *Game) Update() {
	if game.Timer.Expired() {
		game.Timer.ElapsedSeconds = game.Timer.DurationInSeconds
		game.IsWon = true
		game.GameState = Finished
	}

	switch game.GameState {
	case NotStarted:
		if rl.IsKeyPressed(rl.KeySpace) {
			game.Countdown.Started = true
		}
		game.Countdown.Tick(rl.GetFrameTime())
		if game.Countdown.Expired() {
			game.Start()
			game.GameState = OnGoing
		}
		break

	case OnGoing:
		game.Timer.Tick(rl.GetFrameTime())

		game.Player.Update(game.Grid)

		// Update enemies
		for i, _ := range game.Enemies {
			// note to self: using the value is actually a copy and thus making changes on it does not affect anything
			//enemy.Update(game.Player.Position)
			game.Enemies[i].Update(game.Player.Position, game.Grid)
		}

		if game.CheckLoseCondition() {
			game.GameState = Finished
		}
		break

	case Finished:
		if rl.IsKeyPressed(rl.KeySpace) {
			game.Countdown.Reset()
			game.Start()
			game.GameState = NotStarted
			game.Countdown.Started = true
		}
		break

	default:
		panic("Unhandled GameState encountered while updating game data.")
	}
}

func (game *Game) DrawTextElements() {
	// Calculate player screen space position (with a little offset to be in top)
	playerScreenPosition := rl.GetWorldToScreen(rl.Vector3{X: game.Player.Position.X, Y: game.Player.Position.Y + 1.5, Z: game.Player.Position.Z}, game.Camera)
	halfMeasuredPlayerText := rl.MeasureText(game.Player.Name, 20) / 2
	rl.DrawText(game.Player.Name, int32(playerScreenPosition.X)-halfMeasuredPlayerText, int32(playerScreenPosition.Y), 20, rl.Black)

	switch game.GameState {
	case NotStarted:
		text := "Press Space to start!"
		if game.Countdown.Started && !game.Countdown.Expired() {
			num := int(game.Countdown.DurationInSeconds) - int(game.Countdown.ElapsedSeconds) - 1
			text = fmt.Sprintf("%d", num)
			if num == 0 {
				text = "Start!"
			}
		}
		rl.DrawText(text, (screenWidth-rl.MeasureText(text, 40))/2, screenHeight/2, 40, rl.Gray)
		break

	case OnGoing:
		timerText := fmt.Sprintf("%d", int(game.Timer.DurationInSeconds)-int(game.Timer.ElapsedSeconds))
		rl.DrawText(timerText, (screenWidth-rl.MeasureText(timerText, 50))/2, 20, 50, rl.SkyBlue)
		break

	case Finished:
		startOverText := "Press Space to start over!"
		timerText := "Game over"
		if game.IsWon {
			timerText = "You win!"
		}

		rl.DrawText(timerText, (screenWidth-rl.MeasureText(timerText, 40))/2, screenHeight/2, 40, rl.Black)
		rl.DrawText(startOverText, (screenWidth-rl.MeasureText(startOverText, 30))/2, 25, 30, rl.Gray)
		break

	default:
		panic("Unhandled GameState encountered while drawing Text elements.")
	}
}

func (game *Game) CheckLoseCondition() bool {
	// the player loses when he gets caught by one of the big cubes.
	return game.Grid.Check(GetGridCoordinatesFromPlayerPosition(game.Player.Position))
}

func (game *Game) Start() {
	game.IsWon = false
	game.Player.Position = playerStartPosition
	game.Enemies[0].Position = enemy1StartPosition
	game.Enemies[1].Position = enemy2StartPosition
	game.Enemies[2].Position = enemy3StartPosition
	game.Timer.Reset()
	game.Timer.Started = true
	grid := NewGrid(enemy1StartPosition, enemy2StartPosition, enemy3StartPosition)
	game.Grid = &grid
}
