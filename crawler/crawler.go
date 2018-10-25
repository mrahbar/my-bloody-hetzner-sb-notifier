package crawler

import (
	"fmt"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/hetzner"
	"os"
	"sort"
	"text/tabwriter"
)

type Crawler struct {
	tabWriter *tabwriter.Writer

	minPrice float64
	maxPrice float64

	minRam int
	maxRam int

	minHddSize int
	maxHddSize int

	minHddCount int
	maxHddCount int

	minBenchmark int
	maxBenchmark int
}

func NewCrawler(minPrice float64, maxPrice float64, minRam int, maxRam int, minHddSize int, maxHddSize int, minHddCount int, maxHddCount int, minBenchmark int, maxBenchmark int) *Crawler {
	crawler := &Crawler{
		tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.Debug|tabwriter.AlignRight),
		minPrice,
		maxPrice,
		minRam,
		maxRam,
		minHddSize,
		maxHddSize,
		minHddCount,
		maxHddCount,
		minBenchmark,
		maxBenchmark,
	}

	return crawler
}

func (c *Crawler) Filter(servers []hetzner.Server) []hetzner.Server {
	var filteredServers []hetzner.Server
	for _, server := range servers {
		if !c.isFiltered(server) {
			filteredServers = append(filteredServers, server)
		}
	}

	sort.Slice(servers, func(i, j int) bool {
		return servers[i].Score() > servers[j].Score()
	})
	return filteredServers
}

func (c *Crawler) Print(servers []hetzner.Server) {
	fmt.Fprintf(c.tabWriter, "%s\n", servers[0].Header())
	for _, server := range servers {
		fmt.Fprintf(c.tabWriter, "%s\n", server.ToString())
	}
	c.tabWriter.Flush()
}

func (c *Crawler) isFiltered(server hetzner.Server) bool {
	filtered := true

	priceParsed := server.ParsePrice()
	if server.CpuBenchmark >= c.minBenchmark && server.CpuBenchmark <= c.maxBenchmark &&
		priceParsed >= c.minPrice && priceParsed <= c.maxPrice &&
		server.Ram >= c.minRam && server.Ram <= c.maxRam &&
		server.TotalHdd() >= c.minHddSize && server.TotalHdd() <= c.maxHddSize &&
		server.HddCount >= c.minHddCount && server.HddCount <= c.maxHddCount {
		filtered = false
	}

	return filtered
}
