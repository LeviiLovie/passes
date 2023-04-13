package file

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppVersion float64
	DBVersion  float64
}

func readConfigFile(filePath string) (*Config, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := &Config{}

	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines and comments
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Invalid config line: %q", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "app.version":
			if appVersion, err := strconv.ParseFloat(value, 64); err == nil {
				config.AppVersion = appVersion
			} else {
				log.Printf("Invalid app.version value: %q", value)
			}
		case "db.version":
			if dbVersion, err := strconv.ParseFloat(value, 64); err == nil {
				config.DBVersion = dbVersion
			} else {
				log.Printf("Invalid db.version value: %q", value)
			}
		default:
			log.Printf("Unknown config key: %q", key)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}
