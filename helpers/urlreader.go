package helpers

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func (m *MemoryData) GetConf() (*MemoryData, error) {

	yamlFile, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return m, err
	}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return m, err
}

func (m *MemoryData) LinesToMap() error {
	resp, err := http.Get(m.URL)
	if err != nil {
		log.Panicln("Could not get the Url. Error: ", err)
		return err
	}
	defer resp.Body.Close()
	m.RLock()
	defer m.RUnlock()
	m.Ips, err = m.linesFromReader(resp.Body)
	if err != nil {
		log.Panicln("There was a problem when processing a file. Error: ", err)
	}
	log.Println("Now we have ", m.Ips)
	return err
}

func (m *MemoryData) linesFromReader(r io.Reader) (map[string]bool, error) {
	ips := make(map[string]bool)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if text := scanner.Text(); ValidateIp(text) {
			ips[text] = true
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ips, nil
}
