package myAtoi

import _ "fmt"

func MyAtoi(s string) int {
    neg := int64(1)
    val := int64(0)
    start := true
    for i := 0; i < len(s) ; i++ {
        if s[i] == ' ' && start {
            continue
        }

        if s[i] == '-' && start {
            neg = -1
            start = false
            continue
        } else if s[i] == '+' && start {
            start = false
            continue
        }

        start = false
        if s[i] < '0' || s[i] > '9' {
            break;
        }

        val = val * 10 + int64( s[i] - '0' )
        if val > 2147483648 {
            break
        }
    }
    val = neg * val

    if val > 2147483647 {
        return 2147483647
    } else if val < -2147483648 {
        return -2147483648
    }

    return int(val)
}