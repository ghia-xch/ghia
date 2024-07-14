package protocol

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia/ghia/pkg"
	"github.com/ghia/ghia/pkg/protocol/capability"
	"github.com/ghia/ghia/pkg/protocol/full_node"
	"github.com/ghia/ghia/pkg/protocol/message"
	"github.com/ghia/ghia/pkg/protocol/network"
	"github.com/ghia/ghia/pkg/protocol/node"
	"github.com/gorilla/websocket"
)

var DefaultHandshake = &Handshake{
	NetworkId:       network.DefaultNetwork,
	ProtocolVersion: full_node.ProtocolVersion,
	SoftwareVersion: pkg.Version,
	ServerPort:      8444,
	NodeType:        node.FullNode,
	Capabilities:    capability.DefaultSet,
}

type Handshake struct {
	NetworkId       network.Network
	ProtocolVersion string
	SoftwareVersion string
	ServerPort      uint16
	NodeType        node.Type
	Capabilities    capability.Set
}

func (h *Handshake) Encode(enc *message.MessageEncoder) (em message.EncodedMessage, err error) {

	if em, err = enc.Encode(
		string(h.NetworkId),
		string(h.ProtocolVersion),
		string(h.SoftwareVersion),
		h.ServerPort,
		uint8(h.NodeType),
		h.Capabilities,
	); err != nil {
		return
	}

	return enc.Bytes(), nil
}

func (h *Handshake) Decode(em message.EncodedMessage) (err error) {
	return nil
}

func PerformHandshake(conn *websocket.Conn, enc *message.MessageEncoder, h1 *Handshake) (h2 *Handshake, err error) {

	var em message.EncodedMessage

	enc.Reset(HandshakeType, nil)

	if em, err = h1.Encode(enc); err != nil {
		return nil, err
	}

	spew.Dump(em)

	if err = conn.WriteMessage(websocket.BinaryMessage, em); err != nil {
		return nil, err
	}

	if _, em, err = conn.ReadMessage(); err != nil {
		return nil, err
	}

	spew.Dump(em)

	//h2.Decode(enc, em)

	return nil, nil
}
