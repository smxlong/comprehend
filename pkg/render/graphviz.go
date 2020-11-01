package render

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/smxlong/xonnex/pkg/graph"
)

// GraphViz is a renderer for GraphViz.
type GraphViz struct {
	style Style
}

// NewGraphViz returns a new GraphViz.
func NewGraphViz(s Style) *GraphViz {
	return &GraphViz{
		style: s,
	}
}

func dotColor(color Color) string {
	if color.Name != "" {
		return color.Name
	}
	return fmt.Sprintf("\"%02x%02x%02x\"", color.RGB.R, color.RGB.G, color.RGB.B)
}

func (r *GraphViz) render(buff *bytes.Buffer, format string) ([]byte, error) {
	cmd := exec.Command("dot", []string{
		"-T" + format,
	}...)
	cmd.Stdin = buff
	outBuff := &bytes.Buffer{}
	cmd.Stdout = outBuff
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return outBuff.Bytes(), nil
}

func (r *GraphViz) generateDotLanguage(g *graph.Graph, buff *bytes.Buffer) {
	fmt.Fprintln(buff, "digraph {")
	g.ForEachNode(func(g *graph.Graph, n *graph.Node) bool {
		p := r.style.NodePen(g, n)
		fmt.Fprintf(buff, "  \"%s\" [color=%s, fillcolor=%s, fontcolor=%s];\n", n.Metadata().Name(),
			dotColor(p.Draw), dotColor(p.Fill), dotColor(p.Draw))
		return true
	})
	g.ForEachEdge(func(g *graph.Graph, e *graph.Edge) bool {
		fmt.Fprintf(buff, "  \"%s\" -> \"%s\";\n", e.From().Metadata().Name(), e.To().Metadata().Name())
		return true
	})
	fmt.Fprintln(buff, "}")
}

// RenderToMediaObject renders the graph with GraphViz
func (r *GraphViz) RenderToMediaObject(g *graph.Graph, mimeType string) ([]byte, error) {
	buff := &bytes.Buffer{}
	r.generateDotLanguage(g, buff)
	switch mimeType {
	case "text/vnd.graphviz":
		return buff.Bytes(), nil
	case "image/png":
		return r.render(buff, "png")
	case "application/pdf":
		return r.render(buff, "pdf")
	}
	return nil, ErrorUnsupportedMimeType
}
