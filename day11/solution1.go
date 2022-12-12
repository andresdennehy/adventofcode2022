package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strconv"
)

type Monkey struct {
    items []int
    operation func(int)int
    divisibleBy int
    monkeyIfTrue int
    monkeyIfFalse int
    inspected int
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    monkeys := []*Monkey{}

    // Parse all monkeys
    for scanner.Scan() {
        if len(scanner.Text()) > 0 && scanner.Text()[:6] == "Monkey" {
            monkey := Monkey{}

            // Scan items
            scanner.Scan()
            re := regexp.MustCompile("[0-9]+")
			for _, s := range re.FindAllString(scanner.Text(), -1) {
				n, _ := strconv.ParseInt(s, 0, 0)
				monkey.items = append(monkey.items, int(n))
			}

			// Scan operation
			scanner.Scan()
			var operation string
			var amount string
			fmt.Sscanf(scanner.Text(), "  Operation: new = old %s %s", &operation, &amount)
			if amount == "old" {
			    switch operation {
                    case "*":
                        monkey.operation = func(old int) int {
                            return old * old
                        }
                    case "+":
                        monkey.operation = func(old int) int {
                            return old + old
                        }
                }
			} else {
			    convertedAmount, _ := strconv.Atoi(amount)
                switch operation {
                    case "*":
                        monkey.operation = func(old int) int {
                            return old * convertedAmount
                        }
                    case "+":
                        monkey.operation = func(old int) int {
                            return old + convertedAmount
                        }
                }
            }

            // Scan divisor
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &monkey.divisibleBy)

			// Scan throwing
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &monkey.monkeyIfTrue)
			scanner.Scan()
            fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &monkey.monkeyIfFalse)

			monkeys = append(monkeys, &monkey)
        }
    }

    // Run rounds
    for i:=0; i<20; i++ {
        for _, monkey := range monkeys {
            for _, item := range monkey.items {
                item = monkey.operation(item)
                item = item / 3
                if item % monkey.divisibleBy == 0 {
                    monkeys[monkey.monkeyIfTrue].items = append(monkeys[monkey.monkeyIfTrue].items, item)
                } else {
                    monkeys[monkey.monkeyIfFalse].items = append(monkeys[monkey.monkeyIfFalse].items, item)
                }
                monkey.inspected++
            }
            monkey.items = []int{}
        }
    }

    sort.Slice(monkeys, func(i, j int) bool {
        return monkeys[i].inspected > monkeys[j].inspected
    })

    fmt.Println(monkeys[0].inspected * monkeys[1].inspected)

    file.Close()
}
