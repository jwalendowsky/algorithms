package gosort

// IntsSort is the type used to make ints sort possible using sort.Interface
type IntsSort []int

func (items IntsSort) Len() int           { return len(items) }
func (items IntsSort) Swap(i, j int)      { items[i], items[j] = items[j], items[i] }
func (items IntsSort) Less(i, j int) bool { return items[i] < items[j] }
