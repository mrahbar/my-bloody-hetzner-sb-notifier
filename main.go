package main

import (
	"flag"
	"fmt"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/client"
	c "github.com/mrahbar/my-bloody-hetzner-sb-notifier/crawler"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/hetzner"
	n "github.com/mrahbar/my-bloody-hetzner-sb-notifier/notifier"
)

var (
	minPrice = flag.Float64(
		"min-price",
		0,
		"set min price",
	)
	maxPrice = flag.Float64(
		"max-price",
		297,
		"set max price",
	)

	minRam = flag.Int(
		"min-ram",
		0,
		"set min ram",
	)
	maxRam = flag.Int(
		"max-ram",
		256,
		"set max ram",
	)

	minHddSize = flag.Int(
		"min-hdd-size",
		0,
		"set min hdd size",
	)
	maxHddSize = flag.Int(
		"max-hdd-size",
		6144,
		"set max hdd size",
	)

	minHddCount = flag.Int(
		"min-hdd-count",
		0,
		"set min hdd count",
	)
	maxHddCount = flag.Int(
		"max-hdd-count",
		15,
		"set max hdd count",
	)

	minBenchmark = flag.Int(
		"min-benchmark",
		0,
		"set min benchmark",
	)
	maxBenchmark = flag.Int(
		"max-benchmark",
		20000,
		"set max benchmark",
	)

	notifierSender = flag.String(
		"notifier-sender",
		"",
		"set notifier sender",
	)

	notifierPassword = flag.String(
		"notifier-password",
		"",
		"set notifier password",
	)

	notifierRecipient = flag.String(
		"notifier-recipient",
		"",
		"set notifier recipient",
	)

	alertOnScore = flag.Int(
		"alert-on-score",
		0,
		"set alert on score",
	)
)

func main() {
	flag.Parse()

	offers := &hetzner.Offers{}
	err := client.NewClient().DoRequest(hetzner.MakeUrl(), offers)
	if err != nil {
		panic(fmt.Errorf("failed to get hetzner live data: %s", err))
	}

	if len(offers.Server) > 0 {
		crawler := c.NewCrawler(*minPrice, *maxPrice, *minRam, *maxRam, *minHddSize, *maxHddSize, *minHddCount, *maxHddCount, *minBenchmark, *maxBenchmark)
		servers := crawler.Filter(offers.Server)

		fmt.Printf("Got %d offers. Filtered offers: %d\n", len(offers.Server), len(servers))
		crawler.Print(servers)

		notifier := n.NewNotifier(*notifierRecipient, *notifierSender, *notifierPassword)
		notifier.Act(servers)
	} else {
		fmt.Println("Got no offers.")
	}
}
