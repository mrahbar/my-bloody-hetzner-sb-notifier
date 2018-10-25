package hetzner

import (
	"fmt"
	"time"
)

const hetznerlivedataurl = "https://www.hetzner.de/a_hz_serverboerse/live_data.json"

func MakeURL() string {
	return fmt.Sprintf("%s?m=%v", hetznerlivedataurl, time.Now().UnixNano())
}
