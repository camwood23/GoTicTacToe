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
    printBoard()
    if player == 0 {
        fmt.Println("\nPlayer 2 won!")
    } else {
        fmt.Println("\nPlayer 1 won!")
    }
}

func winCondition() bool {
    for i := 0; i < 3; i++ {
        if board[i][0] == board[i][1] && board[i][0] == board[i][2] {
            return false
        }
        if board[0][i] == board[1][i] && board[0][i] == board[2][i] {
            return false
        }
    }
    if board[0][0] == board[1][1] && board[0][0] == board[2][2] {
        return false
    }
    if board[0][2] == board[1][1] && board[0][2] == board[2][0] {
        return false
    }
    return true
}

func printBoard() {
    a, q := 0, 0
    for i := 0; i < 5; i++ {
        if i % 2 == 0 {
            for j := 0; j < 5; j++ {
                if j % 2 == 0 {
                    fmt.Print(board[a][q]);
                    q += 1
                } else {
                    fmt.Print("|");
                }
            }
            fmt.Println();
        } else {
            fmt.Println("-----");
            q = 0
            a += 1
        }
    }
}
