package comprehend

import (
	"errors"

	"github.com/smxlong/comprehend/internal/spec"
	"github.com/smxlong/xonnex/pkg/graph"
)

// Comprehend is the main data structure.
type Comprehend struct {
	graph *graph.Graph
}

// ErrorNodeAlreadyDeclared error
var ErrorNodeAlreadyDeclared = errors.New("the node has already been declared")

// ErrorNodeNotDeclared error
var ErrorNodeNotDeclared = errors.New("the node has not been declared")

// NewComprehend initialized a Comprehend from a spec.Spec.
func NewComprehend(s spec.Spec) (*Comprehend, error) {
	c := &Comprehend{
		graph: graph.NewGraph(),
	}
	err := c.buildGraph(s)
	if err != nil {
		return nil, err
	}
	return c, nil
}
