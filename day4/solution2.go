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

        // If not strictly before or after
        if !(first[1] < second[0] || first[0] > second[1]) {
            fmt.Println(pairs, "overlaps")
            count++
        }
    }

    fmt.Println(count)

    file.Close()
}
