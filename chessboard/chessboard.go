package chessboard

import "slices"

type File []bool
type Chessboard map[string]File

// chessboard := map[string]File{"A": {true, false,true, false...}, "B": {false, true, false},...}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	rank := cb[file]
	var occupied int

	if slices.Contains(rank, true) {
		for _, v := range rank {
			if v {
				occupied++
			}
		}
	}
	return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var occupied int
	for _, file := range cb {
		for i, f := range file {
			if i+1 == rank && f {
				occupied++
			}
		}
	}
	return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var occupied int
	for _, file := range cb {
		for range file {
			occupied++
		}
	}
	return occupied
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var occupied int
	for _, file := range cb {
		for _, f := range file {
			if f {
				occupied++
			}
		}
	}
	return occupied
}
