package node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/gorilla/websocket"
	"time"
)

func (c *Client) inboundQueuing() {

	l.Infoln("starting inbound queuing")

	var err error
	var em []byte

	for {

		if _, em, err = c.conn.ReadMessage(); err != nil {

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {

				l.Errorln("error reading from connection: %v", err)

				c.isClosing <- true

				return
			}

			continue
		}

		c.inbound <- em
	}
}

const PingInterval = 60 * time.Second

func (c *Client) outboundQueuing() {

	l.Infoln("starting outbound queuing")

	var err error
	var em protocol.EncodedMessage
	var ok bool

	ticker := time.NewTicker(PingInterval)

	defer ticker.Stop()

	for {
		select {
		case em, ok = <-c.outbound:

			if !ok {
				continue
			}

			if err = c.conn.WriteMessage(websocket.BinaryMessage, em); err != nil {
				l.Errorln("error writing to connection: %v", err)
			}

		case <-ticker.C:

			l.Infoln("sending ping")

			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				l.Errorln("error writing ping to connection: %v", err)
				return
			}

			ticker.Reset(10 * time.Second)
		}
	}
}

func (c *Client) drainInboundQueue(dec *protocol.MessageDecoder) (err error) {

	var inboundLen int

	inboundLen = len(c.inbound)

	for i := 0; i < inboundLen; i++ {
		if err = c.handleInboundMessage(dec, <-c.inbound); err != nil {
			return
		}
	}

	return nil
}

func (c *Client) drainOutboundQueue() (err error) {

	var outboundLen int

	outboundLen = len(c.outbound)

	for i := 0; i < outboundLen; i++ {
		if err = c.conn.WriteMessage(websocket.BinaryMessage, <-c.outbound); err != nil {
			l.Errorln("error writing to connection: %v", err)
		}
	}

	return nil
}
