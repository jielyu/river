package models

type TOMLConfig struct {
	Name        string `toml:"name"`
	ProjectType string `toml:"project_type"`
}
