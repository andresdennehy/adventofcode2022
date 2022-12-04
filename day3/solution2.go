package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "unicode"
)

func create_charmap(items string) map[rune]bool {
    set := make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return set
}

func common_item(first string, second string, third string) int {
    first_charmap := create_charmap(first)
    second_charmap := create_charmap(second)
    third_charmap := create_charmap(third)

    var result int
    for char := range first_charmap {
        if second_charmap[char] && third_charmap[char] {
            result = int(unicode.ToLower(char)) - 96
            if unicode.IsUpper(char) {
                result += 26
            }
            break
        }
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
        first := scanner.Text()
        scanner.Scan()
        second := scanner.Text()
        scanner.Scan()
        third := scanner.Text()
        priority += common_item(first, second, third)
    }

    fmt.Println(priority)

    file.Close()
}
