// stats.go
package stats

import (
	"cmp"
	"errors"
	"math"
	"slices"
)

// Integer represents all signed and unsigned integer types.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float represents all floating-point types.
type Float interface {
	~float32 | ~float64
}

// IntegerOrFloat represents types that are either integers or floats.
type IntegerOrFloat interface {
	Integer | Float
}

// Average returns the arithmetic mean of a slice of integers or floats.
// It returns an error if the slice is empty.
func Average[T IntegerOrFloat](slice []T) (float64, error) {
	if len(slice) == 0 {
		return 0, errors.New("cannot calculate average of an empty slice")
	}
	var sum float64
	for _, v := range slice {
		sum += float64(v)
	}
	return sum / float64(len(slice)), nil
}

// Median returns the median of a slice of integers or floats.
// It clones and sorts the slice internally to avoid modifying the input.
// It returns an error if the slice is empty.
func Median[T IntegerOrFloat](slice []T) (float64, error) {
	if len(slice) == 0 {
		return 0, errors.New("cannot calculate median of an empty slice")
	}
	cloned := slices.Clone(slice)
	slices.Sort(cloned)

	n := len(cloned)
	if n%2 == 1 {
		return float64(cloned[n/2]), nil
	}
	return (float64(cloned[n/2-1]) + float64(cloned[n/2])) / 2.0, nil
}

// Mode returns the value that appears most frequently in the slice.
// If multiple values have the same maximum frequency, it returns the one that appeared first.
// It returns an error if the slice is empty.
func Mode[T cmp.Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("cannot calculate mode of an empty slice")
	}
	counts := make(map[T]int)
	for _, v := range slice {
		counts[v]++
	}

	var mode T
	maxCount := 0
	for _, v := range slice {
		if counts[v] > maxCount {
			maxCount = counts[v]
			mode = v
		}
	}
	return mode, nil
}

// SumSlice returns the sum of a slice of integers or floats.
// It returns an error if the slice is empty, or if the sum overflows the bounds of the type T (for integer types).
func SumSlice[T IntegerOrFloat](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("cannot calculate sum of an empty slice")
	}

	var zero T
	switch any(zero).(type) {
	case float32:
		var sum float32
		for _, v := range slice {
			sum += any(v).(float32)
		}
		return any(sum).(T), nil
	case float64:
		var sum float64
		for _, v := range slice {
			sum += any(v).(float64)
		}
		return any(sum).(T), nil
	case int:
		var sum int
		for _, v := range slice {
			val := any(v).(int)
			if val > 0 && sum > math.MaxInt-val {
				return zero, errors.New("sum overflows type")
			}
			if val < 0 && sum < math.MinInt-val {
				return zero, errors.New("sum underflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case int8:
		var sum int8
		for _, v := range slice {
			val := any(v).(int8)
			if val > 0 && sum > math.MaxInt8-val {
				return zero, errors.New("sum overflows type")
			}
			if val < 0 && sum < math.MinInt8-val {
				return zero, errors.New("sum underflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case int16:
		var sum int16
		for _, v := range slice {
			val := any(v).(int16)
			if val > 0 && sum > math.MaxInt16-val {
				return zero, errors.New("sum overflows type")
			}
			if val < 0 && sum < math.MinInt16-val {
				return zero, errors.New("sum underflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case int32:
		var sum int32
		for _, v := range slice {
			val := any(v).(int32)
			if val > 0 && sum > math.MaxInt32-val {
				return zero, errors.New("sum overflows type")
			}
			if val < 0 && sum < math.MinInt32-val {
				return zero, errors.New("sum underflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case int64:
		var sum int64
		for _, v := range slice {
			val := any(v).(int64)
			if val > 0 && sum > math.MaxInt64-val {
				return zero, errors.New("sum overflows type")
			}
			if val < 0 && sum < math.MinInt64-val {
				return zero, errors.New("sum underflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case uint:
		var sum uint
		for _, v := range slice {
			val := any(v).(uint)
			if sum > math.MaxUint-val {
				return zero, errors.New("sum overflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case uint8:
		var sum uint8
		for _, v := range slice {
			val := any(v).(uint8)
			if sum > math.MaxUint8-val {
				return zero, errors.New("sum overflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case uint16:
		var sum uint16
		for _, v := range slice {
			val := any(v).(uint16)
			if sum > math.MaxUint16-val {
				return zero, errors.New("sum overflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case uint32:
		var sum uint32
		for _, v := range slice {
			val := any(v).(uint32)
			if sum > math.MaxUint32-val {
				return zero, errors.New("sum overflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case uint64:
		var sum uint64
		for _, v := range slice {
			val := any(v).(uint64)
			if sum > math.MaxUint64-val {
				return zero, errors.New("sum overflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	case uintptr:
		var sum uintptr
		for _, v := range slice {
			val := any(v).(uintptr)
			if sum > ^uintptr(0)-val {
				return zero, errors.New("sum overflows type")
			}
			sum += val
		}
		return any(sum).(T), nil
	default:
		return zero, errors.New("unsupported type")
	}
}

// MinSlice returns the minimum value in a slice of ordered values.
// It returns an error if the slice is empty.
func MinSlice[T cmp.Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("cannot calculate minimum of an empty slice")
	}
	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// MaxSlice returns the maximum value in a slice of ordered values.
// It returns an error if the slice is empty.
func MaxSlice[T cmp.Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("cannot calculate maximum of an empty slice")
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}
