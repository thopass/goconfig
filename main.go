package goconfig

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Configuration map[string]map[string]string

func New() Configuration {
	result := make(map[string]map[string]string)
	return Configuration(result)
}

func (c *Configuration) AddSection(name string) {
	if _, status := (*c)[name]; status != true {
		(*c)[name] = make(map[string]string)
	}
}

func (c *Configuration) AddValue(section, key, value string) error {
	if _, status := (*c)[section]; status != true {
		return errors.New("No such section: " + section)
	}
	(*c)[section][key] = value
	return nil
}

func (c *Configuration) GetValue(section, key string) (string, error) {
	if _, status := (*c)[section]; status != true {
		return "", errors.New("No such section: " + section)
	}
	if _, status := (*c)[section][key]; status != true {
		return "", errors.New("No such key " + key + " for section " + section)
	}
	return (*c)[section][key], nil
}

func (c *Configuration) ReadFromIni(filePath string) error {
	input, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	section := "UNNAMED"
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		switch {
		case len(line) == 0:
			// empty line - drop it
			continue
		case line[0] == '[' && line[len(line)-1] == ']':
			// this is new section - multisection files not supported yet
			section = strings.Trim(line, " []\t")
			continue
		case line[0] == ';':
			// this is comment - skip it
			continue
		case strings.Count(line, "=") == 1:
			// this seems to be proper option
			tokens := strings.Split(line, "=")
			if _, status := (*c)[section]; status != true {
				(*c)[section] = make(map[string]string)
			}
			c.AddValue(section, strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1]))
		default:
			return errors.New("Unsupported line format:" + line)
		}
	}

	return nil
}

func (c *Configuration) WriteToIni(filePath string) error {
	output, err := os.Create(filePath)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(output)

	for section, config := range *c {
		writer.WriteString("; start new section")
		writer.WriteString("[" + section + "]")
		for key, value := range config {
			writer.WriteString("\t" + key + " = " + value)
		}
		// print empty line just to make some separation
		writer.WriteString("\n")
	}
	return nil
}

func (c *Configuration) ReadFromYaml(filePath string) error {
	return errors.New("Not implemented")
}

func (c *Configuration) WriteToYaml(filePath string) error {
	return errors.New("Not implemented")
}
