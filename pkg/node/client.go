package node

import (
	"context"
	"errors"
	"github.com/ghia-xch/ghia/pkg/node/capability"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/peer"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Client struct {
	sync.Mutex

	info *peer.PeerInfo

	conn      *websocket.Conn
	handshake *Handshake

	inbound  chan protocol.EncodedMessage
	outbound chan protocol.EncodedMessage

	callbacks chan protocol.Callback
	handlers  map[protocol.MessageType]protocol.Callback

	isClosing chan bool
	isClosed  chan bool
}

func (c *Client) Open(ctx context.Context, timeout time.Duration) (err error) {

	c.Lock()

	defer c.Unlock()

	if c.conn != nil {
		return nil
	}

	l.WithField("peer", c.info.Url()).Info("connection to peer, opening")

	websocket.DefaultDialer.TLSClientConfig = DefaultTLSConfig
	websocket.DefaultDialer.HandshakeTimeout = timeout

	if c.conn, _, err = websocket.DefaultDialer.DialContext(ctx, c.info.Url().String(), nil); err != nil {
		return err
	}

	if c.handshake, err = performHandshake(c.conn, protocol.NewMessageEncoder(128), DefaultHandshake); err != nil {
		return err
	}

	l.Infoln("handshake successful, connection established.")

	go c.inboundQueuing()
	go c.outboundQueuing()
	go c.handlerQueuing()

	return nil
}

func (c *Client) Handle(handlers ...protocol.MessageHandler) {

	c.Lock()

	defer c.Unlock()

	for _, handler := range handlers {
		c.handlers[handler.Type] = handler.Callback
	}
}

func (c *Client) Send(em protocol.EncodedMessage) (err error) {

	c.Lock()

	defer c.Unlock()

	if !protocol.HasNoExpectedResponse(em.Type()) {
		return errors.New("message has an expected response.")
	}

	select {
	case c.outbound <- em:
		return nil
	default:
		return errors.New("peer outbound is full.")
	}

	return nil
}

func (c *Client) SendWith(em protocol.EncodedMessage, cb protocol.Callback) (err error) {

	c.Lock()

	defer c.Unlock()

	if !protocol.HasExpectedResponse(em.Type()) {
		return errors.New("sending, message has no expected response: " + protocol.TypeAsString(em.Type()))
	}

	select {
	case c.outbound <- em:
		c.callbacks <- cb
		return nil
	default:
		return errors.New("peer outbound is full.")
	}

	return nil
}

func (c *Client) IsClosed() chan bool {
	return c.isClosed
}

func (c *Client) Close() (err error) {

	c.Lock()

	defer c.Unlock()

	if c.conn == nil {
		return nil
	}

	err = c.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)

	if err != nil {
		return
	}

	return nil
}

func (c *Client) IsCapableOf(cap capability.Capability) bool {
	return c.handshake.Capabilities.IsEnabled(cap)
}

var (
	MaxQueuedInboundMessages  = 128
	MaxQueuedOutboundMessages = 128
)

func NewClient(peerInfo *peer.PeerInfo) (c *Client) {

	var client = Client{
		info: peerInfo,

		inbound:  make(chan protocol.EncodedMessage, MaxQueuedInboundMessages),
		outbound: make(chan protocol.EncodedMessage, MaxQueuedOutboundMessages),

		callbacks: make(chan protocol.Callback, MaxQueuedOutboundMessages),
		handlers:  make(map[protocol.MessageType]protocol.Callback),

		isClosing: make(chan bool, 1),
		isClosed:  make(chan bool, 1),
	}

	return &client
}
