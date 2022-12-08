package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "strconv"
)

// Brute force solution, no dynamic programming
func measureLeft(x, y int, forest [][]int) int {
    counter := 0
    for r:=y-1; r>=0; r-- {
        counter ++
        if forest[x][r] >= forest[x][y]  {
            break
        }
    }
    return counter
}


func measureRight(x, y int, forest [][]int) int {
    counter := 0
    for r:=y+1; r<len(forest[0]); r++ {
        counter++
        if forest[x][r] >= forest[x][y] {
            break
        }
    }
    return counter
}

func measureTop(x, y int, forest [][]int) int {
    counter := 0
    for r:=x-1; r>=0; r-- {
        counter++
        if forest[r][y] >= forest[x][y] {
            break
        }
    }
    return counter
}

func measureBottom(x, y int, forest [][]int) int {
    counter := 0
    for r:=x+1; r<len(forest); r++ {
        counter++
        if forest[r][y] >= forest[x][y] {
            break
        }
    }
    return counter
}

func treeScore(x, y int, forest [][]int) int {
    return measureLeft(x, y, forest) * measureRight(x, y, forest) * measureTop(x, y, forest) * measureBottom(x, y, forest)
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    forest := [][]int{}

    for scanner.Scan() {
        row := []int{}
        for _, char := range scanner.Text() {
            appendChar, _ := strconv.Atoi(string(char))
            row = append(row, appendChar)
        }
        forest = append(forest, row)
    }

    // Initialize
    width := len(forest[0])
    height := len(forest)
    highest := 0

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            if score := treeScore(x, y, forest); score > highest {
                highest = score
            }
        }
    }

    fmt.Println(highest)

    file.Close()
}
