package main

import (
    "log"
    "fmt"
    "strings"
    "strconv"
)

// should work for any square board
const COL_SIZE = 3
const BOARD_SIZE = COL_SIZE*COL_SIZE

const CONTINUE = "CONTINUE"
const WIN = "WIN"
const DRAW = "DRAW"
var USERID = map[string]int {
    "master": 1,
    "player": -1,
}

type BoardProperties struct {
    WinningBoard        [BOARD_SIZE - 1]int
    MasterPositions     []int
    PlayerPositions     []int
}

type NewIndex struct {
    Col         int
    Row         int
    Position    int
}

func checkNewPos (pos int, positions []int) bool {
    if len(findIntInArray(pos, positions)) > 0 || pos >= BOARD_SIZE {
        return false
    }

    return true
}

func checkWin(board_status BoardProperties) string {
    positions := append(board_status.MasterPositions, board_status.PlayerPositions...)
    board := board_status.WinningBoard[:]
    log.Println("WINING BOARD: ", board)
    if len(findIntInArray(-COL_SIZE, board)) > 0 || len(findIntInArray(COL_SIZE, board)) > 0 {
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

func generateMatrix(master_pos, player_pos []int) {
    var current_board [BOARD_SIZE]string
    for i := 0; i < BOARD_SIZE; i++ {
        if len(findIntInArray(i, master_pos)) > 0 {
            current_board[i] = "O"
        } else if len(findIntInArray(i, player_pos)) > 0 {
            current_board[i] = "X"
        } else {
            current_board[i] = "-"
        }
    }

    drawBoard(current_board)
}

func drawBoard(currentBoard [BOARD_SIZE]string) {
    for i := 0; i < COL_SIZE; i++ {
        log.Println(currentBoard[i*COL_SIZE:i*COL_SIZE+COL_SIZE])
    }
}

func moveMaster (board_status BoardProperties) (NewIndex) {
    var new_index NewIndex
    for {
        master_row, master_col := applyStrategy(board_status)
        new_master_pos := COL_SIZE*master_row+master_col
        if checkNewPos (new_master_pos, append(board_status.MasterPositions, board_status.PlayerPositions...)) {
            new_index = NewIndex{Col: master_col, Row: master_row, Position: new_master_pos}

            return new_index
        }
    }
}

func movePlayer (board_status BoardProperties) (NewIndex) {
    var new_index NewIndex
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
        if checkNewPos(new_player_pos, append(board_status.MasterPositions, board_status.PlayerPositions...)) {
            new_index = NewIndex{Col: player_col, Row: player_row, Position: new_player_pos}

            return new_index
        }
    }
}

func main() {
    log.Println("Starting tic-tac-toe game!")
    var last_player, game_status string
    var board_status BoardProperties
    for {
        log.Println("Master moves")
        new_index := moveMaster(board_status)
        board_status = BoardProperties{MasterPositions: append(board_status.MasterPositions, new_index.Position),
                                       PlayerPositions: board_status.PlayerPositions,
                                       WinningBoard: updateWinningBoard(board_status.WinningBoard, new_index.Col, new_index.Row, "master")}
        log.Println("Master positions: ", board_status.MasterPositions)
        generateMatrix(board_status.MasterPositions, board_status.PlayerPositions)
        game_status = checkWin(board_status)
        if game_status != CONTINUE {
            last_player = "IA"
            break
        }
        log.Println("Player moves")
        new_index = movePlayer(board_status)
        board_status = BoardProperties{PlayerPositions: append(board_status.PlayerPositions, new_index.Position),
                                       MasterPositions: board_status.MasterPositions,
                                       WinningBoard: updateWinningBoard(board_status.WinningBoard, new_index.Col, new_index.Row, "player")}
        log.Println("Player positions: ", board_status.PlayerPositions)
        generateMatrix(board_status.MasterPositions, board_status.PlayerPositions)
        game_status = checkWin(board_status)
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

