package util

import "os"

func GetEnv(key, defaultValue string) string {
	env := os.Getenv(key)
	if len(env) == 0 {
		return defaultValue
	}
	return env
}
