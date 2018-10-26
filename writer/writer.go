package writer

import "github.com/mrahbar/my-bloody-hetzner-sb-notifier/hetzner"

type Writer interface {
	Print(deals hetzner.Deals)
}
