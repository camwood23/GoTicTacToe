package main

import "fmt"

var board [9]int


func main() {
    i := 0
    for i < 9 {
        board[i] = i
        i += 1
    }
    printBoard()
}

func printBoard() {
    i := 0
    j := 0
    q := 0
    for i < 5 {
        if i % 2 == 0 {
            for j < 5 {
                if j % 2 == 0 {
                    fmt.Print(board[q]);
                    q += 1
                } else {
                    fmt.Print("|");
                }
                j += 1
            }
            j = 0
            fmt.Println();
        } else {
            fmt.Println("-----");
        }
        i += 1
    }
}
