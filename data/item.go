package Data

// Models and item from the atelier series
type Item struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
	Types       []string `json:"types"`
}

func (i Item) Copy() Item {
	var j Item
	j.Name = i.Name
	j.Ingredients = make([]string, len(i.Ingredients))
	for k, v := range i.Ingredients {
		j.Ingredients[k] = v
	}
	j.Types = make([]string, len(i.Types))
	for k, v := range i.Types {
		j.Types[k] = v
	}
	return j
}

func (i Item) equals(b Item) bool {
	if i.Name != b.Name {
		return false
	}
	if !compArray(i.Ingredients, b.Ingredients) {
		return false
	}
	if !compArray(i.Types, b.Types) {
		return false
	}
	return true
}

func compArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, _ := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

// returns the item data as a string
func (i *Item) String() string {
	var ans string
	ans += i.Name
	ans += "\nIs of types: \n"
	for _, v := range i.Types {
		ans += "\t" + v + "\n"
	}
	if len(i.Ingredients) > 0 {
		ans += "Is made from: \n"
		for _, v := range i.Ingredients {
			ans += "\t" + v + "\n"
		}
	}
	return ans
}
