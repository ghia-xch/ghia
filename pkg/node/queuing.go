package node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/gorilla/websocket"
)

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
