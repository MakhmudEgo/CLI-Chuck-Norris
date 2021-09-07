package main

import (
	"CLI-Chuck-Norris/cmd/kek/handler"
	"CLI-Chuck-Norris/pkg/parse"
	"fmt"
	"log"
	"os"
)

func main() {
	d, err := parse.Scan(os.Args[1:])
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	switch d.Type {
	case parse.Random:
		err := handler.Random()
		if err != nil {
			log.Fatal(err)
		}
	case parse.Dump:
		err := handler.Dump(d.Count)
		if err != nil {
			log.Fatal(err)
		}
	}
}
