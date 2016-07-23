package gosort

import "sort"

// SelectionSort uses the selection sort algorith to order the given items and calls the listener
// whenever two items will be swapped.
//
// The complexity for this algorith is N * N/2 = O(N^2)
func SelectionSort(items sort.Interface, listener func(i int, j int)) {
	size := items.Len()

	// The basic of the selection sort algorithm is traversing the items, comparing index to index, from left
	// to right, checking whether an item is less than the other and swapping them when it occurs.
	// The execution is like this:
	//
	// [S O R T M E]
	// [E S R T O M]
	// [E M S T R O]
	// [E M O T S R]
	// [E M O R T S]
	// [E M O R S T]
	// [E M O R S T]
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if items.Less(j, i) {
				items.Swap(i, j)

				if listener != nil {
					listener(i, j)
				}
			}
		}
	}
}
