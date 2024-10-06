// file: api/config/config.go
package config

import "os"

func IsDevelopmentMode() bool {
	return os.Getenv("APP_ENV") == "development"
}
