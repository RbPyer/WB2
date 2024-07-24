package patterns

type Context struct {
	strategy IStrategy
}

func (c *Context) SetStrategy(strategy IStrategy) {
	c.strategy = strategy
}

func (c *Context) SortArray(arr []int) []int {
	return c.strategy.Sort(arr)
}

type IContext interface {
	SetStrategy(strategy IStrategy)
	SortArray(arr []int) []int
}

type IStrategy interface {
	Sort([]int) []int
}

type BubbleSortStrategy struct{}

func (s *BubbleSortStrategy) Sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

type QuickSortStrategy struct{}

func (s *QuickSortStrategy) Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	elem := arr[0]

	left, mid, right := getSlices(arr, elem)

	result := make([]int, 0)
	result = append(result, s.Sort(left)...)
	result = append(result, mid...)
	result = append(result, s.Sort(right)...)

	return result

}

func getSlices(s []int, elem int) ([]int, []int, []int) {
	ltSlice := make([]int, 0)
	gtSlice := make([]int, 0)
	eqSlice := make([]int, 0)

	for _, item := range s {
		if item > elem {
			gtSlice = append(gtSlice, item)
		} else if item < elem {
			ltSlice = append(ltSlice, item)
		} else {
			eqSlice = append(eqSlice, item)
		}
	}
	return ltSlice, eqSlice, gtSlice
}
