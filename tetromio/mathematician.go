package tetromio

/*
The function takes a 4x4 slice of string and returns the same,
only making some modifications on the tetromino block. It checks whether
there exists any empty column and attempts to push the tetromino block to the left
as well it checks for an empty row and tries to push the tetromino to the
very top. The aim is to place the tetrtomino as top-left as possible
*/
func OptimizeTetromino(tetromino [4][4]string) [4][4]string {
	rowIndex := 0
	for {
		emptyBlocks := 0
		for columnIndex := 0; columnIndex < 4; columnIndex++ {
			if tetromino[rowIndex][columnIndex] == "." {
				emptyBlocks++
			}
		}
		if emptyBlocks == 4 {
			tetromino = PushUp(tetromino)
			continue
		}
		break
	}
	for {
		emptyBlocks := 0
		for row := 0; row < 4; row++ {
			if tetromino[row][rowIndex] == "." {
				emptyBlocks++
			}
		}
		if emptyBlocks == 4 {
			tetromino = PushLeft(tetromino)
			continue
		}
		break
	}
	return tetromino
}

/*
The function will check if there is an empty space at the top(a column)
If found the function returns the tetromino pushed to the top.
*/
func PushUp(tetromino [4][4]string) [4][4]string {
	temp := tetromino[0]
	tetromino[0] = tetromino[1]
	tetromino[1] = tetromino[2]
	tetromino[2] = tetromino[3]
	tetromino[3] = temp
	return tetromino
}

/*
This will push the tetromino to the very left of its 4x4 grid
this make the placements of the tetrominoes in the board to
be much efficient.
*/
func PushLeft(tetromino [4][4]string) [4][4]string {
	tetromino = flipper(tetromino)
	tetromino = PushUp(tetromino)
	tetromino = flipper(tetromino)
	return tetromino
}

/*
The function flips the tetromino like a square, geometrically placed,
such that the edges(2 of the vertex) probably the opposite ones.
If a straight vertical line is drawn to march the vertices,
we mark the other vertex to the left as 'a' and the other to the
right as b, now we flip the square, such that the a occupies the 
previous position of b, and vice versa.
*/
func flipper(tetromino [4][4]string) [4][4]string {
	xl := len(tetromino[0])
	yl := len(tetromino)
	var result [4][4]string
	for i := range result {
		result[i] = [4]string{}
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = tetromino[j][i]
		}
	}
	return result
}
