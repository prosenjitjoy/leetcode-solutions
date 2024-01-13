func strStr(haystack string, needle string) int {
    N := len(needle)
    LPS := make([]int, N)

    for prevLPS, i := 0, 1; i<N; {
        if needle[i] == needle[prevLPS] {
            LPS[i] = prevLPS + 1
            prevLPS++
            i++
        } else {
            if prevLPS == 0 {
                LPS[i] = 0
                i++
            } else {
                prevLPS = LPS[prevLPS - 1]
            }
        }
    }

    for i, j := 0, 0; i<len(haystack); {
        if haystack[i] == needle[j] {
            i++
            j++
        } else {
            if j == 0 {
                i++
            } else {
                j = LPS[j-1]
            }
        }

        if j == N {
            return i - N
        }
    }

    return -1
}