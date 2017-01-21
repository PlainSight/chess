package main

import (
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	startGame()
	setup()
}

const (
	GRIDLENGTH    uint = 8
	playerFaction      = 0
)

type piece struct {
	faction uint
	class   uint
}

type position struct {
	x uint
	y uint
}

type move struct {
	from position
	to   position
}

func (move move) Value() int {
	val := 0
	target := getGrid(move.to)
	if target != nil {
		switch target.class {
		case 0:
			val = 1
		case 1:
			val = 3
		case 2:
			val = 3
		case 3:
			val = 5
		case 4:
			val = 9
		case 5:
			val = 20
		}
	}
	return val
}

type moveSet struct {
	moves []move
}

func (moveSet moveSet) Len() int {
	return len(moveSet.moves)
}
func (moveSet moveSet) Less(i, j int) bool {
	// i value
	ival := moveSet.moves[i].Value()

	// j value
	jval := moveSet.moves[j].Value()

	return ival > jval
}
func (moveSet moveSet) Swap(i, j int) {
	temp := moveSet.moves[i]
	moveSet.moves[i] = moveSet.moves[j]
	moveSet.moves[j] = temp
}

var (
	grid     [GRIDLENGTH][GRIDLENGTH]*piece
	selected *position
)

func setGrid(pos position, piece *piece) {
	grid[pos.x][pos.y] = piece
}

func getGrid(pos position) *piece {
	return grid[pos.x][pos.y]
}

func grabPiece(x uint, y uint) {
	if grid[x][y] != nil {
		if grid[x][y].faction == playerFaction {
			selected = &position{x: x, y: y}
		}
	}
}

func movePiece(x uint, y uint) {
	to := position{x: x, y: y}
	if selected != nil {
		if validateMove(move{from: *selected, to: to}) {
			setGrid(to, getGrid(*selected))
			setGrid(*selected, nil)
			selected = nil
			enemyMove()
		}
	}
}

func enemyMove() {
	moveSet := findAllValidMoves(1)
	if len(moveSet.moves) == 0 {
		os.Exit(0)
	}
	sort.Sort(moveSet)

	choices := len(moveSet.moves)
	if choices > 5 {
		choices = 5
	}

	chosen := moveSet.moves[rand.Intn(choices)]
	setGrid(chosen.to, getGrid(chosen.from))
	setGrid(chosen.from, nil)
}

func findAllValidMoves(faction uint) moveSet {

	validMoves := [1024]move{}
	moveCount := uint(0)

	for fx := uint(0); fx < GRIDLENGTH; fx++ {
		for fy := uint(0); fy < GRIDLENGTH; fy++ {
			from := position{x: fx, y: fy}
			fromPiece := getGrid(from)
			if fromPiece != nil && fromPiece.faction == faction {
				for x := uint(0); x < GRIDLENGTH; x++ {
					for y := uint(0); y < GRIDLENGTH; y++ {
						to := position{x: x, y: y}
						move := move{from: from, to: to}
						if validateMove(move) {
							validMoves[moveCount] = move
							moveCount++
						}
					}
				}
			}
		}
	}

	return moveSet{moves: validMoves[:moveCount]}
}

