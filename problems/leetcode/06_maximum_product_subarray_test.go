package b75

import "testing"

/*
Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

A subarray is a contiguous subsequence of the array.

Example 1:

Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.


Constraints:

1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

*/

func TestMaximumProductSubarray(t *testing.T) {

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{2, 3, -2, 4},
			want: 6,
		},
		{
			in:   []int{-2, 0, -1},
			want: 0,
		},
		{
			in:   []int{-2, -1, 0},
			want: 2,
		},

		{
			in:   []int{-2, -1, 0, 3, 4},
			want: 12,
		},
		{
			in:   []int{-2, 3, -4},
			want: 24,
		},
	}

	testfuncs := []func([]int) int{
		MaximumProductSubarrayv10,
		MaximumProductSubarrayv9,
		MaximumProductSubarrayv8,
		MaximumProductSubarrayv7,
		MaximumProductSubarrayv6,
		MaximumProductSubarrayv5,
		MaximumProductSubarrayv4,
		MaximumProductSubarrayv3,
		MaximumProductSubarrayv2,
		MaximumProductSubarrayv1,
		MaximumProductSubarray,
	}

	for i, f := range testfuncs {
		t.Logf("test function: %d", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
