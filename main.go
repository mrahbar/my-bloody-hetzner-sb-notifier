package main

import (
	"flag"
	"fmt"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/client"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/crawler"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/hetzner"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/strcase"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/writer"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

const (
	flagMinPrice     = "min-price"
	flagMaxPrice     = "max-price"
	flagMinRam       = "min-ram"
	flagMaxRam       = "max-ram"
	flagMinHddSize   = "min-hdd-size"
	flagMaxHddSize   = "max-hdd-size"
	flagMinHddCount  = "min-hdd-count"
	flagMaxHddCount  = "max-hdd-count"
	flagMinBenchmark = "min-cpu-benchmark"
	flagMaxBenchmark = "max-cpu-benchmark"

	flagOutput        = "output"
	flagServeHttpPort = "serve-http-port"
	flagServeHttp     = "serve-http"
)

var (
	zeroIntValue   = int64(0)
	zeroFloatValue = float64(0)

	defaultMaxPriceValue        = float64(297)
	defaultMaxRamValue          = int64(256)
	defaultMaxHddSizeValue      = int64(6144)
	defaultMaxHddCountValue     = int64(15)
	defaultMaxCpuBenchmarkValue = int64(20000)

	minPrice = flag.Float64(
		flagMinPrice,
		0,
		"set min price",
	)
	maxPrice = flag.Float64(
		flagMaxPrice,
		297,
		"set max price",
	)

	minRam = flag.Int64(
		flagMinRam,
		0,
		"set min ram",
	)
	maxRam = flag.Int64(
		flagMaxRam,
		256,
		"set max ram",
	)

	minHddSize = flag.Int64(
		flagMinHddSize,
		0,
		"set min hdd size",
	)
	maxHddSize = flag.Int64(
		flagMaxHddSize,
		6144,
		"set max hdd size",
	)

	minHddCount = flag.Int64(
		flagMinHddCount,
		0,
		"set min hdd count",
	)
	maxHddCount = flag.Int64(
		flagMaxHddCount,
		15,
		"set max hdd count",
	)

	minBenchmark = flag.Int64(
		flagMinBenchmark,
		0,
		"set min benchmark",
	)
	maxBenchmark = flag.Int64(
		flagMaxBenchmark,
		20000,
		"set max benchmark",
	)

	serveHttp = flag.Bool(
		flagServeHttp,
		false,
		"set serve http",
	)

	serveHttpPort = flag.Int(
		flagServeHttpPort,
		8080,
		"set serve http port",
	)

	output = flag.String(
		flagOutput,
		"table",
		"set output: one of table, json",
	)
)

func main() {
	flag.Parse()

	if *serveHttp {
		runHttp()
	} else {
		p := crawler.Parameter{
			MinPrice: *minPrice,
			MaxPrice: *maxPrice,
			MinRam: *minRam,
			MaxRam: *maxRam,
			MinHddSize: *minHddSize,
			MaxHddSize: *maxHddSize,
			MinHddCount: *minHddCount,
			MaxHddCount: *maxHddCount,
			MinBenchmark: *minBenchmark,
			MaxBenchmark: *maxBenchmark,
		}

		run(os.Stdout, p, *output)
	}
}

func runHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := httputil.DumpRequest(r, false)
		if err == nil {
			fmt.Printf("Got request: %s", string(b))
		}
		parameter := crawler.Parameter{
			MinPrice:     zeroFloatValue,
			MaxPrice:     defaultMaxPriceValue,
			MinRam:       zeroIntValue,
			MaxRam:       defaultMaxRamValue,
			MinHddSize:   zeroIntValue,
			MaxHddSize:   defaultMaxHddSizeValue,
			MinHddCount:  zeroIntValue,
			MaxHddCount:  defaultMaxHddCountValue,
			MinBenchmark: zeroIntValue,
			MaxBenchmark: defaultMaxCpuBenchmarkValue,
		}
		output := "table"
		values := r.URL.Query()
		for k, v := range values {
			value := v[0]

			switch k {
			case strcase.ToLowerCamel(flagMinPrice):
				parameter.MinPrice = parseFloatFlag(w, strcase.ToLowerCamel(flagMinPrice), value)
				break
			case strcase.ToLowerCamel(flagMaxPrice):
				parameter.MaxPrice = parseFloatFlag(w, strcase.ToLowerCamel(flagMaxPrice), value)
				break
			case strcase.ToLowerCamel(flagMinRam):
				parameter.MinRam = parseIntFlag(w, strcase.ToLowerCamel(flagMinRam), value)
				break
			case strcase.ToLowerCamel(flagMaxRam):
				parameter.MaxRam = parseIntFlag(w, strcase.ToLowerCamel(flagMaxRam), value)
				break
			case strcase.ToLowerCamel(flagMinHddSize):
				parameter.MinHddSize = parseIntFlag(w, strcase.ToLowerCamel(flagMinHddSize), value)
				break
			case strcase.ToLowerCamel(flagMaxHddSize):
				parameter.MaxHddSize = parseIntFlag(w, strcase.ToLowerCamel(flagMaxHddSize), value)
				break
			case strcase.ToLowerCamel(flagMinHddCount):
				parameter.MinHddCount = parseIntFlag(w, strcase.ToLowerCamel(flagMinHddCount), value)
				break
			case strcase.ToLowerCamel(flagMaxHddCount):
				parameter.MaxHddCount = parseIntFlag(w, strcase.ToLowerCamel(flagMaxHddCount), value)
				break
			case strcase.ToLowerCamel(flagMinBenchmark):
				parameter.MinBenchmark = parseIntFlag(w, strcase.ToLowerCamel(flagMinBenchmark), value)
				break
			case strcase.ToLowerCamel(flagMaxBenchmark):
				parameter.MaxBenchmark = parseIntFlag(w, strcase.ToLowerCamel(flagMaxBenchmark), value)
				break
			case strcase.ToLowerCamel(flagOutput):
				output = value
				break
			}
		}

		w.WriteHeader(http.StatusOK)
		run(w, parameter, output)
	})
	address := fmt.Sprintf(":%d", *serveHttpPort)
	fmt.Printf("Running http server on address %s\n", address)
	fmt.Println(http.ListenAndServe(address, nil))
}

func parseFloatFlag(w http.ResponseWriter, flag string, value string) float64 {
	parseValue, err := strconv.ParseFloat(value, 32)
	if err == nil {
		return parseValue
	} else {
		writeBadRequestResponseForQueryParameter(w, flag, err)
		return zeroFloatValue
	}
}

func parseIntFlag(w http.ResponseWriter, flag string, value string) int64 {
	parseValue, err := strconv.ParseInt(value, 10, 32)
	if err == nil {
		return parseValue
	} else {
		writeBadRequestResponseForQueryParameter(w, flag, err)
		return zeroIntValue
	}
}

func writeBadRequestResponseForQueryParameter(w http.ResponseWriter, parameter string, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf("Error parsing query parameter %s: %s", parameter, err)))
}

func run(w io.Writer, parameter crawler.Parameter, output string) error {
	offers := &hetzner.Offers{}
	err := client.NewClient().DoRequest(hetzner.MakeURL(), offers)
	if err != nil {
		return fmt.Errorf("failed to get hetzner live data: %s", err)
	}
	if len(offers.Server) > 0 {
		c := crawler.NewCrawler(parameter)
		deals := c.Filter(offers.Server)

		switch output {
		case "json":
			writer.NewJsonWriter(w).Print(deals)
			break
		default:
		case "table":
			writer.NewTableWriter(w).Print(deals)
			break
		}
		return nil
	} else {
		return fmt.Errorf("got no offers.")
	}
}
