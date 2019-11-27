package main

import (
    "log"
    "math/rand"
    "time"
)

func fillLine (pos []int) (int, int) {
    log.Println("Filling line!")
    var master_row, master_col int
    if pos[0] < COL_SIZE {
        master_row = rand.Intn(COL_SIZE)
        master_col = pos[0]
    } else if (pos[0] >= COL_SIZE) && (pos[0] < COL_SIZE*2) {
        master_row = pos[0]-COL_SIZE
        master_col = rand.Intn(COL_SIZE)
    } else if pos[0] == 2*COL_SIZE {
        master_row = rand.Intn(COL_SIZE)
        master_col = master_row
    } else if pos[0] == 2*COL_SIZE+1 {
        master_row = rand.Intn(COL_SIZE)
        master_col = COL_SIZE-1-master_row
    }

    return master_row, master_col
}

func selectCorner () (int, int) {
    rand.Seed(time.Now().UnixNano())
    index := []int{
            0,
            COL_SIZE-1}

    return index[rand.Intn(len(index))], index[rand.Intn(len(index))]
}

func selectNextMove (board_status BoardProperties) (int, int) {
    var master_row, master_col int
    if len(board_status.MasterPositions) == 0 {
        master_row, master_col = selectCorner()
    } else {
        master_row = rand.Intn(COL_SIZE)
        master_col = rand.Intn(COL_SIZE)
    }

    return master_row, master_col
}

func applyStrategy (board_status BoardProperties) (int, int) {
    rand.Seed(time.Now().UnixNano())
    board := board_status.WinningBoard[:]
    pos := findIntInArray((COL_SIZE-1), board)
    pos_opp := findIntInArray(-(COL_SIZE-1), board)
    opportunity := findIntInArray(0, board)
    var master_row, master_col int
    if len(pos) > 0 {
        master_row, master_col = fillLine(pos)
    } else if len(pos_opp) > 0{
        master_row, master_col = fillLine(pos_opp)
    } else if len(opportunity) > 0 {
        master_row, master_col = fillLine(opportunity)
    } else {
        master_row, master_col = selectNextMove(board_status)
    }

    return master_row, master_col
}

