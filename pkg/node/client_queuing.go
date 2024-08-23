package node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/gorilla/websocket"
	"time"
)

func (c *Client) handlerQueuing() {

	var err error
	var em protocol.EncodedMessage
	var ok bool

	var dec = protocol.NewMessageDecoder()

	for {

		select {
		case em, ok = <-c.inbound:

			if !ok {
				continue
			}

			l.Info("received message[", em.Type(), "] ", protocol.TypeAsString(em.Type()))

			if err = c.handleInboundMessage(dec, em); err != nil {
				l.Errorf("error handling inbound message: %v", err)
			}

		case <-c.isClosing:

			l.Infoln("closing inbound queuing")

			if err = c.drainInboundQueue(dec); err != nil {
				l.Errorf("error closing inbound message queue: %v", err)
			}

			l.Infoln("closing outbound queuing")

			if err = c.drainOutboundQueue(); err != nil {
				l.Errorf("error closing outbound message queue: %v", err)
			}

			c.isClosed <- true

			return
		}
	}
}

func (c *Client) inboundQueuing() {

	l.Infoln("starting inbound queuing")

	var err error
	var em []byte

	for {

		if _, em, err = c.conn.ReadMessage(); err != nil {

			l.Errorln("error reading from connection: %v", err)

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {

				l.Errorln("unexpected close: %v", err)

				return
			}

			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {

				l.Infoln("expected close: %v", err)

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

			l.Infoln(protocol.TypeAsString(em.Type()))

			if err = c.conn.WriteMessage(websocket.BinaryMessage, em); err != nil {
				l.Errorln("error writing to connection: %v", err)
				return
			}

		case <-ticker.C:

			l.Infoln("pinging connection")

			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				l.Errorln("error pinging connection: %v", err)
				return
			}

			ticker.Reset(PingInterval)
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
