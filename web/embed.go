package webassets

import "embed"

//go:embed index.html css images js
var FS embed.FS
