package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	// "github.com/tokopedia/grace"
	gcfg "gopkg.in/gcfg.v1"
	// "github.com/Somesh/pin-config/common/constant"
)

func init() {
	CF = &Config{}
	// GOPATH := os.Getenv("GOPATH")
	ok := ReadConfig(CF, "", "timeConfig")
	// ok := ReadConfig(CF, "/etc", "pin-config") || ReadConfig(CF, GOPATH+"/src/github.com/Somesh/pin-config/files/etc", "pin-config") || ReadConfig(CF, "files/etc", "pin-config")
	if !ok {
		log.Fatal("Failed to read config file")
	}
	makeCountryMap(CF)
	fmt.Printf("Mayank final config>>>%+v\n", CF)
}

func makeCountryMap(cfg *Config) {
	cfg.Countries = make(map[string]map[string]*StateCfg)

	for k, v := range cfg.States {
		log.Println(k, v)
		values := strings.Split(k, "_")
		country := values[0]
		state := values[1]
		pin := values[2]
		state_pin := fmt.Sprintf("%s_%s", state, pin)
		if _, okay := cfg.Countries[country]; !okay {
			cfg.Countries[country] = make(map[string]*StateCfg)
		}
		cfg.Countries[country][state_pin] = v
	}
}

// ReadConfig is file handler for reading configuration files into variable
// Param: - config pointer of Config, filepath string
// Return: - boolean
func ReadConfig(cfg *Config, path string, module string) bool {
	environ := os.Getenv("ENV")
	if environ == "" {
		environ = "DEV" //constant.ENV_DEVELOPMENT
	}

	environ = strings.ToLower(environ)

	parts := []string{"main"}
	var configString []string
	for _, v := range parts {
		fname := path + "/" + module + "/" + environ + "/" + module + "." + v + ".ini"
		fname = "config/timeConfig.main.ini"
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "Reading", fname)
		// config, err := ioutil.ReadFile(fname)
		config, err := ioutil.ReadFile("config/timeConfig.main.ini")
		if err != nil {
			log.Println("ERR 1: function ReadConfig", err)
			return false
		}
		configString = append(configString, string(config))
	}

	err := gcfg.ReadStringInto(cfg, strings.Join(configString, "\n\n"))
	if err != nil {
		log.Println("Err 2: function ReadConfig", err)
		return false
	}
	return true
}

func GetConfig() *Config {
	return CF
}

// func (g GraceCfg) ToGraceConfig() grace.Config {
// 	timeout, err := time.ParseDuration(g.Timeout)
// 	if err != nil {
// 		timeout = time.Second * 5
// 	}

// 	readTimeout, err := time.ParseDuration(g.HTTPReadTimeout)
// 	if err != nil {
// 		timeout = time.Second * 10
// 	}

// 	writeTimeout, err := time.ParseDuration(g.HTTPWriteTimeout)
// 	if err != nil {
// 		timeout = time.Second * 10
// 	}

// 	return grace.Config{
// 		Timeout:          timeout,
// 		HTTPReadTimeout:  readTimeout,
// 		HTTPWriteTimeout: writeTimeout,
// 	}
// }

func GetLogger() *log.Logger {
	var slog *log.Logger

	slog = log.New(ioutil.Discard, "", 0)
	return slog
}
