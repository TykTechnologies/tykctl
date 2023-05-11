package internal

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"

	"github.com/TykTechnologies/tykctl/util"
)

func CreateViper(dir, file string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(file)
	v.AddConfigPath(dir)
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	err := v.ReadInConfig()

	return v, err
}

func CreateCoreViper() (*viper.Viper, error) {
	dir, err := GetCoreDir()
	if err != nil {
		return nil, err
	}

	return CreateViper(dir, CoreConfig)
}

func GetCoreDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, DefaultConfigDir)

	return configDir, nil
}

func GetDefaultConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, DefaultConfigDir, Config)

	return configDir, nil
}

func CreateConfigFile(dir, fileName string) error {
	err := util.CheckDirectory(dir)
	if err != nil {
		return err
	}

	fullFileName := fmt.Sprintf("config_%s.yaml", fileName)

	return CreateFile(dir, fullFileName)
}

func GetAllConfig(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	infos := make([]string, 0, len(entries))

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		if isConfigFile(entry) {
			infos = append(infos, info.Name())
		}
	}

	return infos, nil
}

func isConfigFile(info os.DirEntry) bool {
	if !info.IsDir() && strings.HasPrefix(info.Name(), "config_") && (strings.HasSuffix(info.Name(), ".yaml") || strings.HasSuffix(info.Name(), ".yml")) {
		return true
	}

	return false
}

type ConfigEntry interface {
	GetCurrentActiveConfig() (string, error)
	GetDefaultConfigDir() (string, error)
	GetCoreDir() (string, error)
	GetAllConfig() ([]string, error)
	CreateConfigFile(fileName string, makeActive bool) error
}

var _ ConfigEntry = (*FileConfigEntry)(nil)

type FileConfigEntry struct{}

func (f FileConfigEntry) CreateConfigFile(fileName string, makeActive bool) error {
	dir, err := f.GetDefaultConfigDir()
	if err != nil {
		return err
	}

	err = CreateConfigFile(dir, fileName)
	if err != nil || !makeActive {
		return err
	}

	v, err := CreateCoreViper()
	if err != nil {
		return err
	}

	v.Set(CurrentConfig, fileName)

	return v.WriteConfig()
}

func (f FileConfigEntry) GetCurrentActiveConfig() (string, error) {
	v, err := CreateCoreViper()
	if err != nil {
		return "", err
	}

	return v.GetString(CurrentConfig), nil
}

func (f FileConfigEntry) GetDefaultConfigDir() (string, error) {
	return GetDefaultConfigDir()
}

func (f FileConfigEntry) GetCoreDir() (string, error) {
	return GetCoreDir()
}

func (f FileConfigEntry) GetAllConfig() ([]string, error) {
	configDir, err := f.GetDefaultConfigDir()
	if err != nil {
		return nil, err
	}

	return GetAllConfig(configDir)
}

// CreateFile creates a file in a given directory is it does not exist.
func CreateFile(dir, file string) error {
	result := filepath.Join(dir, file)

	_, err := os.Stat(result)
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	f, err := os.Create(result)
	if err != nil {
		return err
	}

	return f.Close()
}
