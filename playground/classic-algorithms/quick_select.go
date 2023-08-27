package classicalgorithms

func quickSelect(nums []int, k int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		pivot := quickSelectPartition(nums, l, r)
		if pivot == k-1 {
			return nums[pivot]
		} else if pivot < k {
			l = pivot + 1
		} else {
			r = pivot - 1
		}
	}
	return -1
}

func quickSelectPartition(nums []int, l, r int) int {
	pviot := nums[r]
	pviotIdx := l

	for i := l; i < r; i++ {
		if nums[i] > pviot {
			nums[i], nums[pviotIdx] = nums[pviotIdx], nums[i]
			pviotIdx++
		}
	}

	nums[pviotIdx], nums[r] = nums[r], nums[pviotIdx]
	return pviotIdx
}
