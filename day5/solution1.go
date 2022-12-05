package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "strings"
)


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file);

    stacks := make([]string, 9)
    for scanner.Scan() {
        // Parse containers
        line := scanner.Text()
        if string(line[1]) == "1" {
            break
        } else {
            for i:=0; i<9; i++ {
                if char := string(line[1 + 4 * i]); char != "" {
                    stacks[i] = char + stacks[i]
                }
            }
        }
    }

    // Remove empty spaces
    for i:=0; i<9; i++ {
        stacks[i] = strings.Trim(string(stacks[i]), " ")
    }

    // Skip empty line
    scanner.Scan()

    // Keep parsing instructions
    var i int
    for scanner.Scan() {
        i++
        var count, from, to int
        _, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &count, &from, &to)
        from = from - 1
        to = to - 1
        if err == nil {
            for i:=0; i<count; i++ {
                char := stacks[from][len(stacks[from])-1]
                stacks[from] = stacks[from][:len(stacks[from])-1]
                stacks[to] += string(char)
            }
        }
    }

    // Show first letter of every stack
    var result string
    for i:=0; i<9; i++ {
        result += string(stacks[i][len(stacks[i])-1])
    }

    fmt.Println(result)

    file.Close()
}
