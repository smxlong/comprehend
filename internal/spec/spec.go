package spec

// Dependency is a data structure describing a reference to a dependency.
type Dependency struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data,omitempty"`
}

// Node is a data structure describing a dependency Node.
type Node struct {
	Name         string                 `json:"name"`
	Labels       map[string]string      `json:"labels,omitempty"`
	Type         string                 `json:"type"`
	Data         map[string]interface{} `json:"data,omitempty"`
	Dependencies []Dependency           `json:"dependencies,omitempty"`
}

// Group is a group of Nodes
type Group struct {
	Group string `json:"group"`
	Nodes []Node `json:"nodes,omitempty"`
}

// Spec is a data structure describing a dependency Spec, or graph of
// dependencies.
type Spec struct {
	Groups []Group `json:"groups,omitempty"`
}
