package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)

	//new handlers
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// new functions
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	//convert price to string
	f, _ := strconv.ParseInt(price, 0, 64)

	//checks to see if an item already exists
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	//else, creates item key and adds to map db
	db[item] = dollars(f)

	//print status
	fmt.Fprintf(w, "item created: %q with price %v\n", item, f)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	//convert price to string
	f, _ := strconv.ParseInt(price, 0, 64)

	//checks to see if an item is not present in the map
	if _, ok := db[item]; !ok { //if ok is false, the item isnt present
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item does not exist: %q\n", item)
		return
	}

	//else, updates item value
	db[item] = dollars(f)

	//print status
	fmt.Fprintf(w, "item updated: %q to price %v\n", item, f)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	//checks to see if an item exists
	if _, ok := db[item]; !ok { //if ok is false, the item isnt present
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item does not exist: %q\n", item)
		return
	}

	//else, deletes item key and value
	delete(db, item)

	//print status
	fmt.Fprintf(w, "item deleted: %q\n", item)
}
