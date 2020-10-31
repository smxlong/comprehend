package comprehend

import (
	"errors"

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
