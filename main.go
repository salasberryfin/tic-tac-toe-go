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
var USERID = map[string]int {
    "master": 1,
    "player": -1,
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

func moveMaster (master_pos, player_pos []int) (int, int, int) {
    rand.Seed(time.Now().UnixNano())
	for {
		master_row := rand.Intn(COL_SIZE)
		master_col := rand.Intn(COL_SIZE)
		new_master_pos := COL_SIZE*master_row+master_col
		if checkNewPos (new_master_pos, append(master_pos, player_pos...)) {

			return new_master_pos, master_row, master_col
		}
	}
}

func movePlayer (master_pos, player_pos []int) (int, int, int) {
    for {
        fmt.Print("Enter your next move (row, col): ")
		var new_player_index string
		fmt.Scanln(&new_player_index)
        player_row, errRow := strconv.Atoi(strings.Split(new_player_index, ",")[0])
        if errRow != nil {
            log.Print(errRow)
        }
        player_col, errCol := strconv.Atoi(strings.Split(new_player_index, ",")[1])
        if errCol != nil {
            log.Print(errCol)
        }
        new_player_pos := int(COL_SIZE*player_row+player_col)
		if checkNewPos(new_player_pos, append(master_pos, player_pos...)) {

			return new_player_pos, player_row, player_col
		}
    }
}

func checkNewPos (pos int, positions []int) bool {
	if findIntInArray(pos, positions) || pos >= BOARD_SIZE {
		return false
	}

	return true
}

func checkWin(winning_board [BOARD_SIZE-1]int, positions []int) string {
	//log.Println("WINNING BOARD: ", winning_board)
	board := winning_board[:]
	if findIntInArray(-3, board) || findIntInArray(3, board) {
		return WIN
	} else if len(positions) == BOARD_SIZE {
		return DRAW
	}

    return CONTINUE
}

func updateWinningBoard (winning_board [BOARD_SIZE-1]int, row, col int, player string) [BOARD_SIZE-1]int {
    winning_board[row] += USERID[player]
    winning_board[col+COL_SIZE] += USERID[player]
	if row == col {
		winning_board[2*COL_SIZE] += USERID[player]
	}
    if (row + col) == (COL_SIZE - 1) {
		winning_board[(2*COL_SIZE)+1] += USERID[player]
	}

    return winning_board
}

func main() {
    log.Println("Starting tic-tac-toe game!")
    var master_pos, player_pos []int
    var last_player, game_status string
	var winning_board [BOARD_SIZE-1]int
    for {
        log.Println("Master moves")
        new_master_pos, master_row, master_col := moveMaster(master_pos, player_pos)
        master_pos = append(master_pos, new_master_pos)
        winning_board = updateWinningBoard(winning_board, master_row, master_col, "master")
        generateMatrix(master_pos, player_pos)
        game_status = checkWin(winning_board, append(master_pos, player_pos...))
        if game_status != CONTINUE {
            last_player = "IA"
            break
        }
        log.Println("Player moves")
        new_player_pos, player_row, player_col := movePlayer(master_pos, player_pos)
        player_pos = append(player_pos, new_player_pos)
        winning_board = updateWinningBoard(winning_board, player_row, player_col, "player")
        generateMatrix(master_pos, player_pos)
        game_status = checkWin(winning_board, append(master_pos, player_pos...))
        if game_status != CONTINUE {
            last_player = "Player"
            break
        }
    }
	if game_status == WIN {
		log.Println(last_player, " wins!")
	} else {
		log.Println(game_status)
	}
}

