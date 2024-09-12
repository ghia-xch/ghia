package node

import (
	"github.com/ghia-xch/ghia/pkg"
	"github.com/ghia-xch/ghia/pkg/network"
	"github.com/ghia-xch/ghia/pkg/node/capability"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/codec"
	"github.com/ghia-xch/ghia/pkg/node/protocol/full_node"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
	"github.com/gorilla/websocket"
)

var (
	DefaultHandshake = &Handshake{
		NetworkId:       network.DefaultNetwork,
		ProtocolVersion: full_node.ProtocolVersion,
		SoftwareVersion: pkg.Version,
		ServerPort:      8444,
		NodeType:        FullNode,
		Capabilities:    capability.DefaultSet,
	}
)

type Handshake struct {
	NetworkId       network.Network
	ProtocolVersion string
	SoftwareVersion string
	ServerPort      uint16
	NodeType        Type
	Capabilities    capability.Set
}

func (h *Handshake) Type() message.MessageType {
	return protocol.HandshakeType
}

func performHandshake(conn *websocket.Conn, h1 *Handshake) (h2 *Handshake, err error) {

	var em message.EncodedMessage

	if em, err = codec.Encode(nil, h1); err != nil {
		return nil, err
	}

	if err = conn.WriteMessage(websocket.BinaryMessage, em); err != nil {
		return nil, err
	}

	if _, em, err = conn.ReadMessage(); err != nil {
		return nil, err
	}

	h2 = MakeHandshake()

	if err = codec.Decode(h2, em); err != nil {
		return nil, err
	}

	return h2, nil
}

func MakeHandshake() *Handshake {

	var handshake Handshake

	handshake.Capabilities = map[capability.Capability]string{}

	return &handshake
}
