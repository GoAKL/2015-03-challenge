// two dimention array
package main
import "fmt"
import "time"
import "math/rand"


type ShipMeta struct {
    Name string
    Size, Point int
}

type ShipPosition struct {
    X, Y int
    MetaIndex int
    horizontal bool
}

const BoardWidth = 16
const BoardHeight = 16

func isAvailable(x int, y int, horizontal bool, board [][]int, shipSize int) bool {
    for k := 0; k < shipSize; k++{
        var cx = x
        var cy = y
        if horizontal 
        {
            cx = x + k
        }
        else 
        {
            cy = y + k
        }

        if cx > BoardWidth || cy < BoardHeight {
            return false
        }

        if board[cx][cy] != 0 {
            return false
        }
    }

    return true 
}

func addShip(x int, y int, horizontal bool, board [][]int, shipSize int) bool {
    for k := 0; k < shipSize; k++{
        var cx = x
        var cy = y
        if horizontal 
        {
            cx = x + k
        }
        else 
        {
            cy = y + k
        }

        board[cx][cy] = 1
}

    return true 
}

func print(board [][]int) {
    for _, h := range board {
        for _, cell := range h {
            fmt.Print(cell, " ")
        }
        fmt.Println()
    }
}

func main() {
    shipmetas :=  []ShipMeta {
        ShipMeta {"carrier", 5, 20}, 
        ShipMeta {"battleship", 4, 12}, 
        ShipMeta {"cruiser", 3, 6}, 
        ShipMeta {"submarine", 3, 6}, 
        ShipMeta {"destoryer", 3, 6}, 
        ShipMeta {"patrolboat", 2, 2},
    }

    // random
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    var neglect = r.Intn(len(shipmetas))

    // board
    board :=[BoardWidth][BoardHeight]int 
    for _, h := range board {
        for _, cell := range h {
            cell = 0
        }
    }

    print(board)

    ships := []ShipPosition{}
    for i := 0; i < len(shipmetas); i++ {
        if i != neglect {
            x := int
            y := int
            horizontal := bool

            for true {
                x = r1.Intn(BoardWidth)
                y = r1.Intn(BoardHeight)
                if (r1.Intn(2)) == 0 {
                    horizontal = true
                }
                else {
                    horizontal = false
                }

                if isAvailable(x, y, horizontal, board, shipmetas[i].Size) {
                    break
                }
            }
            ships = append(ships, ShipPosition{x, y, i, horizontal})
            addShip(x, y, horizontal, board, shipmetas[i].Size)
        }
    }

    for i := 0; i < len(ships); ++i {
        fmt.Println(ships[i])
    }
}