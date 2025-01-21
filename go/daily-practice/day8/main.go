package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
    var s int
    var nums [4]int
    var total int
    for i := 0; i < n; i++ {
        fmt.Scan(&s)
        nums[s-1]++
    }
    if nums[2] > 0 || nums[1] & 1 > 0 {
        nums[0] = max(nums[0]-nums[2], 0)
        nums[0] = max(nums[0]-(nums[1] & 1)*2, 0)
        nums[0] = (nums[0]+4-1)/4
    } else {
        nums[0] = (nums[0]+4-1)/4
    }
    nums[1] = (nums[1]+2-1)/2
    total = nums[3]+nums[2]+nums[1]+nums[0]
    fmt.Println(total)
}
