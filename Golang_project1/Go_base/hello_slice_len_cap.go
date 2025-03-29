package Go_base

import "fmt"

// DemonstrateSliceLenCap shows examples of slice length and capacity
func DemonstrateSliceLenCap() []string {
	var results []string

	// Basic slice example
	s := make([]int, 5, 10)
	msg := fmt.Sprintf("Basic Slice:\nLength: %d, Capacity: %d\nElements: %v", len(s), cap(s), s)
	fmt.Println(msg)
	results = append(results, msg)

	// Reslicing example
	s2 := s[:3]
	msg = fmt.Sprintf("\nResliced Slice:\nLength: %d, Capacity: %d\nElements: %v", len(s2), cap(s2), s2)
	fmt.Println(msg)
	results = append(results, msg)

	// Capacity expansion example
	s = append(s, 1, 2, 3, 4, 5, 6)
	msg = fmt.Sprintf("\nAfter Append:\nLength: %d, Capacity: %d\nElements: %v", len(s), cap(s), s)
	fmt.Println(msg)
	results = append(results, msg)

	// Edge cases
	var nilSlice []int
	msg = fmt.Sprintf("\nNil Slice:\nLength: %d, Capacity: %d", len(nilSlice), cap(nilSlice))
	fmt.Println(msg)
	results = append(results, msg)

	zeroSlice := make([]int, 0)
	msg = fmt.Sprintf("\nZero-Length Slice:\nLength: %d, Capacity: %d", len(zeroSlice), cap(zeroSlice))
	fmt.Println(msg)
	results = append(results, msg)

	return results
}
