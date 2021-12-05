package main

//go:generate templeGen -dir templates -pkg main -var myTemplates -o templates.go

import (
	"encoding/json"
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

	fruit := Fruit{
		Name:  "Strawberry",
		Color: "red",
	}

	res, err := PrettyStruct(fruit)
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir(config.FolderContents))

	http.Handle("/", fileServer)

	http.HandleFunc(config.FirstUrlRedirect, func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(res))
	})

	http.HandleFunc(config.SecondUrlRedirect, func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(res))
	})

	http.HandleFunc(config.UrlRedirect, redirect)
	http.ListenAndServe(config.HostPort, nil)

}

func redirect(w http.ResponseWriter, r *http.Request) {

	config, err := config.LoadConfig("./properties")
	if err != nil {
		log.Fatal("Cannot load config file: ", err)
	}

	if (counter % 2) == 0 {
		http.Redirect(w, r, config.Protocol+config.Localhost+config.HostPort+config.FirstUrlRedirect, 303)
		counter++

	} else {
		http.Redirect(w, r, config.Protocol+config.Localhost+config.HostPort+config.SecondUrlRedirect, 303)
		counter++
	}

}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

type Fruit struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}
