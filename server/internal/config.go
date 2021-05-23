package internal

import (
	dotenv "github.com/joho/godotenv"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func Setup() error {
	return dotenv.Load(basepath + "/.env")
}
