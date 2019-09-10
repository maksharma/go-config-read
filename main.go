package main

import (
	"fmt"
	config "go-config-read/config"
)

type Configuration struct {
	Users  []string
	Groups []string
}

type TimeZoneConfig struct {
	timezoneList []*TimezoneList
}

type TimezoneList struct {
	country map[string]*Country
	// country Country
}

type Country struct {
	state_pincode map[string]StatePincode
	// state_pincode *StatePincode
}

type StatePincode struct {
	timezone int
	dst      int
}

func main() {

	fmt.Println("in main first")
	x := *config.CF
	fmt.Printf("in main %+v \n", x)
	fmt.Println("in main", x.Countries["au"]["sydney_1234"].DST)
	fmt.Println("in main last")
}
