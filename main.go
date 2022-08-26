package main

import (
	"embed"
	"github.com/NubeIO/rubix-edge-bios/cmd"
	"github.com/NubeIO/rubix-edge-bios/release"
)

//go:embed systemd/*
var systemdFs embed.FS

//go:embed VERSION
var versionFs embed.FS

func main() {
	cmd.SystemdFs = systemdFs
	release.VersionFs = versionFs
	cmd.Execute()
}
