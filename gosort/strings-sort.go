package gosort 

// StringsSort is the type used to make strings sort possible using sort.Interface
type StringsSort []string

func (items StringsSort) Len() int           { return len(items) }
func (items StringsSort) Swap(i, j int)      { items[i], items[j] = items[j], items[i] }
func (items StringsSort) Less(i, j int) bool { return items[i] < items[j] }