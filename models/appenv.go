package models

import (
	"os"
	"os/user"
	"path/filepath"
)

type AppEnv struct {
	
	


}

/**
 * get the path to the directory holding the settings and configs (json format)
 * If the directory not exists, it will be created
 */
func GetPathForFile(filename string) string {
	user, err := user.Current()
	if (err != nil) {
		panic(err)
	}
	configDir := filepath.Join(user.HomeDir,".config","go-agent-smith")
	if _,err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir,os.ModePerm)
	}

	return filepath.Join(configDir,filename)
}