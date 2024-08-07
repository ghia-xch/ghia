package peer

import (
	"github.com/ghia-xch/ghia/pkg/protocol/primitive"
	"time"
)

type PeerInfo struct {
	Host string
	Port uint16
}

type TimestamptedPeerInfo struct {
	PeerInfo
	Ts uint64
}

func (r *TimestamptedPeerInfo) Timestamp() time.Time {
	return time.Unix(int64(r.Ts), 0)
}

func (r *TimestamptedPeerInfo) Encode(enc *primitive.MessageEncoder) (em primitive.EncodedMessage, err error) {
	return enc.Encode(r.Host, r.Port, r.Ts)
}

func (r *TimestamptedPeerInfo) Decode(dec *primitive.MessageDecoder, em primitive.EncodedMessage) (err error) {

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