// from and to
func validateMove(move move) bool {

	fx := move.from.x
	fy := move.from.y
	x := move.to.x
	y := move.to.y

	// selected piece
	selectedPiece := grid[fx][fy]
	// target piece
	target := grid[x][y]

	switch selectedPiece.class {
	// pawn
	case 0:
		var canDoubleStep bool
		var advance uint
		if selectedPiece.faction == 0 {
			// white
			advance = fy - y
			canDoubleStep = fy == 6 && grid[x][5] == nil
		} else {
			// black
			advance = y - fy
			canDoubleStep = fy == 1 && grid[x][2] == nil
		}

		if fx == x {
			return target == nil && (advance == 1 || (canDoubleStep && advance == 2))
		} else {
			sidemovement := fx - x
			if sidemovement < 0 {
				sidemovement = -sidemovement
			}
			return advance == 1 && sidemovement == 1 && (target != nil && target.faction != selectedPiece.faction)
		}
	// knight
	case 1:
		dx := int(x) - int(fx)
		if dx < 0 {
			dx = -dx
		}
		dy := int(y) - int(fy)
		if dy < 0 {
			dy = -dy
		}

		if (dx == 2 && dy == 1) || (dy == 2 && dx == 1) {
			if target == nil || (target.faction != selectedPiece.faction) {
				return true
			}
		}
		return false
	// bishop
	case 2:
		dx := int(x) - int(fx)
		vecx := dx
		if dx < 0 {
			dx = -dx
		}

		dy := int(y) - int(fy)
		vecy := dy
		if dy < 0 {
			dy = -dy
		}

		if dx != dy || dx == 0 {
			return false
		}

		incx := vecx / dx
		incy := vecy / dy

		for tx, ty := int(fx)+incx, int(fy)+incy; tx != int(x); tx, ty = tx+incx, ty+incy {
			if grid[tx][ty] != nil {
				return false
			}
		}
		return target == nil || target.faction != selectedPiece.faction
	// rook
	case 3:
		dx := int(x) - int(fx)
		vecx := dx
		if dx < 0 {
			dx = -dx
		}

		dy := int(y) - int(fy)
		vecy := dy
		if dy < 0 {
			dy = -dy
		}

		// check movement is only in one dimension
		if (dx != 0) == (dy != 0) {
			return false
		}

		incx := 0
		if dx != 0 {
			incx = vecx / dx
		}
		incy := 0
		if dy != 0 {
			incy = vecy / dy
		}

		for tx, ty := int(fx)+incx, int(fy)+incy; tx != int(x) || ty != int(y); tx, ty = tx+incx, ty+incy {
			if grid[tx][ty] != nil {
				return false
			}
		}
		return target == nil || target.faction != selectedPiece.faction
	// queen
	case 4:
		dx := int(x) - int(fx)
		vecx := dx
		if dx < 0 {
			dx = -dx
		}

		dy := int(y) - int(fy)
		vecy := dy
		if dy < 0 {
			dy = -dy
		}

		// check movement is valid
		if dx != dy && ((dx != 0) == (dy != 0)) {
			return false
		}

		incx := 0
		if dx != 0 {
			incx = vecx / dx
		}
		incy := 0
		if dy != 0 {
			incy = vecy / dy
		}

		for tx, ty := int(fx)+incx, int(fy)+incy; tx != int(x) || ty != int(y); tx, ty = tx+incx, ty+incy {
			if grid[tx][ty] != nil {
				return false
			}
		}
		return target == nil || target.faction != selectedPiece.faction
	// king
	case 5:
		dx := int(x) - int(fx)
		vecx := dx
		if dx < 0 {
			dx = -dx
		}

		dy := int(y) - int(fy)
		vecy := dy
		if dy < 0 {
			dy = -dy
		}

		// check movement is valid
		if dx > 1 || dy > 1 {
			return false
		}

		incx := 0
		if dx != 0 {
			incx = vecx / dx
		}
		incy := 0
		if dy != 0 {
			incy = vecy / dy
		}

		for tx, ty := int(fx)+incx, int(fy)+incy; tx != int(x) || ty != int(y); tx, ty = tx+incx, ty+incy {
			if grid[tx][ty] != nil {
				return false
			}
		}
		return target == nil || target.faction != selectedPiece.faction
	}

	return false
}

func isCheck(faction uint) bool {
	return false
}

func isCheckMate(faction uint) bool {
	return false
}

func startGame() {

	// 0 = white
	// 1 = black

	// 0 = pawn
	// 1 = knight
	// 2 = bishop
	// 3 = rook
	// 4 = queen
	// 5 = king

	// set up board

	// white pawns
	for x := uint(0); x < GRIDLENGTH; x++ {
		grid[x][6] = &(piece{
			class: 0, faction: 0,
		})
	}

	grid[0][7] = &(piece{
		class: 3, faction: 0,
	})

	grid[1][7] = &(piece{
		class: 1, faction: 0,
	})

	grid[2][7] = &(piece{
		class: 2, faction: 0,
	})

	grid[3][7] = &(piece{
		class: 4, faction: 0,
	})

	grid[4][7] = &(piece{
		class: 5, faction: 0,
	})

	grid[5][7] = &(piece{
		class: 2, faction: 0,
	})

	grid[6][7] = &(piece{
		class: 1, faction: 0,
	})

	grid[7][7] = &(piece{
		class: 3, faction: 0,
	})

	// black pawns
	for x := uint(0); x < GRIDLENGTH; x++ {
		grid[x][1] = &(piece{
			class: 0, faction: 1,
		})
	}

	grid[0][0] = &(piece{
		class: 3, faction: 1,
	})

	grid[1][0] = &(piece{
		class: 1, faction: 1,
	})

	grid[2][0] = &(piece{
		class: 2, faction: 1,
	})

	grid[3][0] = &(piece{
		class: 4, faction: 1,
	})

	grid[4][0] = &(piece{
		class: 5, faction: 1,
	})

	grid[5][0] = &(piece{
		class: 2, faction: 1,
	})

	grid[6][0] = &(piece{
		class: 1, faction: 1,
	})

	grid[7][0] = &(piece{
		class: 3, faction: 1,
	})

}

func winCheck() bool {
	return false
}
