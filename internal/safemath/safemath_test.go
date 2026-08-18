package safemath

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMulInt64SaturatesAtRepresentationBounds(t *testing.T) {
	const unit = int64(time.Minute)
	maxExact := int64(math.MaxInt64 / unit)
	minExact := int64(math.MinInt64 / unit)

	assert.Equal(t, maxExact*unit, MulInt64(maxExact, unit))
	assert.Equal(t, minExact*unit, MulInt64(minExact, unit))
	assert.Equal(t, int64(math.MaxInt64), MulInt64(maxExact+1, unit))
	assert.Equal(t, int64(math.MinInt64), MulInt64(minExact-1, unit))

	assert.Equal(t, MulInt64(int64(math.MaxInt), unit), MulInt(math.MaxInt, unit))
	assert.Equal(t, MulInt64(int64(math.MinInt), unit), MulInt(math.MinInt, unit))
}

func TestMulFloat64SaturatesAndPreservesNormalFractions(t *testing.T) {
	assert.Equal(t, int64(10), MulFloat64(10.5, 1))
	assert.Equal(t, int64(0), MulFloat64(math.NaN(), 1024))
	assert.Equal(t, int64(math.MaxInt64), MulFloat64(float64(math.MaxInt64), 2))
	assert.Equal(t, int64(math.MinInt64), MulFloat64(float64(math.MinInt64), 2))
}

func TestFloat64ToIntSaturatesBeforeNarrowing(t *testing.T) {
	minInclusive := float64(math.MinInt)
	maxExclusive := -minInclusive
	largestInRange := math.Nextafter(maxExclusive, math.Inf(-1))
	belowMin := math.Nextafter(minInclusive, math.Inf(-1))

	assert.Equal(t, 42, Float64ToInt(42.9))
	assert.Equal(t, -42, Float64ToInt(-42.9))
	assert.Equal(t, 0, Float64ToInt(math.NaN()))
	assert.Equal(t, math.MaxInt, Float64ToInt(maxExclusive))
	assert.Equal(t, int(largestInRange), Float64ToInt(largestInRange))
	assert.Equal(t, math.MinInt, Float64ToInt(minInclusive))
	assert.Equal(t, math.MinInt, Float64ToInt(belowMin))
	assert.Equal(t, math.MaxInt, Float64ToInt(math.Inf(1)))
	assert.Equal(t, math.MinInt, Float64ToInt(math.Inf(-1)))
}

func TestNegateInt64SaturatesMinInt64(t *testing.T) {
	assert.Equal(t, int64(-42), NegateInt64(42))
	assert.Equal(t, int64(42), NegateInt64(-42))
	assert.Equal(t, int64(math.MaxInt64), NegateInt64(math.MinInt64))
}

func TestAddInt64SaturatesInsteadOfChangingSign(t *testing.T) {
	assert.Equal(t, int64(6), AddInt64(1, 2, 3))
	assert.Equal(t, int64(math.MaxInt64), AddInt64(math.MaxInt64, 1))
	assert.Equal(t, int64(math.MinInt64), AddInt64(math.MinInt64, -1))
}

func TestSumExceedsInt64TreatsPositiveOverflowAsExceeded(t *testing.T) {
	assert.False(t, SumExceedsInt64(6, 1, 2, 3))
	assert.True(t, SumExceedsInt64(5, 1, 2, 3))
	assert.True(t, SumExceedsInt64(math.MaxInt64, math.MaxInt64, 1))
}
