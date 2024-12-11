package engine

import (
	"fmt"
)

type Piece int
type Color int

const (
	Pawn Piece = iota
	Rook
	Knight
	Bishop
	Queen
	King
)
const (
	White Color = iota
	Black
)

type Bitboard struct {
	// defines bitboards for each piece types of white color
	whitePawns   uint64
	whiteRooks   uint64
	whiteKnights uint64
	whiteBishops uint64
	whiteQueen   uint64
	whiteKing    uint64

	// defines bitboards for each piece types of black color
	blackPawns   uint64
	blackRooks   uint64
	blackKnights uint64
	blackBishops uint64
	blackQueen   uint64
	blackKing    uint64

	turn Color // defines current turn
}

// Initializes new chess board with starting positions
// Every piece type of every color has its own bitboard
func NewBitboard() *Bitboard {
	return &Bitboard{
		whitePawns:   0x000000000000FF00, // Pawns on rank 2 (a2-h2)
		whiteRooks:   0x0000000000000081, // Rooks on a1, h1
		whiteKnights: 0x0000000000000042, // Knights on b1, g1
		whiteBishops: 0x0000000000000024, // Bishops on c1, f1
		whiteQueen:   0x0000000000000008, // Queen on d1
		whiteKing:    0x0000000000000010, // King on e1

		blackPawns:   0x00FF000000000000, // Pawns on rank 7 (a7-h7)
		blackRooks:   0x8100000000000000, // Rooks on a8, h8
		blackKnights: 0x4200000000000000, // Knights on b8, g8
		blackBishops: 0x2400000000000000, // Bishops on c8, f8
		blackQueen:   0x0800000000000000, // Queen on d8
		blackKing:    0x1000000000000000, // King on e8

		turn: White, // white pieces move first
	}
}

// Simple movement function to move pieces on the bitboard
func (b *Bitboard) MovePiece(piece Piece, fromSquare, toSquare Square) error {
	bitboard, err := b.getPieceBitboard(piece)

	if err != nil {
		return err
	}

	fromMask := uint64(1) << fromSquare
	toMask := uint64(1) << toSquare

	*bitboard &= ^fromMask // remove piece from bitboard
	*bitboard |= toMask    // place piece on bitboard

	// changing turn after moving the piece
	if b.turn == White {
		b.turn = Black
	} else {
		b.turn = White
	}

	return nil
}

func (b *Bitboard) getPieceBitboard(piece Piece) (*uint64, error) {
	switch piece {
	case Pawn:
		if b.turn == White {
			return &b.whitePawns, nil
		}
		return &b.blackPawns, nil
	case Rook:
		if b.turn == White {
			return &b.whiteRooks, nil
		}
		return &b.blackRooks, nil
	case Knight:
		if b.turn == White {
			return &b.whiteKnights, nil
		}
		return &b.blackKnights, nil
	case Bishop:
		if b.turn == White {
			return &b.whiteBishops, nil
		}
		return &b.blackBishops, nil
	case Queen:
		if b.turn == White {
			return &b.whiteQueen, nil
		}
		return &b.blackQueen, nil
	case King:
		if b.turn == White {
			return &b.whiteKing, nil
		}
		return &b.blackKing, nil
	default:
		return nil, fmt.Errorf("Unknown piece: %v", piece)
	}
}
