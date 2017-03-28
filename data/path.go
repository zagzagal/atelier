package Data

// ItemPath defines a path struct
type ItemPath struct {
	Item []string `json`
}

func (i *ItemPath) String() string {
	var ans string
	for k, v := range i.Item {
		if k+1 != len(i.Item) {
			ans += v + " -> "
		} else {
			ans += v
		}
	}
	return ans
}

// ToArray returns the path as an Array
func (i *ItemPath) ToArray() []string {
	return i.Item
}
