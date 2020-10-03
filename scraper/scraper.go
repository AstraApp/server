package scraper

import (
	"encoding/json"
	"net/http"
)

// JSON to GO: https://mholt.github.io/json-to-go/
// LIST: https://spdf.gsfc.nasa.gov/data_orbits.html
// EXAMPLE: https://hpde.io/SMWG/Observatory/GOES.json

// Scraper .
type Scraper struct {
}

// New .
func New() *Scraper {
	return &Scraper{}
}

// Scrape .
func (*Scraper) Scrape(satelliteName string) (*Satellite, error) {
	url := "https://hpde.io/SMWG/Observatory/" + satelliteName + ".json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var satellite Satellite
	if err := json.NewDecoder(resp.Body).Decode(&satellite); err != nil {
		return nil, err
	}

	return &satellite, nil
}
