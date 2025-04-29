package main

import (
	_ "embed"
	"encoding/json"
)

//go:embed .release-please-manifest.json
var version_strings []byte

type VersionSchema struct {
	Path string `json:"."`
}

func GetCLIVersion() string {
	var version VersionSchema
	json.Unmarshal(version_strings, &version)
	return version.Path
}
