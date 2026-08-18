// Package safemath provides saturating arithmetic at representation boundaries.
package safemath

import "math"

// MulInt multiplies an int by a non-negative int64 unit and saturates on overflow.
func MulInt(value int, unit int64) int64 {
	return MulInt64(int64(value), unit)
}

// MulInt64 multiplies an int64 by a non-negative unit and saturates on overflow.
func MulInt64(value, unit int64) int64 {
	if unit <= 0 || value == 0 {
		return value * unit
	}
	if value > math.MaxInt64/unit {
		return math.MaxInt64
	}
	if value < math.MinInt64/unit {
		return math.MinInt64
	}
	return value * unit
}

// MulFloat64 multiplies a float64 by a non-negative int64 unit, truncates the
// result like an ordinary integer conversion, and saturates on overflow.
func MulFloat64(value float64, unit int64) int64 {
	if math.IsNaN(value) {
		return 0
	}
	product := value * float64(unit)
	if product >= float64(math.MaxInt64) {
		return math.MaxInt64
	}
	if product <= float64(math.MinInt64) {
		return math.MinInt64
	}
	return int64(product)
}

// Float64ToInt truncates value toward zero and saturates at host-int bounds.
func Float64ToInt(value float64) int {
	if math.IsNaN(value) {
		return 0
	}
	min := float64(math.MinInt)
	maxExclusive := -min
	if value >= maxExclusive {
		return math.MaxInt
	}
	if value <= min {
		return math.MinInt
	}
	return int(value)
}

// NegateInt64 negates value and maps the unrepresentable -MinInt64 to MaxInt64.
func NegateInt64(value int64) int64 {
	if value == math.MinInt64 {
		return math.MaxInt64
	}
	return -value
}

// AddInt64 adds values in order and saturates instead of wrapping.
func AddInt64(values ...int64) int64 {
	var total int64
	for _, value := range values {
		if value > 0 && total > math.MaxInt64-value {
			return math.MaxInt64
		}
		if value < 0 && total < math.MinInt64-value {
			return math.MinInt64
		}
		total += value
	}
	return total
}

// SumExceedsInt64 reports whether a sum of non-negative values exceeds limit.
// Positive overflow is necessarily greater than every representable limit.
func SumExceedsInt64(limit int64, values ...int64) bool {
	var total int64
	for _, value := range values {
		if value > 0 && total > math.MaxInt64-value {
			return true
		}
		total += value
		if total > limit {
			return true
		}
	}
	return false
}
