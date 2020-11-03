package comprehend

// NodeData is the node data and state.
type NodeData struct {
	Labels map[string]string      `json:"labels,omitempty"`
	Type   string                 `json:"type"`
	Data   map[string]interface{} `json:"data"`
	State  map[string]interface{} `json:"-"`
}
