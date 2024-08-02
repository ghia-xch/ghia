package peer

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
	"github.com/gorilla/websocket"
)

type Peer struct {
	conn      websocket.Conn
	inbound   []primitive.EncodedMessage
	outbound  []primitive.EncodedMessage
	callbacks []primitive.Callback
}

func (p *Peer) Send(em primitive.EncodedMessage) (err error) {
	return err
}

func (p *Peer) SendWithCallback(em primitive.EncodedMessage, cb primitive.Callback) (err error) {
	return err
}

func (p *Peer) Close() (err error) {

	err = p.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)

	if err != nil {
		return
	}

	return nil
}
