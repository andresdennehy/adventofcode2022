package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "strconv"
)

// Brute force solution, no dynamic programming
func isVisibleLeft(x, y int, forest [][]int) bool {
    for r:=0; r<y; r++ {
        if forest[x][r] >= forest[x][y]  {
            return false
        }
    }
    return true
}


func isVisibleRight(x, y int, forest [][]int) bool {
    for r:=y+1; r<len(forest[0]); r++ {
        if forest[x][r] >= forest[x][y] {
            return false
        }
    }
    return true
}

func isVisibleTop(x, y int, forest [][]int) bool {
    for r:=0; r<x; r++ {
        if forest[r][y] >= forest[x][y] {
            return false
        }
    }
    return true
}

func isVisibleBottom(x, y int, forest [][]int) bool {
    for r:=x+1; r<len(forest); r++ {
        if forest[r][y] >= forest[x][y] {
            return false
        }
    }
    return true
}

func isVisible(x, y int, forest [][]int) bool {
    if isVisibleLeft(x, y, forest) {
        return true
    }
    if isVisibleRight(x, y, forest) {
        return true
    }
    if isVisibleTop(x, y, forest) {
        return true
    }
    if isVisibleBottom(x, y, forest) {
        return true
    }
    return false
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    forest := [][]int{}

    // Convert to [][]int
    for scanner.Scan() {
        row := []int{}
        for _, char := range scanner.Text() {
            appendChar, _ := strconv.Atoi(string(char))
            row = append(row, appendChar)
        }
        forest = append(forest, row)
    }

    width := len(forest[0])
    height := len(forest)

    // Initialize with border trees as visible
    visible := 2 * width + 2 * height - 4

    for x := 1; x < width-1; x++ {
        for y := 1; y < height-1; y++ {
            if isVisible(x, y, forest) {
                fmt.Println(x, y, "is visible")
                visible += 1
            }
        }
    }

    fmt.Println(visible)

    file.Close()
}
