package main

func main() {}

func wordPattern(pattern string, s string) bool {
    wordFromRune := make(map[rune]string)
    runeFromWord := make(map[string]rune)

    j := 0

    for _, r := range pattern {
        if j == len(s) {
            return false
        }

        newJ := strings.Index(s[j:], " ")
        if newJ != -1 {
            newJ = j + newJ
        } else {
            newJ = len(s)
        }

        sw := s[j:newJ]
        
        j = newJ
        if j < len(s) {
            j++
        }

        pw := wordFromRune[r]

        if pw == "" {
            if _, used := runeFromWord[sw]; used {
                return false
            }
            wordFromRune[r] = sw
            runeFromWord[sw] = r
            continue
        }

        if sw == pw {
            continue
        }

        return false
    }

    return j == len(s)
}
