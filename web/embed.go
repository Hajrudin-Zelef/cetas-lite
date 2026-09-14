package webassets

import "embed"

//go:embed index.html favicon.svg css js
var FS embed.FS
