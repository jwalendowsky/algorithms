package gosort

import (
	"math/rand"
	"sort"
	"time"
)

// QuickSortInts sorts an integer array using the quick sort algorithm. It was created by C.A.R. Hoare.
func QuickSortInts(items []int) {

	// // // The first step of the quick sort is shuffling its entries, used to
	// ShuffleInts(items)

	// // // Now its time to sort the ints.
	// sortInts(items, 0, len(items)-1)
	QuickSort(IntsSort(items))
}

// QuickSortStrings sorts an string array using the quick sort algorithm. It was created by C.A.R. Hoare.
func QuickSortStrings(items []string) {
	QuickSort(StringsSort(items))
}

// QuickSort sorts an array of sort interface using the quick sort algorithm. It was created by C.A.R. Hoare.
func QuickSort(items sort.Interface) {

	// The first step of the quick sort is shuffling its entries, used to
	Shuffle(items)

	// Now its time to sort the ints.
	doSort(items, 0, items.Len()-1)
}

// SHUFFLING FUNCTIONS

// ShuffleInts shuffles an array of ints.
func ShuffleInts(items []int) {
	size := len(items)

	rand.Seed(time.Now().UnixNano()) // Let's update the random seed

	// Let's use the Fisher–Yates modern shuffle algorithm to shuffle this array.
	for i := size - 1; i >= 1; i-- {
		j := rand.Intn(i + 1) // A random integer between 0 and i

		// NOTE: This go construction is amazing for exchanging items without the temp variable.
		items[i], items[j] = items[j], items[i]
	}
}

// Shuffle shuffles an array of items that implement the sort.Interface
func Shuffle(items sort.Interface) {
	size := items.Len()

	rand.Seed(time.Now().UnixNano()) // Let's update the random seed

	// Let's use the Fisher–Yates modern shuffle algorithm to shuffle this array.
	for i := size - 1; i >= 1; i-- {
		j := rand.Intn(i) // A random integer between 0 and i

		items.Swap(i, j)
	}
}

// SORTING FUNCTIONS

func doSort(items sort.Interface, min int, max int) {

	// If the min is equal or greater than the max, our sort endend.
	if min >= max {
		return
	}

	// Let's partition the array in a way that the return is the position of an element that will not
	// change and all elements at his left are smaller than it and all at his right are bigger.
	finalPositionIndex := Partition(items, min, max)

	// Now let's use divide and conquer to sort all elements at the left and at the right of the final
	// position index
	doSort(items, min, finalPositionIndex-1)
	doSort(items, finalPositionIndex+1, max)

}

// PARTITION FUNCTIONS

// Partition creates a QuickSort partition of an array slice returning and index
// (min <= finalPositionIndex <= max) with the following definitions:
// - This index is the final position of the element on the array.
// - All the elements before the index are smaller/equal to the value of the element.
// - All the elements after the index are grater than the value of the element.
func Partition(items sort.Interface, min int, max int) (finalPosition int) {

	// Let's use another algorithm. This one is found on the Princeton course.
	finalPosition = min
	i := min + 1
	j := max

	for true {
		for items.Less(i, min) {
			i++
			if i == max {
				break
			}
		}

		for items.Less(min, j) {
			j--
			if j == min {
				break
			}
		}

		if i >= j {
			break
		}
		items.Swap(i, j)
	}

	items.Swap(min, j)

	finalPosition = j
	return
}

// PartitionOld creates a QuickSort partition of an array slice returning and index
// (min <= finalPositionIndex <= max) with the following definitions:
// - This index is the final position of the element on the array.
// - All the elements before the index are smaller/equal to the value of the element.
// - All the elements after the index are grater than the value of the element.
func PartitionOld(items sort.Interface, min int, max int) (finalPosition int) {
	finalPosition = min  // Let's start assuming that the final position is the minimum.
	rightPosition := max // Let's keep track of the right position of the array

	// Let's walk from both sides of the array until we reach the middle of it.
	for finalPosition < rightPosition {

		// If the value of the element on the final position is greater than the value of the element on
		// finalPosition + 1, let's swap than and increment the final position.
		if items.Less(finalPosition+1, finalPosition) { // items[finalPosition] >= items[finalPosition+1]
			items.Swap(finalPosition, finalPosition+1)
			finalPosition++
		} else if items.Less(rightPosition, finalPosition) { // items[finalPosition] >= items[rightPosition]

			// Otherwise, if the element on the final position is greater than the element on the right position,
			// let's swipe the element on the right position with the element of the finalPosition + 1 (since
			// we know such element is bigger than the element on the final position. After that, let's swap
			// the value of the final position with the element on final position + 1, updating the final position.
			// items[finalPosition+1], items[rightPosition] = items[rightPosition], items[finalPosition+1]
			items.Swap(finalPosition+1, rightPosition)
			items.Swap(finalPosition, finalPosition+1)
			finalPosition++
		} else {

			// Otherwise, let's decrement the right position
			rightPosition--
		}
	}

	return
}
