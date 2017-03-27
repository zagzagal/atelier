package main

import (
	"encoding/gob"
	ad "github.com/zagzagal/Atelier/Data"
	"io"
	"log"
	"os"
)

func loadData(fName string) *ad.AtelierData {
	file, err := os.Open(fName)
	log.Printf("Open File: %s", fName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := ad.NewAtelier()

	dec := gob.NewDecoder(file)
	var i ad.Item
	for err = dec.Decode(&i); err != io.EOF; err = dec.Decode(&i) {
		log.Printf("Json: %v", i)
		data.AddItem(i.Copy())
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		log.Printf("data: %v", data.GetRawItemData(i.Name))
	}
	log.Printf("%v", data)
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
