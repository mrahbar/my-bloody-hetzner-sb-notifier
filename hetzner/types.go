package hetzner

import (
	"fmt"
	"strconv"
	"strings"
)

type Deals struct {
	ResultStats FilterResultStats
	Servers []Server
}

type FilterResultStats struct {
	OriginalCount int `json:"count"`
	FilteredCount int `json:"filtered"`
}

type Offers struct {
	Server []Server `json:"server"`
}

type Server struct {
	Key         int      `json:"key"`
	Name        string   `json:"name"`
	Freetext    string   `json:"freetext"`
	Description []string `json:"description"`
	Dist        []string `json:"dist"`
	Datacenter  []string `json:"datacenter"`
	Specials    []string `json:"specials"`
	Traffic     string   `json:"traffic"`
	Bandwidth   int      `json:"bandwith"`

	Price      string `json:"price"`
	PriceV     string `json:"price_v"`
	SetupPrice string `json:"setup_price"`

	Cpu          string `json:"cpu"`
	CpuBenchmark int64    `json:"cpu_benchmark"`
	CpuCount     int64    `json:"cpu_count"`
	Ram          int64    `json:"ram"`
	RamHr        string `json:"ram_hr"`
	HddSize      int64    `json:"hdd_size"`
	HddHr        string `json:"hdd_hr"`
	HddCount     int64    `json:"hdd_count"`
	SpecialHdd   string `json:"specialHdd"`

	NextReduce   int    `json:"next_reduce"`
	NextReduceHr string `json:"next_reduce_hr"`

	FixedPrice bool `json:"fixed_price"`
	IsHighio   bool `json:"is_highio"`
	IsEcc      bool `json:"is_ecc"`
}

func (s *Server) TotalHdd() int64 {
	return s.HddCount * s.HddSize
}

func (s *Server) Score() float64 {
	return (float64(s.TotalHdd())*0.2 + float64(s.Ram)*0.4 + float64(s.CpuBenchmark)*0.4) / s.ParsePrice()
}

func (s *Server) ParsePrice() float64 {
	priceParsed, err := strconv.ParseFloat(s.Price, 32)
	if err != nil {
		fmt.Printf("Could not parse price %s for server %d: %s", s.Price, s.Key, err)
		return -1
	}
	return priceParsed * 1.19
}

func (s *Server) Header() string {
	return fmt.Sprint("ID\tRam\tHDD\tCPU\tPrice\tScore\tReduce time\tSpecials")
}

func (s *Server) ToString() string {
	fixedPriceSymbol := "*"
	if !s.FixedPrice {
		fixedPriceSymbol = ""
	}
	specials := strings.Join(s.Specials, ", ")
	return fmt.Sprintf("%s-%d\t%s\t%s (%d)\t%s (%d)\t%.2f €%s\t%.2f\t%s\t%s", s.Name, s.Key, s.RamHr, s.HddHr, s.TotalHdd(), s.Cpu, s.CpuBenchmark, s.ParsePrice(), fixedPriceSymbol, s.Score(), s.NextReduceHr, specials)
}
