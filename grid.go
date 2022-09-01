package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	GridSlices  = 12
	GridSpacing = 1.0
)

// Grid - 2 Dimensional array, true values indicate occupied spaces.
type Grid [GridSlices][GridSlices]bool

type GridPoint struct {
	X, Y int
}

func NewGrid(enemy1Pos, enemy2Pos, enemy3Pos rl.Vector3) Grid {
	grid := Grid{}
	enemy1Points := GetGridCoordinatesFromEnemyPosition(enemy1Pos)
	enemy2Points := GetGridCoordinatesFromEnemyPosition(enemy2Pos)
	enemy3Points := GetGridCoordinatesFromEnemyPosition(enemy3Pos)

	for i, _ := range enemy1Points {
		enemy1Point := enemy1Points[i]
		enemy2Point := enemy2Points[i]
		enemy3Point := enemy3Points[i]
		grid[enemy1Point.X][enemy1Point.Y] = true
		grid[enemy2Point.X][enemy2Point.Y] = true
		grid[enemy3Point.X][enemy3Point.Y] = true
	}

	return grid
}

func (grid *Grid) Check(point GridPoint) bool {
	return grid[point.X][point.Y]
}

func (grid *Grid) CheckAny(points []GridPoint) bool {
	for _, point := range points {
		if grid.Check(point) {
			return true
		}
	}
	return false
}

func (grid *Grid) MovePlayer(oldPos rl.Vector3, newPos rl.Vector3) {
	oldPoint := GetGridCoordinatesFromPlayerPosition(oldPos)
	newPoint := GetGridCoordinatesFromPlayerPosition(newPos)

	grid[oldPoint.X][oldPoint.Y] = false
	grid[newPoint.X][newPoint.Y] = true
}

func (grid *Grid) MoveEnemy(oldPos rl.Vector3, newPos rl.Vector3) {
	oldPoints := GetGridCoordinatesFromEnemyPosition(oldPos)
	newPoints := GetGridCoordinatesFromEnemyPosition(newPos)

	for _, point := range oldPoints {
		grid[point.X][point.Y] = false
	}
	for _, point := range newPoints {
		grid[point.X][point.Y] = true
	}
}

func GetGridCoordinatesFromPlayerPosition(position rl.Vector3) GridPoint {
	// pos {X: -5.5, Z: -5.5} is [ 0, 0] -> bottom left
	// pos {X:  5.5, Z: -5.5} is [11, 0] -> bottom right
	// pos {X: -5.5, Z:  5.5} is [0 ,11] -> top left
	// pos {X:  5.5, Z:  5.5} is [11,11] -> top right
	// (-5.5 - 5.5) + 11 = 0
	// (5.5 - 5.5) +  11 = 12
	// (0.5 - 5.5) + 12 = 7
	return GridPoint{int((position.X - 5.5) + GridSlices - 1), int((position.Z - 5.5) + GridSlices - 1)}
}

func GetGridCoordinatesFromEnemyPosition(position rl.Vector3) []GridPoint {
	// Enemy positions are on top of nodes instead of the center of the tiles way, and they occupy 4 tiles at once
	// pos {X: -5, Z: -5} is [0,0],[1,0],[0,1],[1,1] -> top left
	// -5 -3 -1 1 3 5 -> allowed values
	xTopLeft := int((position.X - 6) + GridSlices - 1)
	zTopLeft := int((position.Z - 6) + GridSlices - 1)

	return []GridPoint{
		{xTopLeft, zTopLeft},
		{xTopLeft + 1, zTopLeft},
		{xTopLeft, zTopLeft + 1},
		{xTopLeft + 1, zTopLeft + 1},
	}
}
