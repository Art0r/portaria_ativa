package internal

import (
	"path/filepath"
)

type Config struct {
	ViewsDir  string
	StaticDir string
}

var AppConfig = Config{
	ViewsDir:  "./public/views",
	StaticDir: "./public/static",
}

func GetViewsPath(template string) string {
	return filepath.Join(AppConfig.ViewsDir, template)
}
