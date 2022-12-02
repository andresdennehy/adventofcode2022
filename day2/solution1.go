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
        switch play[1] {  // The shape we selected
            case "X":
                score += 1
                switch play[0] {
                    case "A":
                        score += 3
                    case "B":
                        score += 0
                    case "C":
                        score += 6
                }
            case "Y":
                score += 2
                switch play[0] {
                    case "A":
                        score += 6
                    case "B":
                        score += 3
                    case "C":
                        score += 0
                }
            case "Z":
                score += 3
                switch play[0] {
                    case "A":
                        score += 0
                    case "B":
                        score += 6
                    case "C":
                        score += 3
                }
        }
    }

    fmt.Println(score)

    file.Close()
}
