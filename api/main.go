package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ishkanan/tienlen/api/game"
	"github.com/ishkanan/tienlen/api/utils"
)

var addr = flag.String("addr", "localhost:27000", "HTTP service address")
var uiFolder = flag.String("ui", "dist", "Folder container UI files")

func main() {
	fmt.Print("Tiến lên (aka. Thirteen) server\n" +
		"  A simple server implementation of the popular Vietnamese card game.\n\n",
	)

	flag.Parse()
	log.SetFlags(0)

	theGame := game.NewGame()
	http.Handle("/", http.FileServer(http.Dir(*uiFolder)))
	http.HandleFunc("/api", game.ConnectionHandler(theGame))

	utils.LogInfo("Server listening on %s ...", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
