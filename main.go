//Need to add a win condition
package main

import (
    "fmt"
    "strconv"
)

var board [9]string

func main() {
    i := 0
    for i < 9 {
        board[i] = strconv.Itoa(i)
        i += 1
    }

    player := 0
    //Game loop, need a win condition to break loop
    for 1 > 0 {
        printBoard()
        fmt.Println("\nPlayer " + strconv.Itoa(player + 1) + " pick a space")
        fmt.Scan(&i)
        fmt.Println()
        if board[i] != "X" && board[i] != "O" {
            if player == 0 {
                board[i] = "X"
            } else {
                board[i] = "O"
            }
            player = (player + 1) % 2
        }
    }
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
