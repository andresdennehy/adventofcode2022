package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "strings"
    "strconv"
)

type File struct {
    size int
    name string
}

type Node struct {
    name string
    parent *Node
    children []*Node
    files []File
    size int
}

func setSizes(node *Node) {
    for _, child := range node.children {
        setSizes(child)
        node.size += child.size
    }

    for _, file := range node.files {
        node.size += file.size
    }
}

func sizeSum(node *Node) int {
    sum := 0
    stack := []*Node{node}

    for len(stack) > 0 {
        currentNode := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        if currentNode.size <= 100_000 {
            fmt.Println(currentNode.name, currentNode.size, currentNode.files, currentNode.children)
            sum += currentNode.size
        }
        for _, child := range currentNode.children {
            stack = append(stack, child)
        }
    }

    return sum
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
	    log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    root := Node{"/", nil, nil, nil, 0}
    pointer := &root
    // sum := 0

    for scanner.Scan() {
        cmds := strings.Split(scanner.Text(), " ")
        if cmds[1] == "cd" {
            // Move current directory
            switch cmds[2] {
                case "/":
                    pointer = &root
                case "..":
                    pointer = pointer.parent
                default:
                    for _, child := range pointer.children {
                        if child.name == cmds[2] {
                            pointer = child
                            break
                        }
                    }
            }
        } else if cmds[1] == "ls" {

        } else if cmds[0] == "dir" {
            // Create new dir
            newDir := &Node{cmds[1], pointer, nil, nil, 0}
            pointer.children = append(pointer.children, newDir)
        } else {
            // Else, it's a file
            filesize, _ := strconv.Atoi(cmds[0])
            pointer.files = append(pointer.files, File{filesize, cmds[1]})
        }
    }

    setSizes(&root)

    fmt.Println(sizeSum(&root))

    file.Close()
}
