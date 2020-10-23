package convertZ

import _ "fmt"

func Convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    cellSize := numRows * 2 - 2
    ret := make( [][]byte, numRows )

    for i := 0; i < len(s); i++ {
        index := i % cellSize

        if index >= numRows {
            index = ( 2 * ( numRows - 1 ) ) - index
        }

        ret[index] = append( ret[index], s[i] )
    }

    retStr := string("")
    for _, str := range ret {
        retStr += string( str )
    }

    return retStr
}