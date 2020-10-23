package reverseNumber

func reverse(x int) int {
    neg := 1
    if x < 0 {
        x *= -1;
        neg = -1
    }

    ret := int64(0)
    for x != 0 {
        ret = 10 * ret + int64( x % 10 )
        x = x / 10
    }

    if ret > 2147483647 {
        return 0
    }
    return neg * int( ret )
}