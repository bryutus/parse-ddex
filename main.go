package main

import (
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("resources/721620118165_combined.xml")
	if err != nil {
		panic(err)
	}

	Unmarshal(bytes)
}
