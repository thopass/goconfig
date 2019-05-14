package goconfig

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Configuration struct {
	Data map[string]string
}

func New() *Configuration {
	result := new(Configuration)
	result.Data = make(map[string]string)

	return result
}

func (c *Configuration) ReadFromIni(filePath string) error {
	input, err := os.Open(filePath)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		switch {
		case len(line) == 0:
			// empty line - drop it
			continue
		case line[0] == '[' && line[len(line)-1] == ']':
			// this is new section - multisection files not supported yet
			continue
		case line[0] == ';':
			// this is comment - skip it
			continue
		case strings.Count(line, "=") == 1:
			// this seems to be proper option
			tokens := strings.Split(line, "=")
			c.Data[strings.TrimSpace(tokens[0])] = strings.TrimSpace(tokens[1])
		default:
			return errors.New("Unsupported line format:" + line)
		}
	}

	return nil
}

func (c *Configuration) WriteToIni(filePath string) error {
	return errors.New("Not implemented")
}

func (c *Configuration) ReadFromYaml(filePath string) error {
	return errors.New("Not implemented")
}

func (c *Configuration) WriteToYaml(filePath string) error {
	return errors.New("Not implemented")
}
