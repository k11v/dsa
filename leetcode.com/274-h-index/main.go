package main

func main() {}

func hIndexV3(citations []int) int {
    n := len(citations)

    // Calculate frequencies.
    papers := make([]int, n+1)
    for _, c := range citations {
        papers[min(c, n)]++
    }

    // Calculate suffix sums.
    for i, _ := range slices.Backward(papers) {
        if i == n {
            continue
        }
        papers[i] += papers[i+1]
    }

    // Calculate max h.
    maxH := 0
    for h, p := range papers {
        if p >= h {
            maxH = h
        }
    }

    return maxH
}

func hIndexV2(citations []int) int {
    slices.Sort(citations)

    n := len(citations)

    maxH := 0
    for i := range n {
        h := min(n - i, citations[i])
        maxH = max(h, maxH)
    }

    return maxH
}

func hIndex(citations []int) int {
    maxH := 0

    for h := range slices.Max(citations) + 1 {
        c := 0
        for _, paperCitations := range citations {
            if paperCitations >= h {
                c++
            }
        }
        if c >= h {
            maxH = max(h, maxH)
        }
    }

    return maxH
}
