/*
Package htmlrenderer provides an infrastructure adapter for rendering HTML email contents.

It leverages Go's standard "html/template" package to securely parse and execute HTML
templates embedded directly into the application binary via the "embed" package.

RECOMMENDATION: Always use the factory function NewHTMLEmailContentRendererAdapter to
instantiate the adapter instead of initializing the struct directly.
Using the factory function guarantees that the production-ready embedded file system (FS)
is properly injected into the adapter. Direct instantiation should be reserved exclusively
for testing environments where you need to inject a mocked or faulted file system.

# Usage Example

	// Recommended approach for production:
	renderer := htmlrenderer.NewHTMLEmailContentRendererAdapter()

	html, err := renderer.Render(message)
*/
package htmlrenderer
