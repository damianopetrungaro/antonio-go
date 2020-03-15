package main

import (
	"errors"
	"fmt"
	"github.com/damianopetrungaro/antonio/customer"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	s := http.NewServeMux()

	s.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Ciao user")); err != nil {
			log.Fatalf("impossible scrvere la risposta: %s", err)
		}
	})

	// localhost/customers-save?name=Damiano&surname=Petrungaro&age=21
	s.HandleFunc("/customers-save", func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		name, err := prendiInput(v, "name")
		if err != nil {
			w.Write([]byte("name:" + err.Error()))
			w.WriteHeader(400)
			return
		}

		surname, err := prendiInput(v, "surname")
		if err != nil {
			w.Write([]byte("surname:" + err.Error()))
			w.WriteHeader(400)
			return
		}

		ageString, err := prendiInput(v, "age")
		if err != nil {
			w.Write([]byte("age:" + err.Error()))
			w.WriteHeader(400)
			return
		}

		age, err := strconv.Atoi(ageString)
		if err != nil {
			w.Write([]byte("age non è un intero"))
			w.WriteHeader(400)
			return
		}

		c := customer.New(name, surname, age)
		fmt.Println(c)
	})

	// localhost/customers-get?name=Damiano
	s.HandleFunc("/customers-get", func(w http.ResponseWriter, r *http.Request) {
	})

	if err := http.ListenAndServe(":81", s); err != nil {
		log.Fatalf("Errore web server non funzionante: %s", err)
	}
}

func prendiInput(values url.Values, key string) (string, error) {
	// map[string][]string
	// "name":["Damiano", "Antonio"]
	// "name":["Damiano"]
	// "name":[]
	inputs, ok := values[key]
	if ok == false {
		return "", errors.New("parametro non inviato")
	}

	if len(inputs) != 1 {
		return "", errors.New("il parametro inviato deve essere uno")
	}

	return inputs[0], nil
}
