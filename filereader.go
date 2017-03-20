package main

import (
	// "bufio"
	"io"
	"log"
	"os"
	//"fmt"
	ad "github.com/zagzagal/Atelier/Data"
	//"regexp"
	"encoding/gob"
)

func loadData(fName string) *ad.AtelierData {
	file, err := os.Open(fName)
	log.Printf("Open File: %s", fName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//var validLine = regexp.MustCompile(`"(.*)" -> "(.*)"`)
	data := ad.NewAtelier()

	/*scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := validLine.FindStringSubmatch(scanner.Text())
		if line != nil {
			data.AddPath(line[1], line[2])
			log.Printf("Adding path %s -> %s", line[1], line[2])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}*/
	dec := gob.NewDecoder(file)
	var i ad.Item
	for err = dec.Decode(&i); err != io.EOF; err = dec.Decode(&i) {
		data.AddItem(i)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	return data
}

func writeData(fName string, data *ad.AtelierData) {
	f, err := os.Create(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	for _, v := range data.Items() {
		err = enc.Encode(data.GetItemData(v))
		if err != nil {
			log.Fatal(err)
		}
	}
	f.Sync()
}
