//https://www.hackerrank.com/challenges/simplified-chess-engine/problem
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

const (
	boardSize = 4
	white     = 'W'
	black     = 'B'
)

var maxMove int

func main() {
	var g int
	fmt.Scan(&g)

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < g; i++ {
		whites, blacks := make([]Figure, 0), make([]Figure, 0)

		wbm, _ := reader.ReadString('\n')
		wbmArr := strings.Split(strings.Trim(wbm, "\n"), " ")
		w, _ := strconv.Atoi(wbmArr[0])
		b, _ := strconv.Atoi(wbmArr[1])
		m, _ := strconv.Atoi(wbmArr[2])
		maxMove = m

		for i := 0; i < w+b; i++ {
			str, _ := reader.ReadString('\n')
			t, c, r := []rune(str)[0], []rune(str)[2], []rune(str)[4]
			if i < w {
				whites = append(whites, NewPiece(t, c, r, white))
			} else {
				blacks = append(blacks, NewPiece(t, c, r, black))
			}
		}

		if !makeMove(0, 'W', whites, blacks) {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}

func makeMove(move int, p rune, whites, blacks []Figure) bool {
	move++
	if move > maxMove {
		return false
	}

	boards, win := getMoves(p, whites, blacks)
	//fmt.Println("##########")
	//fmt.Println(Board{Whites: whites, Blacks: blacks})
	//fmt.Println(move)
	//fmt.Println(string(win))
	//for _, b := range boards {
	//	fmt.Println(b)
	//	fmt.Println(".............")
	//}
	if win == black {
		return false
	}
	//if win == black {
	//	return false
	//}

	for _, b := range boards {
		if !makeMove(move, switchPlayer(p), b.Whites, b.Blacks) {
			return false
		}
		//if makeMove(move, switchPlayer(p), b.Whites, b.Blacks) {
		//	break
		//}
	}
	return true
}

func switchPlayer(p rune) rune {
	if p == white {
		return black
	}
	return white
}

func getMoves(p rune, whites, blacks []Figure) ([]Board, rune) {
	boards := make([]Board, 0)
	if p == white {
		for i, f := range whites {
			figures := f.GetMoves()
			if f.GetType() == 'Q' {
				figures = removeDangerousPositions(figures, whites, blacks)
			}
			//fmt.Println(f)
			//fmt.Println(figures)
			for _, newFigure := range figures {
				//fmt.Println(",,,,,,,,,,")
				//fmt.Println(newFigure)
				if newWhites, newBlacks, ok, win := canMove(newFigure, i, whites, blacks); ok {
					//fmt.Println(newWhites)
					//fmt.Println(newBlacks)
					//fmt.Println(ok)
					//fmt.Println(win)
					if win {
						return nil, white
					}
					boards = append(boards, Board{newWhites, newBlacks})
				}
			}
		}
	} else {
		for i, f := range blacks {
			figures := f.GetMoves()
			if f.GetType() == 'Q' {
				figures = removeDangerousPositions(figures, blacks, whites)
			}
			for _, newFigure := range figures {
				if newBlacks, newWhites, ok, win := canMove(newFigure, i, blacks, whites); ok {
					if win {
						return nil, black
					}
					boards = append(boards, Board{newWhites, newBlacks})
				}
			}
		}
	}
	return boards, '0'
}

func removeDangerousPositions(figures, allies, enemies []Figure) []Figure {
	var newFigures []Figure
	for _, f := range figures {
		canMove := true
		x, y := f.GetCoords()
		for _, ef := range enemies {
			efMoves := ef.GetMoves()
			for _, m := range efMoves {
				eX, eY := m.GetCoords()
				if eX == x && eY == y {
					if !findObstacles(ef, m, enemies, allies) {
						canMove = false
					}
				}
			}
		}
		if canMove {
			if newFigures == nil {
				newFigures = make([]Figure, 0)
			}
			newFigures = append(newFigures, f)
		}
	}
	return newFigures
}

func findObstacles(oldFigure, newFigure Figure, allies, enemies []Figure) bool {
	newX, newY := newFigure.GetCoords()
	positions := oldFigure.Move(Pos{newX, newY})
	//fmt.Println(positions)
	for j, p := range positions {
		for _, f := range allies {
			x, y := f.GetCoords()
			if p.X == x && p.Y == y {
				return true
			}
		}
		for _, f := range enemies {
			x, y := f.GetCoords()
			if p.X == x && p.Y == y && j < len(positions)-1 {
				return true

			}
		}
	}
	return false
}

func canMove(newFigure Figure, idx int, allies, enemies []Figure) ([]Figure, []Figure, bool, bool) {
	win := false
	oldFigure := allies[idx]
	newX, newY := newFigure.GetCoords()
	positions := oldFigure.Move(Pos{newX, newY})
	//fmt.Println(positions)
	for j, p := range positions {
		for _, f := range allies {
			x, y := f.GetCoords()
			if p.X == x && p.Y == y {
				return nil, nil, false, false
			}
		}
		for _, f := range enemies {
			x, y := f.GetCoords()
			if p.X == x && p.Y == y && j < len(positions)-1 {
				return nil, nil, false, false

			}
		}
	}

	newAllies, newEnemies := make([]Figure, 0), make([]Figure, 0)
	for i, f := range allies {
		if i == idx {
			newAllies = append(newAllies, newFigure)
			continue
		}
		newAllies = append(newAllies, f)
	}
	for _, f := range enemies {
		x, y := f.GetCoords()
		if newX == x && newY == y {
			if f.GetType() == 'Q' {
				win = true
			}
		} else {
			newEnemies = append(newEnemies, f)
		}
	}

	return newAllies, newEnemies, true, win
}

type Board struct {
	Whites []Figure
	Blacks []Figure
}

func (b Board) String() string {
	board := make([][]string, boardSize)
	for i := 0; i < boardSize; i++ {
		board[i] = make([]string, boardSize)
	}
	for _, f := range append(b.Whites, b.Blacks...) {
		var sign string
		if f.GetPlayer() == white {
			sign = string(f.GetType())
		} else {
			sign = strings.ToLower(string(f.GetType()))
		}
		i, j := f.GetCoords()
		board[i][j] = sign
	}

	var str string
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			sign := "-"
			if board[i][j] != "" {
				sign = board[i][j]
			}
			str += fmt.Sprintf("%s ", sign)
		}
		str += "\n"
	}
	return str
}

