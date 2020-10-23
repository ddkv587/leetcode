package findMedianSortedArrays

func findMedianSortedArrays( nums1 []int, nums2 []int ) float64 {
    mid := ( len( nums1 ) + len( nums2 ) ) >> 1
    carry := ( len( nums1 ) + len( nums2 ) ) & 0x01

    nums1 = append( nums1, 1000001)
    nums2 = append( nums2, 1000001)

    pos1 := 0
    pos2 := 0

    val1 := 0
    val2 := 0

    for i := 0; i <= mid; i++ {
        val1 = val2
        if nums1[pos1] > nums2[pos2] {
            val2 = nums2[pos2]
            pos2++
        } else {
            val2 = nums1[pos1]
            pos1++
        }
    }

    if carry == 1 {
        return float64( val2 )
    } else {
        return ( float64( val1 ) + float64( val2 ) ) / 2
    }
}