package config

import (
// "time"
)

var CF *Config

type Config struct {
  Countries map[string]map[string]*StateCfg
  States    map[string]*StateCfg
}

type StateCfg struct {
  DST      int // daylightsaving time
  Timezone int
}

/*
{
    "au": {
      "sydney_1234": {
        "dst": 2,
        "timezone": 10
      },
      "sydney_5678": {
        "dst": 2,
        "timezone": 10
      }
    }
}
*/
