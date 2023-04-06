package internal

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/tidwall/sjson"
)

var _ TemplateReader = (*FileTemplateReader)(nil)

type TemplateReader interface {
	Read(name string) (string, error)
}

type FileTemplateReader struct{}

func (f FileTemplateReader) Read(path string) (string, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func ReadTemplateAndSetValues(reader TemplateReader, name string, sets []string) (string, error) {
	value, err := reader.Read(name)
	if err != nil {
		return "", err
	}

	return AddSetValues(value, sets)
}

func AddSetValues(jsonString string, sets []string) (string, error) {
	for _, set := range sets {
		var value bool
		var err error

		keyValue := strings.Split(set, "=")
		if keyValue[1] == "true" || keyValue[1] == "false" {
			value, err = strconv.ParseBool(keyValue[1])
			if err != nil {
				return "", err
			}

			jsonString, err = sjson.Set(jsonString, keyValue[0], value)
		} else {
			jsonString, err = sjson.Set(jsonString, keyValue[0], keyValue[1])
		}

		if err != nil {
			return "", err
		}
	}

	return jsonString, nil
}

func HandleSets(object interface{}, sets []string) error {
	if reflect.ValueOf(object).Kind() != reflect.Ptr {
		return errors.New("out put must be a pointer")
	}

	bytes, err := json.Marshal(object)
	if err != nil {
		return err
	}

	values, err := AddSetValues(string(bytes), sets)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(values), object)
}
