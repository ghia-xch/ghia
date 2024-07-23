package protocol

import (
	"github.com/ghia-xch/ghia/pkg"
	"github.com/ghia-xch/ghia/pkg/protocol/capability"
	"github.com/ghia-xch/ghia/pkg/protocol/full_node"
	"github.com/ghia-xch/ghia/pkg/protocol/network"
	"github.com/ghia-xch/ghia/pkg/protocol/node"
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
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
	NetworkId       *network.Network
	ProtocolVersion string
	SoftwareVersion string
	ServerPort      uint16
	NodeType        node.Type
	Capabilities    capability.Set
}

func (h *Handshake) Encode(enc *primitive.MessageEncoder) (em primitive.EncodedMessage, err error) {

	if em, err = enc.Encode(
		h.NetworkId,
		h.ProtocolVersion,
		h.SoftwareVersion,
		h.ServerPort,
		h.NodeType,
		h.Capabilities,
	); err != nil {
		return
	}

	return enc.Bytes(), nil
}

func (h *Handshake) Decode(dec *primitive.MessageDecoder, em primitive.EncodedMessage) (err error) {

	dec.Reset(em)

	var str string

	if str, err = dec.ParseString(); err != nil {
		return
	}

	h.NetworkId = network.NewNetwork(str)

	if str, err = dec.ParseString(); err != nil {
		return
	}

	h.ProtocolVersion = str

	if str, err = dec.ParseString(); err != nil {
		return
	}

	h.SoftwareVersion = str

	var port uint16

	if port, err = dec.ParseUint16(); err != nil {
		return
	}

	h.ServerPort = port

	var nType uint8

	if nType, err = dec.ParseUint8(); err != nil {
		return
	}

	h.NodeType = node.Type(nType)

	if err = h.Capabilities.Decode(dec); err != nil {
		return
	}

	return nil
}

func PerformHandshake(conn *websocket.Conn, enc *primitive.MessageEncoder, h1 *Handshake) (h2 *Handshake, err error) {

	var em primitive.EncodedMessage

	enc.Reset(HandshakeType, nil)

	if em, err = h1.Encode(enc); err != nil {
		return nil, err
	}

	if err = conn.WriteMessage(websocket.BinaryMessage, em); err != nil {
		return nil, err
	}

	if _, em, err = conn.ReadMessage(); err != nil {
		return nil, err
	}

	h2 = MakeHandshake()

	if err = h2.Decode(primitive.NewMessageDecoder(), em); err != nil {
		return
	}

	return h2, nil
}

func MakeHandshake() *Handshake {

	var handshake Handshake

	handshake.Capabilities = map[capability.Capability]string{}

	return &handshake
}
