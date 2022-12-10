package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    // "strconv"
)

func check(counter *int, x int, signalSum *int) {
    if (*counter - 20) % 40 == 0 && *counter <= 220 {
        *signalSum += x * *counter
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

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%s %d", &action, &amount)
        switch action {
            case "noop":
                check(&counter, x, &signalSum)
            case "addx":
                check(&counter, x, &signalSum)
                check(&counter, x, &signalSum)
                x += amount
        }
    }

    fmt.Println(signalSum)

    file.Close()
}
