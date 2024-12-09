package util

import (
	"os"
	"strings"
)

func ResolveEnv(env string) (string, error) {
    filepath, exists := os.LookupEnv(env+"_FILE")

    if exists {
        content, err := os.ReadFile(filepath)
        if err != nil {
            return "", err
        }
        return strings.TrimSpace(string(content)), nil
    }
    return os.Getenv(env), nil
}