type Coords interface {
	GetCoords() (int, int)
}

type Figure interface {
	Coords
	GetMoves() []Figure
	Move(dest Pos) []Pos
	GetType() rune
	GetPlayer() rune
}

type Pos struct {
	X, Y int
}

type Piece struct {
	Position Pos
	Player   rune
	Type     rune
}

func (p *Piece) String() string {
	return fmt.Sprintf("%c|%c|%d:%d", p.Player, p.Type, p.Position.X, p.Position.Y)
}

func (p *Piece) GetType() rune {
	return p.Type
}

func (p *Piece) GetPlayer() rune {
	return p.Player
}

func (p *Piece) GetCoords() (int, int) {
	return p.Position.X, p.Position.Y
}

func NewPiece(t, c, r, p rune) Figure {
	x, y := getCoords(c, r)
	pos := Pos{x, y}
	return newFigure(pos, p, t)
}

func newFigure(pos Pos, p, t rune) Figure {
	switch t {
	case 'Q':
		return &Queen{Piece{pos, p, 'Q'}}
	case 'N':
		return &Knight{Piece{pos, p, 'N'}}
	case 'B':
		return &Bishop{Piece{pos, p, 'B'}}
	case 'R':
		return &Rook{Piece{pos, p, 'R'}}
	}

	panic("somethig went wrong")
}

func getCoords(c, r rune) (x, y int) {
	switch c {
	case 'A':
		y = 0
	case 'B':
		y = 1
	case 'C':
		y = 2
	case 'D':
		y = 3
	}
	switch r {
	case '4':
		x = 0
	case '3':
		x = 1
	case '2':
		x = 2
	case '1':
		x = 3
	}
	return
}

