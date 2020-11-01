package comprehend

import (
	"github.com/smxlong/comprehend/pkg/render"
	"github.com/smxlong/xonnex/pkg/graph"
)

// Style is the style for graphviz rendering with Comprehend.
type Style struct {
	base *render.DefaultStyle
}

// NewStyle creates a new Style.
func NewStyle() *Style {
	return &Style{
		base: render.NewDefaultStyle(),
	}
}

// NodePen returns the pen to use for drawing nodes.
func (s *Style) NodePen(g *graph.Graph, n *graph.Node) render.Pen {
	return s.base.NodePen(g, n)
}

// EdgePen returns the pen to use for drawing edges.
func (s *Style) EdgePen(g *graph.Graph, e *graph.Edge) render.Pen {
	return s.base.EdgePen(g, e)
}
