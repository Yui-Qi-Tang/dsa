package b75

/*
162. Find Peak Element

A peak element is an element that is strictly greater than its neighbors.

Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.

You must write an algorithm that runs in O(log n) time.
*/

func findPeakElementv23(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv22(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv21(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv20(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv19(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func findPeakElementv18(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv17(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv16(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv15(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv14(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}

	}

	return l
}

func findPeakElementv13(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv12(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}

	}

	return l
}

func findPeakElementv11(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv10(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}

	}

	return l
}

func findPeakElementv9(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv8(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv7(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func findPeakElementv6(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv5(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}

	}

	return l
}

func findPeakElementv4(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv3(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {

		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}

	}

	return l

}

func findPeakElementv2(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func findPeakElementv1(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func FindPeakElement(nums []int) int {

	// a peak: exist a index m which satififies nums[m] > nums[m-1] && nums[m] > nums[m+1]
	l, r := 0, len(nums)-1 // l stores the 'first' peak of nums

	for l < r {
		m := (l + r) / 2
		// nums[m] > nums[m+1]
		if nums[m] > nums[m+1] { // m is the peak of m....r , we don't care nums[m+1,...r], so we hust care about nums[l,....,m]
			r = m
		} else { // m+1 is the peak of l...m, so, we don't care nums[l,....m], so we just care about nums[m+1,...,r]
			// let m = m+1 which satifiese nums[m] > nums[m-1]
			l = m + 1
		}
	}
	return l
}

func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
