package main

import (
	"fmt"
	"log"

	"github.com/amaan287/realtimedb/hopper"
)

func main() {

	user := map[string]string{
		"name": "Amaan Mirza",
		"age":  "24",
	}
	db, err := hopper.New()
	if err != nil {
		log.Fatal(err)
	}
	// coll, err := db.CreateCollection("users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	id, err := db.Insert("users", user)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", coll)
	fmt.Printf("%+v\n", id)
}
