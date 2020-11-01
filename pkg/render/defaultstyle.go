package render

import "github.com/smxlong/xonnex/pkg/graph"

// DefaultStyle is a default style.
type DefaultStyle struct {
}

// NewDefaultStyle returns a new DefaultStyle.
func NewDefaultStyle() *DefaultStyle {
	return &DefaultStyle{}
}

// NodePen returns the pen to use to draw a node.
func (s *DefaultStyle) NodePen(g *graph.Graph, n *graph.Node) Pen {
	return Pen{
		Draw: NewColor("black"),
		Fill: NewColor("white"),
		Text: NewColor("black"),
	}
}

// EdgePen returns the pen to use to draw an edge.
func (s *DefaultStyle) EdgePen(g *graph.Graph, e *graph.Edge) Pen {
	return Pen{
		Draw: NewColor("black"),
		Fill: NewColor("white"),
		Text: NewColor("black"),
	}
}
