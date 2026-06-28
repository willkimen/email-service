package htmlrenderer_test

import (
	"emailservice/adapter/output/content_renderer/html"
	"errors"
	"io/fs"
)

var variables = map[string]any{
	"verification_code": "123456",
}

var rendererAdapter = htmlrenderer.NewHTMLEmailContentRendererAdapter()

type MockErroedFS struct{}

func (m MockErroedFS) Open(name string) (fs.File, error) {
	return nil, errors.New("some parse error")
}

var rendererAdapterParseError = &htmlrenderer.HTMLEmailContentRendererAdapter{
	FS: MockErroedFS{},
}
