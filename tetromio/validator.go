package tetromio

/*
The function is meant to validate the tetromino. The function check for the number of hashes found in
the tetromino, which should be exactly 4 and validates the hash conections, which should be either
6 or 8(for the square tetromino).
*/
func MatrixMorphologist(tetromino [4][4]string) bool {
	c, d := 0, 0

	for i, block := range tetromino {
		for x, minorBlock := range block {
			if minorBlock != "." {
				d++
				if i+1 < 4 && tetromino[i+1][x] != "." {
					c++
				}
				if i-1 >= 0 && tetromino[i-1][x] != "." {
					c++
				}
				if x+1 < 4 && tetromino[i][x+1] != "." {
					c++
				}
				if x-1 >= 0 && tetromino[i][x-1] != "." {
					c++
				}
			}
		}
	}

	if d != 4 {
		return false
	}
	if c == 6 || c == 8 {
		return true
	}
	return false
}

/*
The function takes the indices of the row and the column currently
and checks if there will be a mistake of any sort  when
placing the tetrominoes onto the board, i.e. over-lapping the normal
length of thr tetromino or maybe trying to place a tetromino where
another tetromino already occupies at the moment.
*/
func ValidInsertion(currentRowIndex, currentColumnIndex int, tetro [4][4]string) bool {
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			if tetro[a][b] != "." {
				if currentRowIndex+a == len(grid) || currentColumnIndex+b == len(grid) || grid[currentRowIndex+a][currentColumnIndex+b] != "." {
					return false
				}
			}
		}
	}
	return true
}
