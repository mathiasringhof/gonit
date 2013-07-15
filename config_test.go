package main

import "testing"

const EXPECTED_WEBDIR = "./web"
const EXPECTED_PORT = ":9090"
const EXPECTED_REPODIR = "/opt/gitrepos"

func TestLoadCorrectDataset(t *testing.T) {
	config := getTestConfig(EXPECTED_WEBDIR, EXPECTED_PORT, EXPECTED_REPODIR, t)
	if config.WebDir != "./web" {
		t.Errorf("WebDir property does not match parsed JSON, got %s, expected %s", config.WebDir, "./web")
	}
}

func TestLoadPortWithoutColon(t *testing.T) {
	const PORT_WITHOUT_COLON = "9090"
	config := getTestConfig(EXPECTED_WEBDIR, PORT_WITHOUT_COLON, EXPECTED_REPODIR, t)
	if config.Port != EXPECTED_PORT {
		t.Errorf("Failed correctly loading port without colon, expected to add the colon. Got %s, expected %s", config.Port, EXPECTED_PORT)
	}
}

func TestLoadRepositoryDirWithTrailingSlash(t *testing.T) {
	const REPODIR_WITH_SLASH = EXPECTED_REPODIR + "/"
	config := getTestConfig(EXPECTED_WEBDIR, EXPECTED_PORT, REPODIR_WITH_SLASH, t)
	if config.RepositoryDir != REPODIR_WITH_SLASH[:len(REPODIR_WITH_SLASH)-1] {
		t.Errorf("Failed to correctly load repository dir with trailing slash, expected to remove the slash. Got %s, expected %s.", config.RepositoryDir, EXPECTED_REPODIR)
	}
}

func getTestJSON(webdir, port, repodir string) []byte {
	return []byte(`{
						"WebDir": "` + webdir + `",
						"Port": "` + port + `",
						"RepositoryDir": "` + repodir + `"
					}`)
}

func getTestConfig(webdir, port, repodir string, t *testing.T) *Config {
	config := Config{}
	json := getTestJSON(webdir, port, repodir)
	if err := config.Load(json); err != nil {
		t.Errorf("Failed loading JSON with the following error: %s", err.Error())
	}
	return &config
}
