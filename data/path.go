package Data

type ItemPath struct {
	Item []string `json`
}

func (i *ItemPath) ToString() string {
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

func (i *ItemPath) ToArray() []string {
	return i.Item
}
