package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	gcfg "gopkg.in/gcfg.v1"
)

func init() {
	CF = &Config{}
	ok := ReadConfig(CF, "", "timeConfigs")
	if !ok {
		log.Fatal("Failed to read config file")
	}
	makeCountryMap(CF)
	// fmt.Printf("Mayank final config>>>%+v\n", CF)
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

func GetLogger() *log.Logger {
	var slog *log.Logger

	slog = log.New(ioutil.Discard, "", 0)
	return slog
}
