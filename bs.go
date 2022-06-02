package BinS

import (
	"log"
)

func init() {
	log.SetFlags(0)
	log.Print("> Binary Search")
}

type R786 struct{ n, d int }

func (o R786) Ratio(arr []int) float64 { return float64(arr[o.n]) / float64(arr[o.d]) }

// 786m K-th Smallest Prime Fraction
func kthSmallestPrimeFraction(arr []int, k int) []int {
	type R = R786

	l, r := float64(0), float64(1)
	for l < r {
		m := (l + r) / 2

		x, xRatio := 0, R{0, len(arr) - 1}
		d := 1
		for n := 0; n < len(arr)-1; n++ {
			for d < len(arr) && m <= (R{n, d}).Ratio(arr) {
				d++
			}
			if d == len(arr) {
				break
			}

			x += len(arr) - d
			if xRatio.Ratio(arr) < (R{n, d}).Ratio(arr) {
				xRatio = R{n, d}
			}
		}
		log.Printf("l:%g   %d %d   m: %g   r:%g   -> %v", l, x, k, m, r, xRatio)

		if k == x {
			return []int{arr[xRatio.n], arr[xRatio.d]}
		}
		if x > k {
			r = m
		} else {
			l = m
		}
	}

	return []int{}
}

// 1208m Get Equal Substrings Within Budget
func equalSubstring(s string, t string, maxCost int) int {
	Cost := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		Cost[i] = int(s[i]) - int(t[i])
		if Cost[i] < 0 {
			Cost[i] *= -1
		}
	}

	lMax := 0

	curCost, l := 0, 0
	for r := range len(s) {
		curCost += Cost[r]
		for curCost > maxCost {
			curCost -= Cost[l]
			l++
		}

		lMax = max(lMax, r-l+1)
		log.Printf(" -> %d", lMax)
	}

	return lMax
}
