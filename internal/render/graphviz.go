package render

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/smxlong/xonnex/pkg/graph"
)

// GraphViz is a renderer for GraphViz.
type GraphViz struct {
}

// NewGraphViz returns a new GraphViz.
func NewGraphViz() *GraphViz {
	return &GraphViz{}
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

// RenderToMediaObject renders the graph with GraphViz
func (r *GraphViz) RenderToMediaObject(g *graph.Graph, mimeType string) ([]byte, error) {
	buff := &bytes.Buffer{}
	fmt.Fprintln(buff, "digraph {")
	g.ForEachNode(func(g *graph.Graph, n *graph.Node) bool {
		fmt.Fprintf(buff, "  \"%s\";\n", n.Metadata().Name())
		return true
	})
	g.ForEachEdge(func(g *graph.Graph, e *graph.Edge) bool {
		fmt.Fprintf(buff, "  \"%s\" -> \"%s\";\n", e.From().Metadata().Name(), e.To().Metadata().Name())
		return true
	})
	fmt.Fprintln(buff, "}")
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
