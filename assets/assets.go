package assets

import "embed"

// content holds our static web server content.
//
//go:embed *
var Data embed.FS
