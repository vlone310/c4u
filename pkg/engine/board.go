package engine

type Board [8][8]Piece

type Piece string

const (
	// White Pieces
	WhitePawn   Piece = "P"
	WhiteKnight Piece = "N"
	WhiteBishop Piece = "B"
	WhiteRook   Piece = "R"
	WhiteQueen  Piece = "Q"
	WhiteKing   Piece = "K"

	// Black Pieces
	BlackPawn   Piece = "p"
	BlackKnight Piece = "n"
	BlackBishop Piece = "b"
	BlackRook   Piece = "r"
	BlackQueen  Piece = "q"
	BlackKing   Piece = "k"

	// Empty square
	Empty Piece = ""
)

// Function to initialize starting board
func NewBoard() Board {
	return Board{
		{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook},
		{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn},
		{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook},
	}
}

func (board Board) getPawnMoves(x, y int) [][2]int {
	moves := [][2]int{}
	direction := -1 // White pawns move up the board
	isWhite := isWhitePiece(board[x][y])

	if !isWhite {
		direction = 1 // Black pawns move down the board
	}

	if isInBounds(x+direction, y) && board[x+direction][y] == Empty {
		moves = append(moves, [2]int{x + direction, y})

		// Initial move: 2 squares forward if the first move and both squares are empty
		if (isWhite && x == 6) || (!isWhite && x == 1) {
			if board[x+2*direction][y] == Empty {
				moves = append(moves, [2]int{x + 2*direction, y})
			}
		}
	}

	// Captures: Diagonal left and diagonal right
	if isInBounds(x+direction, y-1) && isOpponentPiece(board[x+direction][y-1], isWhite) {
		moves = append(moves, [2]int{x + direction, y - 1})
	}
	if isInBounds(x+direction, y+1) && isOpponentPiece(board[x+direction][y+1], isWhite) {
		moves = append(moves, [2]int{x + direction, y + 1})
	}

	// TODO: En Passant, promotion?

	return moves
}

func (board Board) getKnightMoves(x, y int) [][2]int {
	moves := [][2]int{}
	knightMoves := [][2]int{
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
		{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	}
	isWhite := isWhitePiece(board[x][y])

	for _, move := range knightMoves {
		nextX := x + move[0]
		nextY := y + move[1]

		if isInBounds(nextX, nextY) && !isOpponentPiece(board[nextX][nextY], isWhite) {
			moves = append(moves, [2]int{nextX, nextY})
		}
	}

	return moves
}

// isWhite checks if the piece belongs to the white player
func isWhitePiece(p Piece) bool {
	return p == WhitePawn || p == WhiteKnight || p == WhiteBishop ||
		p == WhiteRook || p == WhiteQueen || p == WhiteKing
}

// isBlack checks if the piece belongs to the black player
func isBlackPiece(p Piece) bool {
	return p == BlackPawn || p == BlackKnight || p == BlackBishop ||
		p == BlackRook || p == BlackQueen || p == BlackKing
}

func isOpponentPiece(p Piece, isWhite bool) bool {
	if isWhite {
		return isBlackPiece(p)
	}
	return isWhitePiece(p)
}

// isEmpty checks if the square is empty
func isEmptyPiece(p Piece) bool {
	return p == Empty
}

// Checks if a position is within the bounds of the board
func isInBounds(x, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}
