package comprehend

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/smxlong/comprehend/pkg/render"
	"github.com/spf13/cobra"

	"github.com/smxlong/comprehend/internal/spec"
	"gopkg.in/yaml.v2"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "comprehend <input file>",
		Short: "interpret a dependency list and spit out a graph",
		RunE:  Main,
		Args:  cobra.ExactArgs(1),
	}
	cmd.Flags().StringP("mime-type", "t", "text/vnd.graphviz", "MIME type to generate")
	return cmd
}

// Main is the main function of Comprehend.
func Main(cmd *cobra.Command, args []string) error {
	input, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer input.Close()
	spec, err := loadSpec(input)
	if err != nil {
		return err
	}
	c, err := NewComprehend(spec)
	if err != nil {
		return err
	}
	r := render.NewGraphViz(NewStyle())
	mimeType, _ := cmd.Flags().GetString("mime-type")
	bytes, err := r.RenderToMediaObject(c.graph, mimeType)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(bytes)
	return err
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
