package main

func main() {}

func productExceptSelfV2(nums []int) []int {
    n := len(nums)

    // Calculate prefix products.
    result := make([]int, n)
    for i := 0; i < n; i++ {
        if i == 0 {
            result[i] = nums[i]
        } else {
            result[i] = result[i-1] * nums[i]
        }
    }

    // Calculate products without self.
    suffixProduct := 1
    for i := n - 1; i >= 0; i-- {
        prefixProductWithoutSelf := 1
        if i-1 >= 0 {
            prefixProductWithoutSelf = result[i-1]
        }

        suffixProductWithoutSelf := suffixProduct

        suffixProduct *= nums[i]

        result[i] = prefixProductWithoutSelf * suffixProductWithoutSelf
    }

    return result
}

func productExceptSelf(nums []int) []int {
    prefixProducts := make([]int, len(nums))
    for i, num := range nums {
        if i == 0 {
            prefixProducts[i] = num
        } else {
            prefixProducts[i] = prefixProducts[i-1] * num
        }
    }

    suffixProducts := make([]int, len(nums))
    for i, num := range slices.Backward(nums) {
        if i == len(nums) - 1 {
            suffixProducts[i] = num
        } else {
            suffixProducts[i] = num * suffixProducts[i+1]
        }
    }

    result := make([]int, len(nums))
    for i := range len(nums) {
        result[i] = 1
        if i-1 >= 0 {
            result[i] *= prefixProducts[i-1]
        }
        if i+1 < len(nums) {
            result[i] *= suffixProducts[i+1]
        }
    }

    return result
}
