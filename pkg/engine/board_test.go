package engine

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sortMoves(moves [][2]int) [][2]int {
	// Sorting the moves slice by first element, then by the second element if needed
	sort.Slice(moves, func(i, j int) bool {
		if moves[i][0] == moves[j][0] {
			return moves[i][1] < moves[j][1]
		}
		return moves[i][0] < moves[j][0]
	})
	return moves
}

func TestNewBoard(t *testing.T) {
	expected := Board{
		{"r", "n", "b", "q", "k", "b", "n", "r"}, // Black's back rank
		{"p", "p", "p", "p", "p", "p", "p", "p"}, // Black's pawns
		{"", "", "", "", "", "", "", ""},         // Empty rows
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"P", "P", "P", "P", "P", "P", "P", "P"}, // White's pawns
		{"R", "N", "B", "Q", "K", "B", "N", "R"}, // White's back rank
	}
	out := NewBoard()

	assert.Equal(t, expected, out, "Board should be equal")

}

func TestPawnMoves(t *testing.T) {
	var tests = []struct {
		name  string
		input struct {
			board Board
			x, y  int
		}
		want [][2]int
	}{
		{
			name: "White pawn initial move (no blocking)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, WhitePawn, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 6, y: 3,
			},
			want: [][2]int{{5, 3}, {4, 3}}, // White pawn can move 1 or 2 squares
		},
		{
			name: "White pawn regular move (one square forward)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, WhitePawn, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 5, y: 3,
			},
			want: [][2]int{{4, 3}}, // Only one square forward since it's not the first move
		},
		{
			name: "White pawn blocked by piece in front",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, WhitePawn, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, BlackPawn, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, WhitePawn, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 6, y: 3,
			},
			want: [][2]int{}, // Blocked by BlackPawn in front
		},
		{
			name: "White pawn diagonal capture",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, BlackKnight, Empty, BlackPawn, Empty, Empty, Empty},
					{Empty, Empty, Empty, WhitePawn, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 5, y: 3,
			},
			want: [][2]int{{4, 3}, {4, 2}, {4, 4}}, // Can move forward and capture diagonally
		},
		{
			name: "Black pawn initial move (no blocking)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, BlackPawn, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 1, y: 4,
			},
			want: [][2]int{{2, 4}, {3, 4}}, // Black pawn can move 1 or 2 squares
		},
		{
			name: "Black pawn blocked by piece in front",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, BlackPawn, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, WhitePawn, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 1, y: 4,
			},
			want: [][2]int{}, // Blocked by WhitePawn in front
		},
		{
			name: "Black pawn diagonal capture",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, BlackPawn, Empty, Empty, Empty},
					{Empty, Empty, WhiteBishop, Empty, Empty, WhiteKnight, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 1, y: 4,
			},
			want: [][2]int{{2, 4}, {3, 4}, {2, 5}}, // Can move forward and capture diagonally
		},
		{
			name: "Black pawn regular move (one square forward)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, BlackPawn, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 2, y: 4,
			},
			want: [][2]int{{3, 4}}, // Only one square forward since it's not the first move
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := tt.input.board.getPawnMoves(tt.input.x, tt.input.y)

			assert.Equal(t, sortMoves(tt.want), sortMoves(out), tt.name)
		})
	}

}

func TestKnightMoves(t *testing.T) {
	var tests = []struct {
		name  string
		input struct {
			board Board
			x, y  int
		}
		want [][2]int
	}{
		{
			name: "White Knight Moves from (3,4)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, WhiteKnight, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 3, y: 4,
			},
			want: [][2]int{
				{1, 3}, {1, 5}, {5, 3}, {5, 5}, {2, 2}, {2, 6}, {4, 2}, {4, 6},
			},
		},
		{
			name: "Black Knight Moves from (3,4)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, BlackKnight, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 3, y: 4,
			},
			want: [][2]int{
				{1, 3}, {1, 5}, {5, 3}, {5, 5}, {2, 2}, {2, 6}, {4, 2}, {4, 6},
			},
		},
		{
			name: "Knight Captures White Piece from (3,4)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, WhiteKnight, Empty, BlackPawn, Empty}, // Captures BlackPawn
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 3, y: 4,
			},
			want: [][2]int{
				{1, 3}, {1, 5}, {5, 3}, {5, 5}, {2, 2}, {2, 6}, {4, 2}, {4, 6}, // Captures BlackPawn at {4,6}
			},
		},
		{
			name: "Knight Captures Black Piece from (3,4)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, WhiteKnight, Empty, Empty, Empty}, // Can capture BlackPawn at {5, 4}
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, BlackPawn, Empty, Empty, Empty}, // BlackPawn at {5, 4}
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 3, y: 4,
			},
			want: [][2]int{
				{1, 3}, {1, 5}, {5, 3}, {5, 5}, {2, 2}, {2, 6}, {4, 2}, {4, 6}, // Captures BlackPawn at {5, 4}
			},
		},
		{
			name: "Knight Captures Opponent's Knight from (3,4)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, WhiteKnight, Empty, Empty, Empty}, // Can capture BlackKnight at {5, 3}
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, BlackKnight, Empty}, // BlackKnight at {5, 3}
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 3, y: 4,
			},
			want: [][2]int{
				{1, 3}, {1, 5}, {5, 3}, {5, 5}, {2, 2}, {2, 6}, {4, 2}, {4, 6}, // Captures BlackKnight at {5, 3}
			},
		},
		{
			name: "Knight Captures Empty Spaces from (0,0)",
			input: struct {
				board Board
				x, y  int
			}{
				board: Board{
					{WhiteKnight, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
					{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				},
				x: 0, y: 0,
			},
			want: [][2]int{
				{2, 1}, {1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := tt.input.board.getKnightMoves(tt.input.x, tt.input.y)

			assert.Equal(t, sortMoves(tt.want), sortMoves(out), tt.name)
		})
	}

}
