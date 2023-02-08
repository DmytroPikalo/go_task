package helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func (m *MemoryData) IpExists(ip string) bool {
	_, ok := m.Ips[ip]
	return ok
}

func (m *MemoryData) Match(ip string) []string {
	var match []string

	for k := range m.Ips {
		if strings.Contains(k, ip) {
			match = append(match, k)
		}
	}
	return match
}

func (m *MemoryData) ChangeApis(data JsonReader) {
	changes := make(map[string][]string)
	m.RLock()
	defer m.RUnlock()

	if data.Add != nil {
		for _, ip := range data.Add {
			if ValidateIp(ip) {
				changes["added"] = append(changes["added"], ip)
				m.Ips[ip] = true
			}
		}
	}

	if data.Remove != nil {
		for _, ip := range data.Remove {
			if _, ok := m.Ips[ip]; ok {
				changes["removed"] = append(changes["removed"], ip)
				delete(m.Ips, ip)
			}
		}
	}

	m.Changes = changes

}

func MakeJsonRes(w http.ResponseWriter, value interface{}) {
	jData, err := json.Marshal(value)
	if err != nil {
		log.Panicln("MakeJsonRes. Error: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func ValidateIp(ip string) bool {
	re := regexp.MustCompile(`^(([1-9]?\d|1\d\d|25[0-5]|2[0-4]\d)\.){3}([1-9]?\d|1\d\d|25[0-5]|2[0-4]\d)$`)
	return re.MatchString(ip)
}

func ValidateStr(ip string) bool {
	re := regexp.MustCompile(`^[\.0-9]*$`)
	return re.MatchString(ip)
}
