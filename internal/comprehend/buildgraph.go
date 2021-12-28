package comprehend

import (
	"github.com/smxlong/comprehend/internal/spec"
	"github.com/smxlong/xonnex/pkg/graph"
)

// buildGraph validates and transforms the spec.Spec into a graph
func (c *Comprehend) buildGraph(s spec.Spec) error {
	// Index all nodes by name and add them to the graph.
	nodesByName := map[string]*graph.Node{}
	for _, group := range s.Groups {
		for _, node := range group.Nodes {
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
	}
	// Look for any undeclared nodes in the dependencies
	for _, group := range s.Groups {
		for _, node := range group.Nodes {
			for _, dependency := range node.Dependencies {
				if _, valid := nodesByName[dependency.Name]; !valid {
					n := graph.NewNodeWithName(dependency.Name)
					nodesByName[dependency.Name] = n
					err := c.graph.AddNode(n)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	// Create the graph edges
	for _, group := range s.Groups {
		for _, node := range group.Nodes {
			for _, dependency := range node.Dependencies {
				err := c.graph.AddEdge(graph.NewEdgeDirectional(nodesByName[dependency.Name], nodesByName[node.Name]))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
