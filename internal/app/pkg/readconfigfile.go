package pkg

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/spinel/go-musthave-shortener/internal/app/config"
)

func ReadJsonFile(f string) (config.Config, error) {
	// Open jsonFile
	jsonFile, err := os.Open(f)
	if err != nil {
		return config.Config{}, err
	}

	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var c config.Config
	json.Unmarshal(byteValue, &c)

	return c, nil
}
