package invoice

import (
	"encoding/json"
	"os"
)

type Project struct {
	Title       string  `json:"title"`
	Cost        float64 `json:"cost"`
	Deliverable string  `json:"deliverable"`
	Client      Client  `json:"client"`
}

type Client struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	NIP     string `json:"nip"`
	REGON   string `json:"regon"`
}

func LoadProject(filename string) *Project {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var project Project
	if err := json.Unmarshal(data, &project); err != nil {
		panic(err)
	}

	return &project
}
