package main

import (
	"crypto/tls"
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg/protocol"
	"github.com/ghia-xch/ghia/pkg/protocol/message"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {

	var conn *websocket.Conn
	var hs *protocol.Handshake
	var err error

	var u = url.URL{
		Scheme: "wss",
		Host:   "203.184.53.208:8444",
		Path:   "/ws",
	}

	c, _ := tls.LoadX509KeyPair("keys/public_full_node.crt", "keys/public_full_node.key")

	websocket.DefaultDialer.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{c},
	}

	if conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil); err != nil {
		log.Fatal("dial:", err)
		return
	}

	defer conn.Close()

	log.Println("connected")

	// var em message.EncodedMessage
	//
	// em, err = full_node.RequestBlocksMessage(1, 100, false).Encode(message.NewMessageEncoder(1024))
	//
	// rbm = full_node.RequestBlocks{}
	// if err = rbm.Decode(em); err != nil {
	//	return
	// }

	if hs, err = protocol.PerformHandshake(conn, message.NewMessageEncoder(1024), protocol.DefaultHandshake); err != nil {
		log.Fatal("handshake:", err)
		return
	}

	//////
	spew.Dump(hs.NetworkId)

	done := make(chan struct{})

	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt)

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

	return
}
