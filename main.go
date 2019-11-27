package main

import (
    "log"
    "fmt"
    "math/rand"
    "time"
    "strings"
    "strconv"
)

// default to 3x3 board
const COL_SIZE = 3
const BOARD_SIZE = COL_SIZE*COL_SIZE
// define move results
const CONTINUE = "CONTINUE"
const WIN = "WIN"
const DRAW = "DRAW"
const MASTER_ID = 1
const PLAYER_ID = -1

func firstMove() ([]int, []int, [BOARD_SIZE-1]int, error) {
    // converting to uni-dimensional array
    // [x][y] = (COL_SIZE * x) + y
    rand.Seed(time.Now().UnixNano())
    var master_pos, player_pos []int
	master_row := rand.Intn(COL_SIZE)
	master_col := rand.Intn(COL_SIZE)
	var winning_board [BOARD_SIZE-1]int
    winning_board[master_row] += MASTER_ID
    winning_board[master_col+COL_SIZE] += MASTER_ID
	if master_row == master_col {
		winning_board[2*COL_SIZE] += MASTER_ID
	} else if (master_row + master_col) == (COL_SIZE - 1) {
		winning_board[(2*COL_SIZE)+1] += MASTER_ID
	}
    master_pos = append(master_pos, COL_SIZE*master_row+master_col)
    log.Println("Master player will start at: ", master_pos)

    return master_pos, player_pos, winning_board, nil
}

func generateMatrix(master_pos, player_pos []int) {
    var current_board [BOARD_SIZE]string
    for i := 0; i < BOARD_SIZE; i++ {
        if findIntInArray(i, master_pos) {
            current_board[i] = "O"
        } else if findIntInArray(i, player_pos) {
            current_board[i] = "X"
        } else {
            current_board[i] = "-"
        }
    }

    drawBoard(current_board)
}

func drawBoard(currentBoard [BOARD_SIZE]string) {
    for i := 0; i < COL_SIZE; i++ {
        fmt.Println(currentBoard[i*COL_SIZE:i*COL_SIZE+3])
    }
}

func movePlayer (master_pos, player_pos []int, winning_board [BOARD_SIZE-1]int) (string, string) {
    for {
        fmt.Print("Enter your next move (row, col): ")
		var new_player_index string
		fmt.Scanln(&new_player_index)
        row, errRow := strconv.Atoi(strings.Split(new_player_index, ",")[0])
        if errRow != nil {
            log.Print(errRow)
        }
        col, errCol := strconv.Atoi(strings.Split(new_player_index, ",")[1])
        if errCol != nil {
            log.Print(errCol)
        }
        new_player_pos := int(COL_SIZE*row+col)
		if checkNewPos(new_player_pos, append(master_pos, player_pos...)) {
			player_pos = append(player_pos, new_player_pos)
			winning_board[row] += PLAYER_ID
			winning_board[col+COL_SIZE] += PLAYER_ID
			if row == col {
				winning_board[2*COL_SIZE] += PLAYER_ID
			} else if (row + col) == (COL_SIZE - 1) {
				winning_board[(2*COL_SIZE)+1] += PLAYER_ID
			}
			generateMatrix(master_pos, player_pos)
			game_status := checkWin(winning_board, append(master_pos, player_pos...))
			if game_status != CONTINUE{
				return game_status, "Player"
			}
			new_master_pos, master_row, master_col := moveMaster(master_pos, player_pos)
			master_pos = append(master_pos, new_master_pos)
			winning_board[master_row] += MASTER_ID
			winning_board[master_col+COL_SIZE] += MASTER_ID
			if master_row == master_col {
				winning_board[(2*COL_SIZE)] += MASTER_ID
			} else if (master_row + master_col) == (COL_SIZE - 1) {
				winning_board[(2*COL_SIZE)+1] += MASTER_ID
			}
			log.Println("Master moves")
			generateMatrix(master_pos, player_pos)
			game_status = checkWin(winning_board, append(master_pos, player_pos...))
			if game_status != CONTINUE{
				return game_status, "Master"
			}
		}
    }
}

func checkNewPos (pos int, positions []int) bool {
	if findIntInArray(pos, positions) || pos >= BOARD_SIZE {
		return false
	}

	return true
}

func moveMaster (master_pos, player_pos []int) (int, int, int) {
	for {
		master_row := rand.Intn(COL_SIZE)
		master_col := rand.Intn(COL_SIZE)
		new_master_pos := COL_SIZE*master_row+master_col
		if checkNewPos (new_master_pos, append(master_pos, player_pos...)) {

			return new_master_pos, master_row, master_col
		}
	}
}

func checkWin(winning_board [BOARD_SIZE-1]int, positions []int) string {
	log.Println("WINNING BOARD: ", winning_board)
	board := winning_board[:]
	if findIntInArray(-3, board) || findIntInArray(3, board) {
		return WIN
	} else if len(positions) == BOARD_SIZE {
		return DRAW
	}

    return CONTINUE
}

func findStrInArray(value string, array [BOARD_SIZE]string) (bool) {
    for x := 0; x < len(array); x++ {
        if value == array[x]{
            return true
        }
    }

	return false
}


func findIntInArray(value int, array[]int) (bool) {
    for x := 0; x < len(array); x++ {
        if value == array[x]{
            return true
        }
    }

    return false
}

func main() {
    log.Println("Starting tic-tac-toe game!")
    master_pos, player_pos, winning_board, errFirst := firstMove()
    if errFirst != nil {
        log.Println(errFirst)
    }
    generateMatrix(master_pos, player_pos)
	result, last_player := movePlayer(master_pos, player_pos, winning_board)
	if result == WIN {
		log.Println(last_player, " wins!")
	} else {
		log.Println(result)
	}
}
