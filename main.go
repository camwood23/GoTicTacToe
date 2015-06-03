//Need to add a win condition
//Change everything to work with a 2D array
package main

import (
    "fmt"
    "strconv"
)

var board [3][3]string

func main() {
    i, j, a, player := 0, 0, 0, 0
    for i < 3 {
        for j < 3 {
            board[i][j] = strconv.Itoa(a)
            j += 1
            a += 1
        }
        j = 0
        i += 1
    }

    for winCondition() {
        i, j = 0, 0
        printBoard()
        fmt.Println("\nPlayer " + strconv.Itoa(player + 1) + " pick a space")
        fmt.Scan(&a)
        fmt.Println()
        for i < 3 {
            for j < 3 {
                if board[i][j] == strconv.Itoa(a) {
                    if player == 0 {
                        board[i][j] = "X"
                    } else {
                        board[i][j] = "O"
                    }
                    player = (player + 1) % 2
                }
                j += 1
            }
            j = 0
            i += 1
        }
    }
}

func winCondition() bool {
    i, j := 0, 0
    for i < 3 {
        for j < 3 {
            j += 1
        }
        j = 0
        i += 1
    }
    return true
}

func printBoard() {
    i, j, a, q := 0, 0, 0, 0
    for i < 5 {
        if i % 2 == 0 {
            for j < 5 {
                if j % 2 == 0 {
                    fmt.Print(board[a][q]);
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
            q = 0
            a += 1
        }
        i += 1
    }
}
