package domain

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Log LogConfig
}

type LogConfig struct {
	Level            logrus.Level `default:"info" env:"LOG_LEVEL"`
	FilePath         string       `default:"/var/log/cott/cott.log" env:"LOG_FILE_PATH"`
	MaxFileSizeInMb  int          `default:"10" env:"LOG_MAX_FILE_SIZE_IN_MB"`
	MaxFilesCount    int          `default:"7" env:"LOG_MAX_FILES_COUNT"`
	MaxFileAgeInDays int          `default:"7" env:"LOG_MAX_FILE_AGE_IN_DAYS"`
	CompressOldFiles bool         `default:"true" env:"LOG_COMPRESS_OLD_FILES"`
}

func (c *Config) ToJson() ([]byte, error) {
	return json.Marshal(c)
}
