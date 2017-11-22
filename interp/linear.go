package interp

// Linear performs a piecewise linear interpolation assuming that x1 and x2 are
// sorted in increasing order and do not contain any duplicate points. Slices
// y1 and x1 must be the same length. The returned slice will be the same
// length as x2.
func Linear(y1, x1, x2 []float64) []float64 {
	y2 := make([]float64, len(x2))
	s := subintervalIndices(x1, x2)
	m := slopes(x1, y1)
	for i := range y2 {
		y2[i] = y1[s[i]] + m[s[i]]*(x2[i]-x1[s[i]])
	}
	return y2
}

// Calculate the slope for each subinterval. The returned slice will be one
// less than the length of the given xs slice.
func slopes(xs, ys []float64) []float64 {
	m := make([]float64, len(xs)-1)
	for i := 0; i < len(xs)-1; i++ {
		m[i] = (ys[i+1] - ys[i]) / (xs[i+1] - xs[i])
	}
	return m
}

// Determine the index of the subinterval in x containing each value in z.
func subintervalIndices(x, z []float64) []int {
	indices := make([]int, len(z))
	for i := range indices {
		indices[i] = subintervalIndex(x, z[i])
	}
	return indices
}

// Locate the index of the subinterval in x containing z using binary search.
// The slice x is assumed to be monotonically increasing.
func subintervalIndex(x []float64, z float64) int {
	left := 0
	right := len(x)
	for right > left+1 {
		mid := (left + right) / 2
		if z < x[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	if left == len(x)-1 {
		return left - 1
	}
	return left
}
