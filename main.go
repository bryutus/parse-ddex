package main

import (
	"io/ioutil"

	"github.com/bryutus/parse-ddex/unmarshal"
)

func main() {
	bytes, err := ioutil.ReadFile("resources/721620118165_combined.xml")
	if err != nil {
		panic(err)
	}

	unmarshal.Exec(bytes)
}
