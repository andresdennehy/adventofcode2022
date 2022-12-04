package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "unicode"
)

func common_item(first string, second string) int {
    var item rune
    charmap := make(map[rune]bool)
    for _, char := range first {
        charmap[char] = true
    }
    for _, char := range second {
        if _, found := charmap[char]; found {
            item = char
            break
        }
    }

    var result int
    result = int(unicode.ToLower(item)) - 96
    if unicode.IsUpper(item) {
        result += 26
    }

    return result
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file);

    // Initialize
    priority := 0

    for scanner.Scan() {
        rucksack := scanner.Text()
        priority += common_item(rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:])
    }

    fmt.Println(priority)

    file.Close()
}
