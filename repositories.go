package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

type Repository struct {
	Name string
}

func getRepositories() *[]Repository {
	var result []Repository
	dirs, _ := ioutil.ReadDir(config.RepositoryDir)
	for _, entry := range dirs {
		if entry.IsDir() {
			result = append(result, Repository{entry.Name()})
		}
	}
	return &result
}

func (repo *Repository) Create() error {
	cmd := exec.Command("git", "init", "--bare", repo.Name)
	cmd.Dir = config.RepositoryDir
	return cmd.Run()
}

func (repo *Repository) Path() string {
	return config.RepositoryDir + "/" + repo.Name
}

func (repo *Repository) Exists() bool {
	if _, err := os.Stat(repo.Path()); os.IsNotExist(err) {
		return false
	}
	return true
}
