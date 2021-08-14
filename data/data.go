package data

// Data store
type Data struct {
	Batches [][]int
	Sum     int
}

// NewData construct a Data
// with separated batches
func NewData(threads int, nums []int) *Data {
	numPerBatch := len(nums) / threads

	data := &Data{
		Sum:     0,
		Batches: [][]int{},
	}

	for i := 0; i < threads; i++ {
		batch := nums[i*numPerBatch : (i+1)*numPerBatch]
		data.Batches = append(data.Batches, batch)
	}

	return data
}

func (d *Data) processBatch(pos int) {
	for _, n := range d.Batches[pos] {
		d.Sum += n
	}
}

// EachBatch a
func EachBatch(d *Data) {
	for i := 0; i < len(d.Batches); i++ {
		d.processBatch(i)
	}
}
