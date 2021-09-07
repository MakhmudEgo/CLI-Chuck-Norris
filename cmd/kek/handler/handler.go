package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Joke struct {
	Value string `json:"value"`
}

func Random() error {
	kek := new(Joke)
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if err := resp.Body.Close(); err != nil {
		return fmt.Errorf(err.Error())
	}
	if err := json.Unmarshal(b, kek); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(kek.Value)
	return nil
}
