package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	// static route
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// webapp route
	http.Handle("/app", http.StripPrefix("/app", http.FileServer(http.Dir("./app"))))

	// websocket

	// listen
	fmt.Println("Listening.. on PORT:", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
