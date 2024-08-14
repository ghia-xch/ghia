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

	go func() {

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
	}()

	return nil
}

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
				l.Println("write:", err)
				return
			}

			ticker.Reset(10 * time.Second)
		}
	}
}

func (c *Client) handleInboundMessage(dec *protocol.MessageDecoder, em protocol.EncodedMessage) (err error) {

	var ok bool
	var cb protocol.Callback

	if protocol.HasNoExpectedResponse(em.Type()) {

		if cb, ok = c.handlers[em.Type()]; !ok {
			return
		}

	} else {

		if cb, ok = <-c.callbacks; !ok {
			return err
		}
	}

	if err = dec.Reset(em); err != nil {
		return
	}

	if err = cb(dec); err != nil {
		return
	}

	return nil

	// Look for callback and exec
	return nil
}

func (c *Client) Handle(handlers ...protocol.MessageHandler) {
	for _, handler := range handlers {
		c.handlers[handler.Type] = handler.Callback
	}
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
