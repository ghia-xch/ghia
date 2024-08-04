package crawler

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/ghia-xch/ghia/pkg"
	"github.com/ghia-xch/ghia/pkg/node"
	"github.com/ghia-xch/ghia/pkg/protocol"
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/url"
	"os"
)

var crawlCommand = &cobra.Command{
	Use:   "crawl",
	Short: "Runs an instance of the crawler",
	Long:  `Runs an instance of the crawler.`,
	Run: func(cmd *cobra.Command, args []string) {

		l.Println("-- ghia (" + viper.GetString("network") + ") - " + pkg.SemVer + " - PoST Freedom. --")

		l.Debugln("DEBUG!")

		var conn *websocket.Conn
		var hs *protocol.Handshake
		var err error

		var u = url.URL{
			Scheme: "wss",
			Host:   "203.184.53.208:8444",
			Path:   "/ws",
		}

		websocket.DefaultDialer.TLSClientConfig = node.DefaultTLSConfig

		if conn, _, err = websocket.DefaultDialer.DialContext(context.Background(), u.String(), nil); err != nil {
			l.Fatalln(err)
			os.Exit(1)
		}

		if hs, err = protocol.PerformHandshake(conn, primitive.NewMessageEncoder(1024), protocol.DefaultHandshake); err != nil {
			l.Fatal("handshake:", err)
			os.Exit(1)
		}

		spew.Dump(hs)

		l.Println("-- fin --")
	},
}
