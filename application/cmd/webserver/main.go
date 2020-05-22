package main

import (
	"github.com/kilsenp/application"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {

	store, close_fn, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	defer close_fn()

	if err != nil {
		log.Fatal(err)
	}

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	} else {
		log.Printf("server started on port 5000")
	}

}
