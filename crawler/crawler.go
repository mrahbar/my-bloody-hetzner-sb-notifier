package crawler

import (
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/hetzner"
	"sort"
)

type Parameter struct {
	MinPrice float64
	MaxPrice float64

	MinRam int64
	MaxRam int64

	MinHddSize int64
	MaxHddSize int64

	MinHddCount int64
	MaxHddCount int64

	MinBenchmark int64
	MaxBenchmark int64
}


type Crawler struct {
	parameter Parameter
}

func NewCrawler(parameter Parameter) *Crawler {
	crawler := &Crawler{
		parameter: parameter,
	}

	return crawler
}

func (c *Crawler) Filter(servers []hetzner.Server) hetzner.Deals {
	var filteredServers []hetzner.Server
	for _, server := range servers {
		if !c.isFiltered(server) {
			filteredServers = append(filteredServers, server)
		}
	}

	sort.Slice(servers, func(i, j int) bool {
		return servers[i].Score() > servers[j].Score()
	})

	deals := hetzner.Deals{
		ResultStats: hetzner.FilterResultStats{OriginalCount: len(servers), FilteredCount: len(filteredServers)},
		Servers: filteredServers,
	}
	return deals
}

func (c *Crawler) isFiltered(server hetzner.Server) bool {
	filtered := true

	priceParsed := server.ParsePrice()
	if server.CpuBenchmark >= c.parameter.MinBenchmark && server.CpuBenchmark <= c.parameter.MaxBenchmark &&
		priceParsed >= c.parameter.MinPrice && priceParsed <= c.parameter.MaxPrice &&
		server.Ram >= c.parameter.MinRam && server.Ram <= c.parameter.MaxRam &&
		server.TotalHdd() >= c.parameter.MinHddSize && server.TotalHdd() <= c.parameter.MaxHddSize &&
		server.HddCount >= c.parameter.MinHddCount && server.HddCount <= c.parameter.MaxHddCount {
		filtered = false
	}

	return filtered
}
