package main

func main() {}

func canConstruct(ransomNote string, magazine string) bool {
    countFromRune := make(map[rune]int)
    for _, r := range magazine {
        countFromRune[r]++
    }

    for _, r := range ransomNote {
        countFromRune[r]--
        if countFromRune[r] < 0 {
            return false
        }
    }

    return true
}
