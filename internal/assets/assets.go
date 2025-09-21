package assets

import _ "embed"

//go:embed profile.html
var ProfileTemplate string

//go:embed redirect.html
var RedirectTemplate string
