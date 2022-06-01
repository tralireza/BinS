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
