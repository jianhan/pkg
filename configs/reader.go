// Package configs provides functions for loading configurations when app boots via viper
package configs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Reader is a single method interface for configs reader.
type Reader interface {
	Read() (*ServiceConfigs, error)
}

// NewReader returns a new interface for ConfigsReader.
func NewReader(env string, extraPaths ...string) Reader {
	return &scfgReader{
		env:        env,
		extraPaths: extraPaths,
	}
}

type scfgReader struct {
	env        string
	extraPaths []string
}

// ServiceConfigs represents common configs for any go-micro service.
type ServiceConfigs struct {
	Name             string
	RegisterTTL      int `mapstructure:"register_ttl"`
	RegisterInterval int `mapstructure:"register_interval"`
	Version          string
	Metadata         map[string]string
}

// Validate checks all required value are set correctly.
func (s *ServiceConfigs) Validate() error {
	// all services must have a name
	if strings.TrimSpace(s.Name) == "" {
		return errors.New("service name can not be empty, please set it in configuration yml file")
	}
	// all service must have a version
	if strings.TrimSpace(s.Version) == "" {
		return errors.New("service version can not be empty, please set it in configuration yml file")
	}
	// TTL is required
	if s.RegisterTTL == 0 {
		return errors.New("TTL to use when registering the service is required and can not be 0, please set it in configuration yml file")
	}
	// interval for re-registering service is required
	if s.RegisterTTL == 0 {
		return errors.New("interval for re-registering service is required and can not be 0, please set it in configuration yml file")
	}
	return nil
}

func (s *scfgReader) Read() (*ServiceConfigs, error) {
	if strings.TrimSpace(s.env) == "" {
		return nil, errors.New("env can not be blank")
	}
	paths := append([]string{"configs/base.yml", fmt.Sprintf("configs/%s.yml", s.env)}, s.extraPaths...)
	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			log.WithFields(log.Fields{
				"config_file_path": path,
			}).Warn("file path not exists")
			return nil, err
		}
		fileName, fileExt := fileInfo.Name(), filepath.Ext(fileInfo.Name())
		fileBaseName := fileName[0 : len(fileName)-len(fileExt)]
		fileDir := filepath.Dir(path)
		viper.SetConfigName(fileBaseName)
		viper.AddConfigPath(fileDir)
		err = viper.MergeInConfig()
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error,
			}).Warn("failed to read/merge configs")
			return nil, err
		}
	}
	c := &ServiceConfigs{}
	if err := viper.UnmarshalKey("service", c); err != nil {
		return nil, err
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}
