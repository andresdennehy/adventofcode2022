package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    //"strconv"
)

type Vertex struct {
    x int
    y int
}

func canMove(from, to byte) bool {
	if from == 'S' {
		from = 'a'
	}
	if from == 'z' && to == 'E' {
		return true
	} else if to == 'E' {
		return false
	}
	return int(to)-int(from) <= 1
}

func bfs(_heightMap []string, start Vertex) int {
    // Solve with BFS
    visited := make(map[Vertex]int)
    queue := []Vertex{}
    queue = append(queue, start)
    visited[start] = 0

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        v := visited[curr]

        next := []Vertex{}
        up := Vertex{curr.x, curr.y - 1}
        down := Vertex{curr.x, curr.y + 1}
        left := Vertex{curr.x - 1, curr.y}
        right := Vertex{curr.x + 1, curr.y}

        // Push up
        if up.y >= 0 && canMove(_heightMap[curr.y][curr.x], _heightMap[up.y][up.x]) {
            next = append(next, up)
        }

        // Push down
        if down.y < len(_heightMap) && canMove(_heightMap[curr.y][curr.x], _heightMap[down.y][down.x]) {
            next = append(next, down)
        }

        // Push left
        if left.x >= 0 && canMove(_heightMap[curr.y][curr.x], _heightMap[left.y][left.x]) {
            next = append(next, left)
        }

        // Push right
        if right.x < len(_heightMap[0]) && canMove(_heightMap[curr.y][curr.x], _heightMap[right.y][right.x]) {
            next = append(next, right)
        }

        for _, n := range next {
            if _heightMap[n.y][n.x] == 'E' {
                return v + 1
            }
            if _, ok := visited[n]; !ok {
                visited[n] = v + 1
                queue = append(queue, n)
            }
        }
    }

    return -1
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    heightMap := []string{}

    // Parse
    for scanner.Scan() {
        heightMap = append(heightMap, scanner.Text())
    }

    // Make it a big enough number
    smallest := len(heightMap) * len(heightMap[0])

    for row, line := range heightMap {
        for col, char := range line {
            if string(char) == "a" {
                if distance := bfs(heightMap, Vertex{col, row}); distance < smallest && distance > 0 {
                    fmt.Println("Smallest is", col, row)
                    smallest = distance
                }
           }
        }
    }

    fmt.Println(smallest)

    file.Close()
}
