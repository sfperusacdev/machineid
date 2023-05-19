package main

import (
	"fmt"
	"log"

	"github.com/panta/machineid"
)

func main() {
	id, err := machineid.ID()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
}
