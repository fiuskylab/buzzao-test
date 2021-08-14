package data

// Data store
type Data struct {
	Batches [][]int
	Sum     int
}

// NewData construct a Data
// with separated batches
func NewData(threads int, nums []int) (data *Data) {
	numPerBatch := len(nums) / threads

	for i := 0; i < threads; i++ {
		batch := nums[i*numPerBatch : (i+1)*numPerBatch]
		data.Batches = append(data.Batches, batch)
	}

	return
}
