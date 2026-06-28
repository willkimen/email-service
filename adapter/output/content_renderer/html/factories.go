package htmlrenderer

import (
	"embed"
)

// templatesFS holds the HTML template files embedded directly into the
// compiled binary. This variable is package-private to ensure the
// encapsulation of the physical infrastructure files.
//
//go:embed templates/*.html
var templatesFS embed.FS

// NewHTMLEmailContentRendererAdapter creates and initializes a new instance of
// HTMLEmailContentRendererAdapter ready for production use.
//
// It automatically injects the embedded file system (templatesFS) containing
// the default HTML files of the package.
func NewHTMLEmailContentRendererAdapter() *HTMLEmailContentRendererAdapter {
	return &HTMLEmailContentRendererAdapter{
		FS: templatesFS,
	}
}
