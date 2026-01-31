package main

import (
	"flag"

	_ "github.com/catcursor/trojan-go/component"
	"github.com/catcursor/trojan-go/log"
	"github.com/catcursor/trojan-go/option"
)

func main() {
	flag.Parse()
	for {
		h, err := option.PopOptionHandler()
		if err != nil {
			log.Fatal("invalid options")
		}
		err = h.Handle()
		if err == nil {
			break
		}
	}
}
