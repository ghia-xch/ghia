package crawler

import "github.com/ghia-xch/ghia/pkg/peer"

type Crawler struct {
	ps peer.Store
}

func Run(ps peer.Store) *Crawler {
	return &Crawler{ps: ps}
}
