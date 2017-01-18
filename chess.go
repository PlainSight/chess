package main

func main() {
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

var (
	grid     [GRIDLENGTH][GRIDLENGTH]*piece
	selected *position
)

func grab(x uint, y uint) {
	if grid[x][y] != nil {
		if grid[x][y].faction == playerFaction {
			selected = &position{x: x, y: y}
		}
	}
}

func move(x uint, y uint) {
	if selected != nil && grid[selected.x][selected.y] != nil {
		if validateMove(x, y) {
			grid[x][y] = grid[selected.x][selected.y]
			grid[selected.x][selected.y] = nil
			selected = nil
		}
	}
}

func validateMove(x uint, y uint) bool {
	// from values
	fx := selected.x
	fy := selected.y
	// selected piece
	selectedPiece := grid[selected.x][selected.y]
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
			return target == nil && advance == 1 || (canDoubleStep && advance == 2)
		} else {
			sidemovement := fy - y
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

func isCheck() bool {
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
