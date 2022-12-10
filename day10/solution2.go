package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    // "strconv"
)

type Screen [6][40]bool


func check(counter *int, x int, screen *Screen) {
    c := *counter - 1
    row, col := c/40, c%40
	if col >= x-1 && col <= x+1 {
		(*screen)[row][col] = true
	}
	*counter++
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    // Initialize
    var action string
    var amount int
    x := 1
    counter := 1
    signalSum := 0
    screen := Screen{}

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%s %d", &action, &amount)
        switch action {
            case "noop":
                check(&counter, x, &screen)
            case "addx":
                check(&counter, x, &screen)
                check(&counter, x, &screen)
                x += amount
        }
    }

    for row := 0; row < len(screen); row++ {
		s := ""
		for col := 0; col < len(screen[row]); col++ {
			if screen[row][col] {
				s += "#"
			} else {
				s += "."
			}
		}
		fmt.Println(s)
	}

    fmt.Println(signalSum)

    file.Close()
}
