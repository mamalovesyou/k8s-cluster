package generator

import (
	"os"
	"path"
)

const (
	terraformDir       = "terraform"
	ansibleDir         = "ansible"
	ansibleGroupVarDir = "ansible/group_vars"
	deploymentsDir     = "deployments"
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
	if err := createDir(path.Join(clusterDir, ansibleDir)); err != nil {
		return err
	}
	if err := createDir(path.Join(clusterDir, ansibleGroupVarDir)); err != nil {
		return err
	}
	if err := createDir(path.Join(clusterDir, deploymentsDir)); err != nil {
		return err
	}
	return nil
}
