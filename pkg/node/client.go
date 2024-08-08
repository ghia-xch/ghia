package node

import (
	"context"
	"errors"
	"github.com/ghia-xch/ghia/pkg/node/capability"
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/peer"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	sync.Mutex
	info      *peer.PeerInfo
	conn      *websocket.Conn
	handshake *Handshake
	inbound   chan protocol.EncodedMessage
	outbound  chan protocol.EncodedMessage
	callbacks chan protocol.Callback
	handlers  map[protocol.MessageType]protocol.Callback

	isClosed chan bool
}

func (c *Client) Open(ctx context.Context, timeout time.Duration) (err error) {

	c.Lock()

	defer c.Unlock()

	if c.conn != nil {
		return nil
	}

	var u = url.URL{
		Scheme: "wss",
		Host:   c.info.Host + ":" + strconv.Itoa(int(c.info.Port)),
		Path:   "/ws",
	}

	l.Info("opening socket to peer: ", u.String())

	websocket.DefaultDialer.TLSClientConfig = DefaultTLSConfig
	websocket.DefaultDialer.HandshakeTimeout = timeout

	if c.conn, _, err = websocket.DefaultDialer.DialContext(ctx, u.String(), nil); err != nil {
		return err
	}

	l.Infoln("performing handshake...")

	if c.handshake, err = PerformHandshake(c.conn, protocol.NewMessageEncoder(1024), DefaultHandshake); err != nil {
		return err
	}

	l.Infoln("handshake succeeded")

	go func() {

		var err error
		var msg []byte

		for {

			if _, msg, err = c.conn.ReadMessage(); err != nil {

				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					l.Errorln("read error: %v", err)
					return
				}

				continue
			}

			c.inbound <- msg
		}
	}()

	go func() {

		var msg protocol.EncodedMessage
		var cb protocol.Callback
		var ok bool

		var msgDecoder = protocol.NewMessageDecoder()

		for {

			msg = <-c.inbound

			l.Info("received message[", msg.Type(), "] ", protocol.TypeAsString(msg.Type()))

			if protocol.HasNoExpectedResponse(msg.Type()) {

				if cb, ok = c.handlers[msg.Type()]; !ok {
					continue
				}

				if err = msgDecoder.Reset(msg); err != nil {
					l.Errorln("failed to decode message: %v", err)
				}

				if err = cb(msgDecoder); err != nil {
					l.Errorf("failed to handle message: %v", err)
				}
			}
		}
	}()

	go func() {

		interrupt := make(chan os.Signal, 1)

		signal.Notify(interrupt, os.Interrupt)

		ticker := time.NewTicker(10 * time.Second)

		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:

				if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					l.Println("write:", err)
					return
				}

			case <-interrupt:

				l.Println("interrupt, closing websocket")

				if err = c.Close(); err != nil {
					l.Errorln("close:", err)
				}

				return
			}
		}
	}()

	return nil
}

func (c *Client) Handle(handlers ...protocol.MessageHandler) {
	for _, handler := range handlers {
		c.handlers[handler.Type] = handler.Callback
	}
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

	c.isClosed <- true

	if err != nil {
		return
	}

	return nil
}

func (c *Client) IsClosed() chan bool {
	return c.isClosed
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
		return errors.New("message has no expected response.")
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

func NewClient(peerInfo *peer.PeerInfo) (c *Client) {

	var client = Client{
		info:      peerInfo,
		inbound:   make(chan protocol.EncodedMessage, 128),
		outbound:  make(chan protocol.EncodedMessage, 128),
		callbacks: make(chan protocol.Callback, 128),
		handlers:  make(map[protocol.MessageType]protocol.Callback),
		isClosed:  make(chan bool, 1),
	}

	return &client
}
