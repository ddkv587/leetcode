package twoSum

func TwoSum( src []int, target int ) []int {
    sub := make( map[int]int )
    for i := 0; i < len(src); i++ {
        index, found := sub[ src[i] ]

        if found && index != i {
            if index < i {
                return []int{ index, i }
            } else {
                return []int{ i, index }
            }
        } else {
            sub[ target - src[i] ] = i
        }
    }

    return []int{}
}