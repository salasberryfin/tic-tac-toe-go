package main

import (
    "log"
    "fmt"
    "math/rand"
    "time"
    "errors"
)

// default to 3x3 board
const BOARD_SIZE = 9
const COL_SIZE = 3

func firstMove() ([BOARD_SIZE]string, error) {
    // (0, 0) (0, 1) (0, 2)
    // (1, 0) (1, 1) (1, 2)
    // (2, 0) (2, 1) (2, 2)
    // converting to uni-dimensional array
    // [x][y] = (BOARD_SIZE * x) + y
    rand.Seed(time.Now().UnixNano())
    var master_pos, player_pos []int
    master_pos = append(master_pos, rand.Intn(BOARD_SIZE))
    log.Println("Master player will start at: ", master_pos)
    current_board, errMat:= generateMatrix(master_pos, player_pos)
    if errMat != nil {
        return [BOARD_SIZE]string{}, errors.New(errMat.Error())
    }

    return current_board, nil
}

func generateMatrix(master_pos, player_pos []int) ([BOARD_SIZE]string, error) {
    var current_board [BOARD_SIZE]string
    for i := 0; i < BOARD_SIZE; i++ {
        if findInArray(i, master_pos) {
            current_board[i] = "O"
        } else if findInArray(i, player_pos) {
            current_board[i] = "X"
        } else {
            current_board[i] = "-"
        }
    }

    return current_board, nil
}

func drawBoard(currentBoard [BOARD_SIZE]string) {
    for i := 0; i < COL_SIZE; i++ {
        fmt.Println(currentBoard[i*COL_SIZE:i*COL_SIZE+3])
    }
}

func findInArray(value int, array[]int) (bool) {
    for x := 0; x < len(array); x++ {
        if value == array[x]{
            return true
        }
    }
    return false
}

func main() {
    log.Println("Starting tic-tac-toe game!")
    current_board, errFirst := firstMove()
    if errFirst != nil {
        log.Println(errFirst)
    }
    drawBoard(current_board)
    //startGame(current_board)
}
