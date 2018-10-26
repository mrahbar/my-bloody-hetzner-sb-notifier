package writer

import (
	"encoding/json"
	"fmt"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/hetzner"
	"io"
)

type JsonWriter struct {
	output io.Writer
}

func NewJsonWriter(output io.Writer)*JsonWriter {
	return &JsonWriter{
		output:output,
	}
}

func (c *JsonWriter) Print(deals hetzner.Deals) {
	b, err := json.Marshal(deals)
	if err != nil {
		fmt.Fprintf(c.output, "{\"error\": \"%s\"}", err)
		return
	}
	fmt.Fprint(c.output, string(b))
}
