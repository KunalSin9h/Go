package utils

// Add adds multiple numbers
func Add(nums ...int32) int32 {
	total := int32(0)
	for _, v := range nums {
		total += v
	}
	return total
}
