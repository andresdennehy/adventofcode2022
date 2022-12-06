package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
)

func isWindowUnique(_window map[byte]int) bool {
    for _, freq := range _window {
        if freq > 1 {
            return false
        }
    }
    return true
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    line := scanner.Text()

    window := make(map[byte]int)
    for i:=0; i<4; i++ {
        window[line[i]]++
    }

    for i:=4; i<len(line); i++ {
        if isWindowUnique(window) {
            fmt.Println(i)
            break
        }
        window[line[i-4]]--
        window[line[i]]++
    }

    file.Close()
}
