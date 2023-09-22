package b75

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.



Example 1:

link: https://leetcode.com/problems/container-with-most-water/
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1


Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
Accepted
1,686,384
Submissions


*/

func ContainerWithMostWaterv29(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv28(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv27(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv26(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv25(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv24(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv23(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv22(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv21(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv20(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv19(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv18(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv17(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv16(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv15(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv14(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv13(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv12(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv11(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv10(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv9(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv8(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv7(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv6(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv5(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv4(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv3(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}

func ContainerWithMostWaterv2(height []int) int {
	l, r, result := 0, len(height)-1, 0

	for l < r {

		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return result
}

func ContainerWithMostWaterv1(height []int) int {
	var result int = 0

	totalLen := len(height)
	for i := 0; i < totalLen; i++ {
		for j := i + 1; j < totalLen; j++ {
			h := min(height[i], height[j])
			result = max(result, h*(j-i))
		}
	}

	return result
}

func ContainerWithMostWater(height []int) int {

	result := 0
	if len(height) == 0 {
		return result
	}

	n := len(height)
	l, r := 0, n-1

	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}

	return result
}
