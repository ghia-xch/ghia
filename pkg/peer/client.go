package peer

import (
	"context"
	"errors"
	"github.com/ghia-xch/ghia/pkg/node"
	"github.com/ghia-xch/ghia/pkg/protocol"
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
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
	info      *PeerInfo
	conn      *websocket.Conn
	handshake *Handshake
	inbound   chan primitive.EncodedMessage
	outbound  chan primitive.EncodedMessage
	callbacks chan primitive.Callback
	handlers  map[primitive.MessageType]primitive.Callback
}

func (p *Client) Open(ctx context.Context, timeout time.Duration) (err error) {

	p.Lock()

	defer p.Unlock()

	if p.conn != nil {
		return nil
	}

	var u = url.URL{
		Scheme: "wss",
		Host:   p.info.Host + ":" + strconv.Itoa(int(p.info.Port)),
		Path:   "/ws",
	}

	l.Info("opening socket to peer: ", u.String())

	websocket.DefaultDialer.TLSClientConfig = node.DefaultTLSConfig
	websocket.DefaultDialer.HandshakeTimeout = timeout

	if p.conn, _, err = websocket.DefaultDialer.DialContext(ctx, u.String(), nil); err != nil {
		return err
	}

	l.Infoln("performing handshake...")

	if p.handshake, err = PerformHandshake(p.conn, primitive.NewMessageEncoder(1024), DefaultHandshake); err != nil {
		return err
	}

	l.Infoln("handshake succeeded")

	go func() {

		var mt int
		var err error
		var msg []byte

		for {

			if mt, msg, err = p.conn.ReadMessage(); err != nil {
				if mt == -1 {
					return
				}

				continue
			}

			p.inbound <- msg
		}
	}()

	go func() {

		var msg primitive.EncodedMessage
		var cb primitive.Callback
		var ok bool

		for {

			msg = <-p.inbound

			l.Infoln("received message: ", msg)

			if protocol.HasNoExpectedResponse(msg.Type()) {

				if cb, ok = p.handlers[msg.Type()]; !ok {
					continue
				}

				if err = cb(msg); err != nil {
					l.Errorf("failed to handle message: %v", err)
				}
			}
		}
	}()

	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt)

	ticker := time.NewTicker(10 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			if err = p.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				l.Println("write:", err)
				return
			}

		case <-interrupt:

			l.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err = p.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				l.Println("write close:", err)
				return
			}
			select {
			case <-time.After(time.Second):
			}
			return
		}
	}

	return nil
}

func (c *Client) Handle(handlers ...primitive.MessageHandler) {
	for _, handler := range handlers {
		c.handlers[handler.Type] = handler.Callback
	}
}

func (p *Client) Close() (err error) {

	p.Lock()

	defer p.Unlock()

	if p.conn == nil {
		return nil
	}

	err = p.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)

	if err != nil {
		return
	}

	return nil
}

func (p *Client) Send(em primitive.EncodedMessage) (err error) {

	p.Lock()

	defer p.Unlock()

	if !protocol.HasNoExpectedResponse(em.Type()) {
		return errors.New("message has an expected response.")
	}

	select {
	case p.outbound <- em:
		return nil
	default:
		return errors.New("peer outbound is full.")
	}

	return nil
}

func (p *Client) SendWith(em primitive.EncodedMessage, cb primitive.Callback) (err error) {

	p.Lock()

	defer p.Unlock()

	if !protocol.HasExpectedResponse(em.Type()) {
		return errors.New("message has no expected response.")
	}

	select {
	case p.outbound <- em:
		p.callbacks <- cb
		return nil
	default:
		return errors.New("peer outbound is full.")
	}

	return nil
}

func NewClient(peerInfo *PeerInfo) (c *Client) {

	var client = Client{
		info:      peerInfo,
		inbound:   make(chan primitive.EncodedMessage, 128),
		outbound:  make(chan primitive.EncodedMessage, 128),
		callbacks: make(chan primitive.Callback, 128),
		handlers:  make(map[primitive.MessageType]primitive.Callback),
	}

	return &client
}
