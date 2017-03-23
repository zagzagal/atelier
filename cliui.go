package main

import (
	"bufio"
	"fmt"
	ad "github.com/zagzagal/Atelier/Data"
	"os"
	"strings"
)

func startCLIUI(dataFileName string) {
	var data *ad.AtelierData
	if dataFileName == "" {
		data = ad.NewAtelier()
	} else {
		data = loadData(dataFileName)
	}
	bio := bufio.NewScanner(os.Stdin)
	for {
		printMenu()
		bio.Scan()
		switch strings.ToUpper(bio.Text()) {
		case "A":
			addItem(data, bio)
		case "P":
			pathFind(data, bio)
		case "L":
			listItems(data)
		case "O":
			printDot(data)
		case "F":
			data, dataFileName = getDataFromFile(bio, dataFileName)
		case "W":
			dataFileName = writeDataToFile(data, dataFileName, bio)
		case "U":
			getUsed(data, bio)
		case "I":
			itemData(data, bio)
		case "X":
			return
		default:
			fmt.Println("\ncommand not recognized")
		}
	}
}

func printMenu() {
	fmt.Println("A - Add New Item")
	fmt.Println("P - Find Path between Items")
	fmt.Println("L - List Items")
	fmt.Println("O - Output .dot file")
	fmt.Println("F - Load Data from file")
	fmt.Println("W - Write Data to file")
	fmt.Println("U - Used in")
	fmt.Println("I - Item Info")
	fmt.Println("X - Exit")
	fmt.Printf("> ")
}

func addItem(data *ad.AtelierData, bio *bufio.Scanner) {
	var i ad.Item
	item := getInput("Enter the Item Name: ", bio)
	if item == "" {
		fmt.Println("yopu entered nothing")
		return
	}
	i.Name = strings.Title(item)
	fmt.Printf("Enter Item Types: ")
	loop := true
	for loop {
		bio.Scan()
		s := bio.Text()
		if s == "" {
			loop = false
			break
		}
		i.Types = append(i.Types, strings.Title(s))
	}

	fmt.Printf("Enter Item components: ")
	loop = true
	for loop {
		bio.Scan()
		s := bio.Text()
		if s == "" {
			loop = false
			break
		}
		i.Ingredients = append(i.Ingredients, strings.Title(s))
	}
	data.AddItem(i)
}

func printDot(d *ad.AtelierData) {
	fmt.Printf("%s\n", d.PrintDot())
}

func listItems(d *ad.AtelierData) {
	items := d.Items()
	for _, v := range items {
		fmt.Printf("%s\n", v)
	}
}

func pathFind(d *ad.AtelierData, bio *bufio.Scanner) {
	startItem := getInput("Enter the starting Item: ", bio)
	endItem := getInput("Enter the destination Item: ", bio)
	path := d.GetPath(strings.Title(startItem), strings.Title(endItem))
	fmt.Printf("%s\n", path.ToString())
}

func getDataFromFile(bio *bufio.Scanner, def string) (*ad.AtelierData, string) {
	i := fmt.Sprintf("Enter data file name: [%s]", def)
	fileName := getInput(i, bio)
	if fileName == "" {
		fileName = def
	}
	return loadData(fileName), fileName
}

func writeDataToFile(data *ad.AtelierData, def string, bio *bufio.Scanner) string {
	i := fmt.Sprintf("Enter data file name: [%s]", def)
	fileName := getInput(i, bio)
	if fileName == "" {
		fileName = def
	}
	writeData(fileName, data)
	return fileName
}

func getUsed(d *ad.AtelierData, bio *bufio.Scanner) {
	fmt.Printf("Pending implementation\n")
	//i := getInput("Enter Item to examine: ", bio)
	//if i != "" {
	//	fmt.Printf("%s\n", d.usedIn(i))
	//}
}

func itemData(d *ad.AtelierData, bio *bufio.Scanner) {
	i := getInput("Enter Item to examine: ", bio)
	if i != "" {
		item := d.GetItemData(strings.Title(i))
		fmt.Printf("%s\n", item.ToString())
	}
}

func getInput(prompt string, bio *bufio.Scanner) string {
	fmt.Printf(prompt)
	bio.Scan()
	i := bio.Text()
	return i
}
