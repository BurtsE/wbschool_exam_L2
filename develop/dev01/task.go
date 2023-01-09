package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}
	time := time.Now().Add(response.ClockOffset)
	fmt.Println(time.Clock())
	fmt.Println(time)
}
