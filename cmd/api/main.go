package main

import (
	"log/slog"

	"github.com/vlone310/c4u/pkg/engine"
)

func main() {
	bitboard := engine.NewBitboard()
	err := bitboard.MovePiece(100, engine.A2, engine.A3)
	if err != nil {
		slog.Error(err.Error())
	}
	bitboard.PrintBoard()
}
