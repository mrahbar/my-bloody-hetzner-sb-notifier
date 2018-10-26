package writer

import (
	"fmt"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/hetzner"
	"io"
	"text/tabwriter"
)

type TableWriter struct {
	tabWriter *tabwriter.Writer
}

func NewTableWriter(output io.Writer)*TableWriter {
	return &TableWriter{
		tabwriter.NewWriter(output, 0, 8, 2, ' ', tabwriter.Debug|tabwriter.AlignRight),
	}
}

func (c *TableWriter) Print(deals hetzner.Deals) {
	fmt.Fprintf(c.tabWriter,"Got %d offers. Filtered offers: %d\n", deals.ResultStats.OriginalCount, deals.ResultStats.FilteredCount)

	fmt.Fprintf(c.tabWriter, "%s\n", deals.Servers[0].Header())
	for _, server := range deals.Servers {
		fmt.Fprintf(c.tabWriter, "%s\n", server.ToString())
	}
	c.tabWriter.Flush()
}
