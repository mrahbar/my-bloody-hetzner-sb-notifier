package hetzner

import (
	"fmt"
	"strconv"
	"strings"
)

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
	Bandwith    int      `json:"bandwith"`

	Price       string `json:"price"`
	Price_v     string `json:"price_v"`
	Setup_price string `json:"setup_price"`

	Cpu           string `json:"cpu"`
	Cpu_benchmark int    `json:"cpu_benchmark"`
	Cpu_count     int    `json:"cpu_count"`
	Ram           int    `json:"ram"`
	Ram_hr        string `json:"ram_hr"`
	Hdd_size      int    `json:"hdd_size"`
	Hdd_hr        string `json:"hdd_hr"`
	Hdd_count     int    `json:"hdd_count"`
	SpecialHdd    string `json:"specialHdd"`

	Next_reduce    int    `json:"next_reduce"`
	Next_reduce_hr string `json:"next_reduce_hr"`

	Fixed_price bool `json:"fixed_price"`
	Is_highio   bool `json:"is_highio"`
	Is_ecc      bool `json:"is_ecc"`
}

func (s *Server) TotalHdd() int {
	return s.Hdd_count*s.Hdd_size
}

func (s *Server) Score() float64 {
	return (float64(s.TotalHdd())*0.2 + float64(s.Ram)*0.4 + float64(s.Cpu_benchmark)*0.4)/s.ParsePrice()
}

func (s *Server) ParsePrice() float64 {
	priceParsed, err := strconv.ParseFloat(s.Price, 32)
	if err != nil {
		fmt.Printf("Could not parse price %s for server %s: %s", s.Price, s.Key, err)
		return -1
	}
	return priceParsed*1.19
}

func (s *Server) Header() string {
	return fmt.Sprint("ID\tRam\tHDD\tCPU\tPrice\tScore\tReduce time\tSpecials")
}

func (s *Server) ToString() string {
	fixedPriceSymbol := "*"
	if !s.Fixed_price {
		fixedPriceSymbol = ""
	}
	specials := strings.Join(s.Specials, ", ")
	return fmt.Sprintf("%s-%d\t%s\t%s (%d)\t%s (%d)\t%.2f €%s\t%.2f\t%s\t%s", s.Name, s.Key, s.Ram_hr, s.Hdd_hr,  s.TotalHdd(), s.Cpu, s.Cpu_benchmark, s.ParsePrice(), fixedPriceSymbol, s.Score(), s.Next_reduce_hr, specials)
}
