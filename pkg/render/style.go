package render

import (
	"github.com/smxlong/xonnex/pkg/graph"
)

// RGB represents an RGB color
type RGB struct {
	R, G, B int
}

// Color represents a named or RGB color
type Color struct {
	Name string
	RGB  RGB
}

// NewColor creates a new color with the given name
func NewColor(name string) Color {
	return Color{
		Name: name,
	}
}

// Pen specifies the draw and fill colors of a style
type Pen struct {
	Draw Color
	Fill Color
	Text Color
}

// Style applies styles to graph nodes and edges during rendering.
type Style interface {
	NodePen(g *graph.Graph, n *graph.Node) Pen
	EdgePen(g *graph.Graph, e *graph.Edge) Pen
}
