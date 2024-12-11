package engine

import "fmt"

// Helper function to get the piece symbol for a given square
func (b *Bitboard) getPieceSymbol(square int) string {
	mask := uint64(1) << square

	// Check each piece type's bitboard
	if b.whitePawns&mask != 0 {
		return "P"
	} else if b.whiteRooks&mask != 0 {
		return "R"
	} else if b.whiteKnights&mask != 0 {
		return "N"
	} else if b.whiteBishops&mask != 0 {
		return "B"
	} else if b.whiteQueen&mask != 0 {
		return "Q"
	} else if b.whiteKing&mask != 0 {
		return "K"
	} else if b.blackPawns&mask != 0 {
		return "p"
	} else if b.blackRooks&mask != 0 {
		return "r"
	} else if b.blackKnights&mask != 0 {
		return "n"
	} else if b.blackBishops&mask != 0 {
		return "b"
	} else if b.blackQueen&mask != 0 {
		return "q"
	} else if b.blackKing&mask != 0 {
		return "k"
	}

	return "." // Empty square
}

// PrintBoard prints the current board state in a human-readable format
// White pieces represented as capital leters, black ones with normal letters
func (b *Bitboard) PrintBoard() {
	fmt.Println("  a b c d e f g h") // Column labels
	fmt.Println(" +----------------")

	// Loop over each rank (row)
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d|", rank+1) // Row label (rank)

		// Loop over each file (column)
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			fmt.Printf("%s ", b.getPieceSymbol(square))
		}

		fmt.Println("|") // End of row
	}

	fmt.Println(" +----------------")
	fmt.Println("  a b c d e f g h") // Column labels at the bottom
}
