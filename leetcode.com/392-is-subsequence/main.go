package main

func main() {}

func isSubsequence(s string, t string) bool {
    j := 0

    for i := 0; i < len(s); i++ {
        for j < len(t) && s[i] != t[j] {
            j++
        }
        if j == len(t) {
            return false
        }
        j++
    }

    return true
}
