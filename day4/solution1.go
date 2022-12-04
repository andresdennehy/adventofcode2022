package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "strconv"
    "strings"
)


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file);

    // Initialize
    var count int

    for scanner.Scan() {
        pairs := strings.Split(scanner.Text(), ",")
        first := make([]int, 2)
        second := make([]int, 2)

        // Convert to int
        for i, s := range strings.Split(pairs[0], "-") {
            first[i], _ = strconv.Atoi(s)
        }

        for i, s := range strings.Split(pairs[1], "-") {
            second[i], _ = strconv.Atoi(s)
        }

        // Second contains first or first contains second
        if (first[0] >= second[0] && first[1] <= second[1]) || (second[0] >= first[0] && second[1] <= first[1]) {
            fmt.Println(pairs, "are contained")
            count++
        }
    }

    fmt.Println(count)

    file.Close()
}
