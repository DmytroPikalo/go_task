package helpers

import "sync"

type MemoryData struct {
	sync.RWMutex
	URL     string `yaml:"url"`
	Ips     map[string]bool
	Changes map[string][]string
}

type JsonReader struct {
	IpAddress string   `json:"ip_address"`
	Add       []string `json:"add"`
	Remove    []string `json:"remove"`
}
