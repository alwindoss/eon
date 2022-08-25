package eon

import "time"

type Config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3030"`
	Password     string        `env:"PASSWORD,unset"`
	IsProduction bool          `env:"PRODUCTION"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
	DatabaseURL  string        `env:"DATABASE_URL"`
}
