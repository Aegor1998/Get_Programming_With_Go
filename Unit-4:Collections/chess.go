package main

import "fmt"

type chessPiece struct {
	pieceName  string
	pieceColor string
}

const (
	kingPiece   = "K"
	queenPiece  = "Q"
	knightPiece = "H"
	bishopPiece = "B"
	rookPiece   = "R"
	pawnPiece   = "P"
	whiteSide   = "w"
	blackSide   = "b"
) //Piece Names + Side Names

var board [8][8]chessPiece

func main() {
	generateBoard()
	showBoard()
}

func generateBoard() {
	for r := 0; r < len(board); r++ { //r==row position
		for c := 0; c < len(board); c++ { //0==A, 1==B, ..., 7==H
			switch r {
			case 0:
				switch c {
				case 0, 7:
					board[r][c].pieceName, board[r][c].pieceColor = rookPiece, blackSide
				case 1, 6:
					board[r][c].pieceName, board[r][c].pieceColor = knightPiece, blackSide
				case 2, 5:
					board[r][c].pieceName, board[r][c].pieceColor = bishopPiece, blackSide
				case 3:
					board[r][c].pieceName, board[r][c].pieceColor = queenPiece, blackSide
				case 4:
					board[r][c].pieceName, board[r][c].pieceColor = kingPiece, blackSide
				}
			case 1:
				board[r][c].pieceName, board[r][c].pieceColor = pawnPiece, blackSide
			case 6:
				board[r][c].pieceName, board[r][c].pieceColor = pawnPiece, whiteSide
			case 7:
				switch c {
				case 0, 7:
					board[r][c].pieceName, board[r][c].pieceColor = rookPiece, whiteSide
				case 1, 6:
					board[r][c].pieceName, board[r][c].pieceColor = knightPiece, whiteSide
				case 2, 5:
					board[r][c].pieceName, board[r][c].pieceColor = bishopPiece, whiteSide
				case 3:
					board[r][c].pieceName, board[r][c].pieceColor = queenPiece, whiteSide
				case 4:
					board[r][c].pieceName, board[r][c].pieceColor = kingPiece, whiteSide
				}
			default:
				board[r][c].pieceName = " "
				board[r][c].pieceColor = " "
			}
		}
	}
} //Makes or Resets Board

func showBoard() {
	for r := 0; r < len(board); r++ { //r==row position
		for c := 0; c < len(board); c++ { //0==A, 1==B, ..., 7==H
			fmt.Printf("|%s%s", board[r][c].pieceName, board[r][c].pieceColor)
		}
		fmt.Printf("|\n")
	}
} //Display's Boards Current State
