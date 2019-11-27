package main

import (
)

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

