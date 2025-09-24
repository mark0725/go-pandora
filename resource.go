package pandora

import "embed"

//go:embed pandora/*
var StaticFiles embed.FS
