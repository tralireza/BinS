package BinS

import (
	"log"
	"testing"
)

func init() {}

// 209m Minimum Size Subarray Size
func Test209(t *testing.T) {
	log.Print("2 ?= ", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	log.Print("1 ?= ", minSubArrayLen(4, []int{1, 4, 4}))
	log.Print("0 ?= ", minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

// 240m Search a 2D Matrix II
func Test240(t *testing.T) {
	lBS := func(nums []int, k int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] >= k {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}
	for _, k := range []int{0, 1, 3, 6, 7} {
		log.Printf("%d -> %d", k, lBS([]int{1, 1, 3, 3, 4, 5, 5, 6}, k))
	}

	log.Print("true ?= ", searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	log.Print("false ?= ", searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
	log.Print("true ?= ", searchMatrix([][]int{{3, 3, 8, 13, 13, 18}, {4, 5, 11, 13, 18, 20}, {9, 9, 14, 15, 23, 23}, {13, 18, 22, 22, 25, 27}, {18, 22, 23, 28, 30, 33}, {21, 25, 28, 30, 35, 35}, {24, 25, 33, 36, 37, 40}}, 21))
}

// 633m Sum of Square Numbers
func Test633(t *testing.T) {
	log.Print("true ?= ", judgeSquareSum(5))
	log.Print("false ?= ", judgeSquareSum(3))
	log.Print("false ?= ", judgeSquareSum(999999999))
}

// 786m K-th Smallest Prime Fraction
func Test786(t *testing.T) {
	log.Print("[2 5] ?= ", kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	log.Print("[1 7] ?= ", kthSmallestPrimeFraction([]int{1, 7}, 1))
	log.Print("[1 47] ?= ", kthSmallestPrimeFraction([]int{1, 29, 47}, 1))
	log.Print("[13 17] ?= ", kthSmallestPrimeFraction([]int{1, 13, 17, 59}, 6))
}

// 1208m Get Equal Substrings Within Budget
func Test1208(t *testing.T) {
	log.Print("3 ?= ", equalSubstring("abcd", "bcdf", 3))
	log.Print("1 ?= ", equalSubstring("abcd", "cdef", 3))
	log.Print("2 ?= ", equalSubstring("abce", "adce", 0))
}
