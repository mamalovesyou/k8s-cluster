package generator

import (
	"os"
	"path"
)

const (
	terraformDir = "terraform"
	playbookDir  = "playbooks"
	confDir      = "conf"
)

// Create a path directory if not exist
func createDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

func CreateAllDirs(destination string, clusterName string) error {
	clusterDir := path.Join(destination, clusterName)
	if err := createDir(clusterDir); err != nil {
		return err
	}
	if err := createDir(path.Join(clusterDir, terraformDir)); err != nil {
		return err
	}
	if err := createDir(path.Join(clusterDir, playbookDir)); err != nil {
		return err
	}
	if err := createDir(path.Join(clusterDir, confDir)); err != nil {
		return err
	}
	return nil
}
