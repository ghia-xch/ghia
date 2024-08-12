package peer

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"net/url"
	"strconv"
	"time"
)

type PeerInfo struct {
	Host string
	Port uint16
}

func (pi *PeerInfo) Url() *url.URL {
	return &url.URL{
		Scheme: "wss",
		Host:   pi.Host + ":" + strconv.Itoa(int(pi.Port)),
		Path:   "/ws",
	}
}

type TimestamptedPeerInfo struct {
	PeerInfo
	Ts uint64
}

func (r *TimestamptedPeerInfo) Timestamp() time.Time {
	return time.Unix(int64(r.Ts), 0)
}

func (r *TimestamptedPeerInfo) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {
	return enc.Encode(r.Host, r.Port, r.Ts)
}

func (r *TimestamptedPeerInfo) Decode(dec *protocol.MessageDecoder, em protocol.EncodedMessage) (err error) {

	if r.Host, err = dec.ParseString(); err != nil {
		return err
	}

	if r.Port, err = dec.ParseUint16(); err != nil {
		return err
	}

	if r.Ts, err = dec.ParseUint64(); err != nil {
		return err
	}

	return nil
}

func NewPeerInfo(host string, port uint16) *PeerInfo {
	return &PeerInfo{
		Host: host,
		Port: port,
	}
}
