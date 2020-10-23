package longestPalindrome

import _ "fmt"

func centerSpread( s string, left int, right int ) string {
    for left >= 0 && right < len( s ) {
        if s[ left ] != s[ right ] {
            break
        }

        left--
        right++
    }

    return string( s[ left + 1 : right ] )
}

func LongestPalindrome( s string ) string {
    maxStr := string( s[0] )

    for pos := 0; pos < len( s ) - 1; pos++ {
        oddStr  := centerSpread( s, pos, pos )
        evenStr := centerSpread( s, pos, pos + 1 )

        subStr := oddStr
        if len( oddStr ) < len( evenStr ) {
            subStr = evenStr
        }

        if len( subStr ) > len( maxStr ) {
            maxStr = subStr
        }
    }

    return maxStr
}