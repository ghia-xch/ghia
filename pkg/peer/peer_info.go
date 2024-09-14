package peer

import (
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

func NewPeerInfo(host string, port uint16) *PeerInfo {
	return &PeerInfo{
		Host: host,
		Port: port,
	}
}

func NewTimestampedPeerInfo(host string, port uint16, ts uint64) *TimestamptedPeerInfo {
	return &TimestamptedPeerInfo{
		PeerInfo: *NewPeerInfo(host, port),
		Ts:       ts,
	}
}
