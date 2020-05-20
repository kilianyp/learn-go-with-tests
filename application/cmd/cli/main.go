package main

import (
	"fmt"
	"github.com/kilsenp/application"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {name} wins to record a win")
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	defer close()
	if err != nil {
		log.Fatalf("Could not create file system, %v", err)
	}

	fn := poker.BlindAlerterFunc(poker.StdOutAlerter)
	game := poker.NewTexasHoldem(fn, store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()

}
