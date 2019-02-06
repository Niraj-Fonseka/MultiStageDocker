package main

import (
	"fmt"
	"time"
)

func main() {

	location, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Println(err)
	}

	t := time.Now().In(location)

	fmt.Println("Time in Berlin:", t.Format("02.01.2006 15:04"))
}
