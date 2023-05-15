package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I'm %s. I'm %s years.", name, age)
	//w.Write([]byte("<h1>Hello Full Cycle</h1>"))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/myfamily/family.txt")
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	fmt.Fprintf(w, "My Family: %s.", string(data))
}
