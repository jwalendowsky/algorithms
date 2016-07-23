package gosort

import "sort"

// InsertionSort uses the insertion sort algorith to order the given items and calls the listener
// whenever two items will be swapped.
//
func InsertionSort(items sort.Interface, listener func(i int, j int)) {
	size := items.Len()

	// The idea of this algorithm is just like adding cards to a poker deck. Once we insert a card into
	// a position, we need to shift all other cards to the right. The good things is that we consider all
	// cards from the left already ordered and therefore, if we are sorting an ordered array, we'll have
	// O(n).
	//
	// [S O R T M E]
	// [O S R T M E]
	// [O R S T M E]
	// [O R S T M E]
	// [M O R S T E]
	// [E M O R S T]

	for i := 1; i < size; i++ {
		for j := i; j > 0 && items.Less(j, j-1); j-- {
			if listener != nil {
				listener(j-1, j)
			}
			items.Swap(j, j-1)
		}
	}
}
