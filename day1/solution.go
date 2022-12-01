package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    file, err := os.Open("input.txt");
    check(err);
    defer file.Close();

    scanner := bufio.NewScanner(file);
    // Initialize counter
    sums := []int{};
    count := 0;

    for scanner.Scan() {
        if scanner.Text() == "" {
            sums = append(sums, count)
            count = 0
        } else {
            cals, err := strconv.Atoi(scanner.Text())
            check(err)
            count += cals
        }
    }

    // Solution 1
    max := 0
    for _, cals := range sums {
        if cals > max {
            max = cals
        }
    }
    fmt.Println(max)

    // Solution 2. Sorting, O(n log n)
    sort.Ints(sums)
    first_three := 0
    for _, val := range sums[len(sums)-3:] {
        first_three += val
    }
    fmt.Println(first_three)

}
