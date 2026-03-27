package common

import (
	_ "embed"
	"encoding/json"
	"os"
)

type PluginPlatform struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

type PluginManifest struct {
	Name        string           `json:"name"`
	Version     string           `json:"version"`
	Description string           `json:"description"`
	Author      string           `json:"author"`
	License     string           `json:"license"`
	Repo        string           `json:"repo"`
	Doc         string           `json:"doc"`
	Platforms   []PluginPlatform `json:"platforms"`
	Tags        []string         `json:"tags"`
}

func LoadManifestByByte(raw []byte) (*PluginManifest, error) {
	var m PluginManifest

	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func LoadManifestByString(raw string) (*PluginManifest, error) {
	return LoadManifestByByte([]byte(raw))
}

func LoadManifestByFile(path string) (*PluginManifest, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return LoadManifestByByte(data)
}
