package comprehend

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/smxlong/comprehend/pkg/render"

	"github.com/smxlong/comprehend/internal/spec"
	"gopkg.in/yaml.v2"
)

// Main is the main function of Comprehend.
func Main(args []string) {
	input, err := os.Open(args[1])
	defer input.Close()
	if err != nil {
		panic(err)
	}
	spec, err := loadSpec(input)
	if err != nil {
		panic(err)
	}
	c, err := NewComprehend(spec)
	if err != nil {
		panic(err)
	}
	r := render.NewGraphViz(NewStyle())
	bytes, err := r.RenderToMediaObject(c.graph, "text/vnd.graphviz")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(bytes)
}

func loadSpec(r io.Reader) (spec.Spec, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return spec.Spec{}, err
	}
	s := spec.Spec{}
	err = yaml.Unmarshal(bytes, &s)
	if err != nil {
		return spec.Spec{}, err
	}
	return s, nil
}
