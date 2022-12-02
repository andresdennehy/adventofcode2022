package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "strings"
)

func main() {
    file, err := os.Open("input.txt");
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file);

    // Initialize score
    score := 0

    for scanner.Scan() {
        play := strings.Fields(scanner.Text())
        switch play[0] {  // The shape they selected
            case "A":
                switch play[1] {
                    case "X":
                        score += 3 + 0
                    case "Y":
                        score += 1 + 3
                    case "Z":
                        score += 2 + 6
                }
            case "B":
                switch play[1] {
                    case "X":
                        score += 1 + 0
                    case "Y":
                        score += 2 + 3
                    case "Z":
                        score += 3 + 6
                }
            case "C":
                switch play[1] {
                    case "X":
                        score += 2 + 0
                    case "Y":
                        score += 3 + 3
                    case "Z":
                        score += 1 + 6
                }
        }
    }

    fmt.Println(score)

    file.Close()
}
