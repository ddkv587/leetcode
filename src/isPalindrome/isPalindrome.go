package isPalindrome

func isPalindrome(x int) bool {
    if x < 0 || ( x != 0 && x % 10 == 0 ) {
        return false
    }

    compare := int64( 0 )
    val := int64( x )
    for val != 0 {
        compare = compare * 10 + val % 10
        val = val / 10
    }

    return int( compare ) == x
}