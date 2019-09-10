package main

import (
	"fmt"
	config "go-config-read/config"
)

func main() {

	fmt.Println("in main first")
	x := *config.CF
	fmt.Printf("in main %+v \n", x)
	fmt.Println("in main", x.Countries["au"]["sydney_1234"].DST)
	fmt.Println("in main last")
}
