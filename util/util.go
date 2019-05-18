package util

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func ParseSetting(line string) Setting {
	keyReg := regexp.MustCompile(`^\*color\d{1,2}`)
	valueReg := regexp.MustCompile(`^#[a-zA-Z0-9]+`)
	result := strings.Fields(line)

	key := strings.Trim(keyReg.FindString(result[0]), "*")
	value := valueReg.FindString(result[1])

	return Setting{Key: key, Value: value}
}

func ReadConfig(path string) ([]Setting, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	config := []Setting{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := ParseSetting(scanner.Text())
		config = append(config, s)
	}

	return config, nil
}

type Setting struct {
	Key   string
	Value string
}
