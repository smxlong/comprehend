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

// buildGraph validates and transforms the spec.Spec into a graph
func (c *Comprehend) buildGraph(s spec.Spec) error {
	// Index all nodes by name and add them to the graph.
	nodesByName := map[string]*graph.Node{}
	for _, node := range s.Nodes {
		if _, valid := nodesByName[node.Name]; valid {
			return ErrorNodeAlreadyDeclared
		}
		n := graph.NewNodeWithName(node.Name)
		nodesByName[node.Name] = n
		err := c.graph.AddNode(n)
		if err != nil {
			return err
		}
	}
	// Look for any undeclared nodes in the dependencies
	for _, node := range s.Nodes {
		for _, dependency := range node.Dependencies {
			if _, valid := nodesByName[dependency.Name]; !valid {
				return ErrorNodeNotDeclared
			}
		}
	}
	// Create the graph edges
	for _, node := range s.Nodes {
		for _, dependency := range node.Dependencies {
			err := c.graph.AddEdge(graph.NewEdgeDirectional(nodesByName[dependency.Name], nodesByName[node.Name]))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
