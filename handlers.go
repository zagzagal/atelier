package main

import (
	ad "bitbucket.org/zagzagal/AtelierComplete/AtelierData"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ilItem struct {
	Name string `json`
	Link string `json`
}

type itemlist struct {
	Item []ilItem `json`
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func ApiIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "api todo")
}

func ItemList(w http.ResponseWriter, r *http.Request) {
	var il []ilItem
	for _, v := range DATA.GetAllItems() {
		link := "/api/item/" + v
		il = append(il, ilItem{Name: v, Link: link})
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(il); err != nil {
		panic(err)
	}
}

func ItemShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemName := vars["itemID"]
	itemData := DATA.GetItemData(itemName)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(itemData); err != nil {
		panic(err)
	}
}

func PathShow(w http.ResponseWriter, r *http.Request) {
	type ipath struct {
		Items   []ilItem
		DotFile string
	}

	vars := mux.Vars(r)
	start := vars["start"]
	dest := vars["dest"]
	var ip ipath
	path, err := DATA.GetPath(dest, start)
	if err == nil {
		ip.DotFile = "digraph test {"
		for k, v := range path.Item {
			if k > 0 {
				ip.DotFile += "\"" + path.Item[k-1] + "\"->\"" + v + "\";"
			}
			link := "/api/item/" + v
			ip.Items = append(ip.Items, ilItem{Name: v, Link: link})
		}
		ip.DotFile += "}"
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(ip); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json;charset-UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(ipath{}); err != nil {
			panic(err)
		}
	}
}

func ItemCreate(w http.ResponseWriter, r *http.Request) {
	var i ad.Item
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&i)
	log.Printf("request - %v\n", r)
	log.Printf("unmarshaled - %v\n", i)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable enity
		log.Printf("%v\n", i)
		log.Printf("the error is %s\n", err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	DATA.AddItem(i)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(i); err != nil {
		panic(err)
	}
}
