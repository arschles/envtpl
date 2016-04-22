package main

import (
	"os"
	"strings"
)

func collectEnv() map[string]string {
	envMap := make(map[string]string)
	envStrs := os.Environ()
	for _, envStr := range envStrs {
		spl := strings.SplitN(envStr, "=", 2)
		if len(spl) != 2 {
			continue
		}
		envMap[spl[0]] = spl[1]
	}
	return envMap
}
