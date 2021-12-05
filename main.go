package main

import (
	"log"
	"net/http"

	"github.com/Pierfabio/redirector/config"
)

var counter int

func main() {

	config, err := config.LoadConfig("./properties")
	if err != nil {
		log.Fatal("Cannot load config file: ", err)
	}

	fileServer := http.FileServer(http.Dir(config.FolderContents))

	http.Handle("/", fileServer)

	// http.Handle("/resources/", http.StripPrefix("/resources", fileServer))

	http.HandleFunc(config.UrlRedirect, redirect)
	http.ListenAndServe(config.HostPort, nil)

}

func redirect(w http.ResponseWriter, r *http.Request) {

	config, err := config.LoadConfig("./properties")
	if err != nil {
		log.Fatal("Cannot load config file: ", err)
	}

	if (counter % 2) == 0 {
		http.Redirect(w, r, config.Protocol+config.FirstUrlRedirect, 303)
		counter++

	} else {
		http.Redirect(w, r, config.Protocol+config.SecondUrlRedirect, 303)
		counter++
	}

}
