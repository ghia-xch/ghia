package main

import (
	"context"
	"github.com/ghia-xch/ghia/pkg/node"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/codec"
	"github.com/ghia-xch/ghia/pkg/node/protocol/full_node"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"github.com/ghia-xch/ghia/pkg/peer"
	l "github.com/sirupsen/logrus"
	"time"
)

func main() {

	c := make(chan struct{}, 0)

	l.Info("Hello World!")

	var client *node.Client

	client = node.NewClient(peer.NewPeerInfo("192.168.8.117", 8444))

	client.Handle(
		protocol.Handler(
			protocol.NewPeak,
			func(em message.EncodedMessage) (err error) {

				var np full_node.NewPeak

				if err = codec.Decode(&np, em); err != nil {
					return err
				}

				l.Info("new peak found: ", np.Height, " [", np.HeaderHash.String(), "] ")
				l.Info("-- weight: ", np.Weight.String())
				l.Info("-- fork: ", np.ForkPointWithPreviousPeak)
				l.Info("-- unfinished block hash: ", np.UnfinishedRewardBlockHash.String())

				return err
			},
		),
	)

	var err error

	if err = client.Open(context.Background(), 10*time.Second); err != nil {
		l.Fatalln(err)
		return
	}

	<-c
}
