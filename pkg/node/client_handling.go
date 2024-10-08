package node

import (
	"errors"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

func (c *Client) getMessageHandler(em message.EncodedMessage) (cb protocol.Callback) {

	if protocol.HasNoExpectedResponse(em.Type()) {
		return c.handlers[em.Type()]
	}

	return <-c.callbacks
}

func (c *Client) handleInboundMessage(em message.EncodedMessage) (err error) {

	var cb protocol.Callback

	if cb = c.getMessageHandler(em); cb == nil {
		return errors.New("handler for '" + protocol.TypeAsString(em.Type()) + "' not found")
	}

	if err = cb(em); err != nil {
		return
	}

	return nil
}
