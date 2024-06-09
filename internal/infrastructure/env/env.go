package env

import (
	"fmt"
	"os"
	"strconv"
)

// All variables for project
var (
	LogLevel = Getter("LOG_LEVEL", "debug")
	Endpoint = fmt.Sprintf("%s:%s", Getter("SERVER_HOST", ""), Getter("SERVER_PORT", "8080"))
)

// Getter -
func Getter(key, defaultValue string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	return defaultValue
}

// GetterInt -
func GetterInt(key string, defaultValue int) int {
	env, ok := os.LookupEnv(key)
	if ok {
		res, err := strconv.ParseInt(env, 10, 32)
		if err == nil {
			return int(res)
		}
	}
	return defaultValue
}
