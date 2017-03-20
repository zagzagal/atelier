package Data

// Models and item from the atelier series
type Item struct {
	Name        string   `json`
	Ingredients []string `json`
	Types       []string `json`
}

// returns the item data as a string
func (i *Item) ToString() string {
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
