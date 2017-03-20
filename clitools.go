package main

import (
	"bufio"
	"fmt"
	ad "github.com/zagzagal/Atelier/Data"
	"strings"
)

const title = "Atelier Complete"
const prompt = "$"

type Menu struct {
	items   []MenuItem
	data    *ad.AtelierData
	scanner *bufio.Scanner
}

type MenuItem struct {
	alias  []string
	desc   string
	action func(*ad.AtelierData, *bufio.Scanner)
}

func NewMenu(data *ad.AtelierData, scanner *bufio.Scanner) (m Menu) {
	m.data = data
	m.scanner = scanner
	return
}

func (m *Menu) Add(mi MenuItem) {
	m.items = append(m.items, mi)
}

func (m *Menu) Parse(s string) {
	a := strings.Split(s, " ")
	if a != nil {
		for k, v := range m.items {
			for _, v2 := range v.alias {
				if strings.ToLower(v2) == strings.ToLower(a[0]) {
					m.items[k].action(m.data, m.scanner)
					return
				}
			}
		}
	}
	fmt.Printf("%s is not a command\n", a[0])
}

func (m *Menu) Print() {
	fmt.Printf("%s\n", title)
	for _, v := range m.items {
		fmt.Printf("%s\t%s\n", v.alias[0], v.desc)
	}
	fmt.Printf("%s", prompt)
}
