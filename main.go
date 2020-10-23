package main

import ( 
    "twoSum"
    "addTwoNumbers"
    "longestPalindrome"
    "convertZ"
    "myAtoi"
    "isMatch"
    "fmt"
)

func main() {
    val := []int{ 7, 2, 11, 15 }
    ret := twoSum.TwoSum( val, 9 )

    fmt.Println( ret )

    add1 := addTwoNumbers.ListNode{ Val: 0, Next: nil }
    add2 := addTwoNumbers.ListNode{ Val: 1, Next: nil }
    sum  := addTwoNumbers.AddTwoNumbers( &add1, &add2 )

    fmt.Println( sum.Val )

    fmt.Println( longestPalindrome.LongestPalindrome( "babaddtattarrattatddetartrateedredividerb" ) )
    fmt.Println( longestPalindrome.LongestPalindrome( "babab" ) )
    fmt.Println( longestPalindrome.LongestPalindrome( "cbbd" ) )
    fmt.Println( longestPalindrome.LongestPalindrome( "a" ) )
    fmt.Println( longestPalindrome.LongestPalindrome( "ccc" ) )
    fmt.Println( longestPalindrome.LongestPalindrome( "abb" ) )
    fmt.Println( longestPalindrome.LongestPalindrome( "bba" ) )
    fmt.Println( longestPalindrome.LongestPalindrome( "caba" ) )

    //fmt.Println( value.testValue1 )
    fmt.Println( child.TestValue2 )
    fmt.Println( parent.TestValue2 )

    fmt.Println( parent.GetValue1() )
    fmt.Println( parent.GetValue2() )

    fmt.Println( convertZ.Convert( "PAYPALISHIRING", 4 ) )
    fmt.Println( "PINALSIGYAHRPI" )

    channel.TestChannel()

    fmt.Println( myAtoi.MyAtoi( "9223372036854775808" ) )

    // fmt.Println( isMatch.IsMatch( "aaa", "a*aaa" ) )
    // fmt.Println( isMatch.IsMatch( "aa", "a" ) )
    // fmt.Println( isMatch.IsMatch( "aa", "a*" ) )
    // fmt.Println( isMatch.IsMatch( "ab", ".*" ) )
    // fmt.Println( isMatch.IsMatch( "aab", "c*a*b" ) )
    // fmt.Println( isMatch.IsMatch( "mississippi", "mis.*i.*s*ip*." ) )
    // fmt.Println( isMatch.IsMatch( "aaa", "ab*a*c*a" ) )
    // fmt.Println( isMatch.IsMatch( "aaca", "ab*a*c*a" ) )
    // fmt.Println( isMatch.IsMatch( "a", ".*.." ) )
    // fmt.Println( isMatch.IsMatch( "ab", ".*" ) )
    // fmt.Println( isMatch.IsMatch( "acaabbaccbbacaabbbb", "a*.*b*.*a*aa*a*" ) )
    // fmt.Println( isMatch.IsMatch( "bbcacbabbcbaaccabc", "b*a*a*.c*bb*b*.*.*" ) )
    // fmt.Println( isMatch.IsMatch( "b", "aaa." ) )

    fmt.Println( isMatch.IsMatch( "bbbcacbaacacaaaba", "b*c*b*.a.*a*.*.*b*" ) )
    fmt.Println( isMatch.IsMatch( "mississippi", "mis*is*ip*." ) )
    fmt.Println( isMatch.IsMatch( "abacacccbbbcbcbb", ".*.*.*ab*.*ab.*c*" ) )
}