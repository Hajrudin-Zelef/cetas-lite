package webassets

import "embed"

//go:embed index.html css images js fonts
var FS embed.FS
