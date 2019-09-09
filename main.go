package main

import (
	"fmt"
	_ "golang/config"
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

	fmt.Println("in main")
}
