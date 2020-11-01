package render

import (
	"errors"

	"github.com/smxlong/xonnex/pkg/graph"
)

// ErrorUnsupportedMimeType is returned when the requested MIME type is unsupported.
var ErrorUnsupportedMimeType = errors.New("unsupported MIME type")

// Renderer is an interface that lets you render a xonnex graph into a media
// object.
type Renderer interface {
	RenderToMediaObject(g *graph.Graph, mimeType string) ([]byte, error)
}
