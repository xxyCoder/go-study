package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
	var p2 person
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p2)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
