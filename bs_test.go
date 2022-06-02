package BinS

import (
	"log"
	"testing"
)

func init() {}

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
