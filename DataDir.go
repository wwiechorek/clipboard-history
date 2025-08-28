package main

import (
	"os"
	"path/filepath"
)

func AppDataDir(appName string) (string, error) {
	base, err := os.UserConfigDir() // diretório de config do usuário
	if err != nil {
		return "", err
	}
	dir := filepath.Join(base, appName)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return dir, nil
}
