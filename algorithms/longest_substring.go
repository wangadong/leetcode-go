package algorithms

func LengthOfLongestSubstring(s string) int {
    max := 0 
    start := 0
    m := make(map[rune]int)
    for i, c := range s {
        m[c]++
        if m[c] == 1 {
            max = maxInt(max, i-start+1)
        } else {
            for m[c] > 1 {
                m[[]rune(s)[start]]--
                                start++

            }
        }
    }
    return max
}

func maxInt(a int, b int) int {
    if a>b {
        return a
    } else {
        return b
    }
}
