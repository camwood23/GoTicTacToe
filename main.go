//Need to incorporate user input and a win condition
package main

import (
    "fmt"
    "strconv"
)

//Need this to be a String so it can store X's and O's
var board [9]string


func main() {
    i := 0
    for i < 9 {
        //Need to figure out how to convert ints into strings for this line
        board[i] = strconv.Itoa(i)
        i += 1
    }
    for 1 > 0 {
        printBoard()
        fmt.Scan(&i)
        board[i] = "X"
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
