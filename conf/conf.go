package conf

import (
	"sync"
)

var (
	once sync.Once
	conf *Conf
)

// Conf
type Conf struct {
	TestUrl string
	ProUrl  string
}

// Conf
func GetConf() *Conf {
	once.Do(func() {
		conf = &Conf{
			TestUrl: "https://openapi.akucun.com",
			ProUrl:  "https://openapi.aikucun.com",
		}
	})
	return conf
}
