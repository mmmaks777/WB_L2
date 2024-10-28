package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {
	execTime, err := ntp.Time("time.google.com")
	if err != nil {
		log.Fatalln("Error og getting time: ", err)
	}
	fmt.Println("Exact time: ", execTime)
}
