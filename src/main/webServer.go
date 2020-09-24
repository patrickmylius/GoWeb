package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8082", nil))
}

/*
Dette er en webserver der åbner html/css filer,
 fra en mappe der hedder static

Virker ikke pt, kæmper med at fixe port.
Den giver siden, men 404 fejl server not found
*/
