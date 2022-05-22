package internal

import (
	"flag"
	"fmt"
	"strings"
)

var (
	url        = flag.String("u", "", "Url for request")
	times      = flag.Int("t", -1, "Times to make request")
	concurrent = flag.Int("c", 5, "Concurrent count")
	method     = flag.String("m", "GET", "Request method")
	configPath = flag.String("config", "", "Path to json config")
	status     = flag.Int("s", 200, "Required response status")
	protocol   = flag.String("p", "https", "Protocol")
)

func Bootstrap() *Config {
	flag.Parse()

	if *configPath == "" {
		if !strings.Contains(*url, "http://") || !strings.Contains(*url, "https://") {
			*url = fmt.Sprintf("%s://%s", *protocol, *url)
		}

		return &Config{
			Url:            *url,
			Times:          *times,
			Concurrent:     *concurrent,
			Method:         *method,
			RequiredStatus: *status,
		}
	} else {
		return &Config{}
	}
}