func getPositions(f Figure, dest Pos, next func(x, y int) (int, int)) []Pos {
	positions := make([]Pos, 0)
	x, y := f.GetCoords()
	for !reached(x, y, dest) {
		x, y = next(x, y)
		positions = append(positions, Pos{x, y})
	}
	return positions
}

func reached(x, y int, dest Pos) bool {
	return x == dest.X && y == dest.Y
}

type Queen struct {
	Piece
}

func (q *Queen) Move(dest Pos) []Pos {
	positions := diagonals(q, dest)
	positions = append(positions, horizontals(q, dest)...)
	return positions
}

func (q *Queen) GetMoves() []Figure {
	return possibleMoves(q.Position.X, q.Position.Y, q.Player, q.Type, func(x, y float64) bool {
		return x == y || x == 0 || y == 0
	})
}

type Knight struct {
	Piece
}

func (k *Knight) Move(dest Pos) []Pos {
	return []Pos{dest}
}

func (k *Knight) GetMoves() []Figure {
	return possibleMoves(k.Position.X, k.Position.Y, k.Player, k.Type, func(x, y float64) bool {
		return (x == 1 && y == 2) || (x == 2 && y == 1)
	})
}

type Bishop struct {
	Piece
}

func (b *Bishop) Move(dest Pos) []Pos {
	return diagonals(b, dest)
}

func (b *Bishop) GetMoves() []Figure {
	return possibleMoves(b.Position.X, b.Position.Y, b.Player, b.Type, func(x, y float64) bool {
		return x == y
	})
}

type Rook struct {
	Piece
}

func (r *Rook) Move(dest Pos) []Pos {
	return horizontals(r, dest)
}

func (r *Rook) GetMoves() []Figure {
	return possibleMoves(r.Position.X, r.Position.Y, r.Player, r.Type, func(x, y float64) bool {
		return x == 0 || y == 0
	})
}

func possibleMoves(x, y int, p, t rune, cond func(x, y float64) bool) []Figure {
	figures := make([]Figure, 0)
	for i := 0; i < boardSize; i++ {
		xDiff := math.Abs(float64(i - x))
		for j := 0; j < boardSize; j++ {
			yDiff := math.Abs(float64(j - y))
			if !(x == i && y == j) && cond(xDiff, yDiff) {
				figures = append(figures, newFigure(Pos{i, j}, p, t))
			}
		}
	}
	return figures
}

func diagonals(f Figure, dest Pos) []Pos {
	var positions []Pos
	x, y := f.GetCoords()
	switch {
	case dest.X > x && dest.Y > y:
		positions = getPositions(f, dest, func(x, y int) (int, int) {
			x++
			y++
			return x, y
		})
	case dest.X > x && dest.Y < y:
		positions = getPositions(f, dest, func(x, y int) (int, int) {
			x++
			y--
			return x, y
		})
	case dest.X < x && dest.Y < y:
		positions = getPositions(f, dest, func(x, y int) (int, int) {
			x--
			y--
			return x, y
		})
	case dest.X < x && dest.Y > y:
		positions = getPositions(f, dest, func(x, y int) (int, int) {
			x--
			y++
			return x, y
		})
	}

	return positions
}

func horizontals(f Figure, dest Pos) []Pos {
	var positions []Pos
	x, y := f.GetCoords()
	switch {
	case dest.X == x && dest.Y > y:
		positions = getPositions(f, dest, func(x, y int) (int, int) {
			y++
			return x, y
		})
	case dest.X == x && dest.Y < y:
		positions = getPositions(f, dest, func(x, y int) (int, int) {
			y--
			return x, y
		})
	case dest.X > x && dest.Y == y:
		positions = getPositions(f, dest, func(x, y int) (int, int) {
			x++
			return x, y
		})
	case dest.X < x && dest.Y == y:
		positions = getPositions(f, dest, func(x, y int) (int, int) {
			x--
			return x, y
		})
	}

	return positions
}
